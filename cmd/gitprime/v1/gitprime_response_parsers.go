package v1

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	runtime "github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	api_client "github.com/cisco-sso/gitprime-cli/client"
	api_teamMembership "github.com/cisco-sso/gitprime-cli/client/team_membership"
	api_teams "github.com/cisco-sso/gitprime-cli/client/teams"
	api_users "github.com/cisco-sso/gitprime-cli/client/users"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getAllUserList(client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) []User {
	numToRequest := int64(1000)
	ordering := "id" // must request in order or else not all results returned (True for teamMemberships)
	userParams := api_users.NewUsersListParams()

	// define the data structure into which to unmarshall
	type tmpStruct struct {
		Count    int64  `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []User
	}
	t := tmpStruct{}

	users := []User{}
	for offset := int64(0); offset <= t.Count; offset += numToRequest {
		userParams.Offset = &offset
		userParams.Limit = &numToRequest
		userParams.Ordering = &ordering
		userResp, _ := client.Users.UsersList(userParams, authInfo) // TODO: handle err
		payload := userResp.Payload

		// Re-Marshal into our own data structure
		////  Marshall and turn the payload back into a json string
		jsonbytes, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}
		////  Unmarshall into our own data structure
		err = json.Unmarshal(jsonbytes, &t)
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}

		for _, ele := range t.Results {
			users = append(users, ele)
		}
	}
	return users
}

type UserIdMap map[int64]User
type UserEmailMap map[string]User

func parseUserListToUserMaps(userList []User) (UserIdMap, UserEmailMap) {
	// load into our own data strutures
	userIdMap := UserIdMap{}
	userEmailMap := UserEmailMap{}
	for _, user := range userList {
		userIdMap[user.Id] = user
		userEmailMap[user.Email] = user
	}
	return userIdMap, userEmailMap
}

type Team struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent"`

	Description string `json:"description"`
}

func getAllTeamList(client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) []Team {
	numToRequest := int64(1000)
	ordering := "id" // must request in order or else not all results returned (True for teamMemberships)
	teamParams := api_teams.NewTeamsListParams()

	// define the data structure into which to unmarshall
	type tmpStruct struct {
		Count    int64  `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []Team
	}
	t := tmpStruct{}

	teams := []Team{}
	for offset := int64(0); offset <= t.Count; offset += numToRequest {
		teamParams.Offset = &offset
		teamParams.Limit = &numToRequest
		teamParams.Ordering = &ordering
		teamResp, _ := client.Teams.TeamsList(teamParams, authInfo) // TODO: handle err
		payload := teamResp.Payload

		// Re-Marshal into our own data structure
		////  Marshall and turn the payload back into a json string
		jsonbytes, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}
		////  Unmarshall into our own data structure
		err = json.Unmarshal(jsonbytes, &t)
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}

		for _, ele := range t.Results {
			teams = append(teams, ele)
		}
	}
	return teams
}

type TeamIdMap map[int64]Team
type TeamNameMap map[string]Team

func parseTeamListToTeamMaps(teamList []Team) (TeamIdMap, TeamNameMap) {
	// load into our own data strutures
	teamIdMap := TeamIdMap{}
	teamNameMap := TeamNameMap{}
	for _, team := range teamList {
		teamIdMap[team.Id] = team
		teamNameMap[team.Name] = team
	}
	return teamIdMap, teamNameMap
}

type TeamMembership struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"apex_user_id"`
	Team   Team
}

func getAllTeamMembershipList(client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) []TeamMembership {
	numToRequest := int64(1000)
	ordering := "id" // must request in order or else not all results returned (True for teamMemberships)
	teamMembershipParams := api_teamMembership.NewTeamMembershipListParams()

	// define the data structure into which to unmarshall
	type tmpStruct struct {
		Count    int64  `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []TeamMembership
	}
	t := tmpStruct{}

	teamMemberships := []TeamMembership{}
	for offset := int64(0); offset <= t.Count; offset += numToRequest {
		teamMembershipParams.Offset = &offset
		teamMembershipParams.Limit = &numToRequest
		teamMembershipParams.Ordering = &ordering
		teamMembershipResp, _ := client.TeamMembership.TeamMembershipList(teamMembershipParams, authInfo) // TODO: handle err
		payload := teamMembershipResp.Payload

		// Re-Marshal into our own data structure
		////  Marshall and turn the payload back into a json string
		jsonbytes, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}
		////  Unmarshall into our own data structure
		err = json.Unmarshal(jsonbytes, &t)
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
		}

		for _, ele := range t.Results {
			teamMemberships = append(teamMemberships, ele)
		}
	}
	return teamMemberships
}

type TeamMembershipMap map[string]map[string]TeamMembershipItem // map[user.email][team.name]
type TeamMembershipItem struct {
	User           User
	UserEmail      string
	Team           Team
	TeamName       string
	TeamMembership TeamMembership
}

func parseTeamMembershipListToMaps(userIdMap UserIdMap, teamMembershipList []TeamMembership) TeamMembershipMap {
	// load into our own data struture
	teamMembershipMap := TeamMembershipMap{}
	for _, teamMembership := range teamMembershipList {
		user := userIdMap[teamMembership.UserId]
		team := teamMembership.Team

		if _, present := teamMembershipMap[user.Email]; !present {
			teamMembershipMap[user.Email] = make(map[string]TeamMembershipItem)
		}

		tmi := TeamMembershipItem{
			UserEmail:      user.Email,
			TeamName:       team.Name,
			User:           user,
			Team:           team,
			TeamMembership: teamMembership,
		}
		if _, present := teamMembershipMap[user.Email][team.Name]; !present {
			teamMembershipMap[user.Email][team.Name] = tmi
		} else {
			log.WithFields(logrus.Fields{"email": user.Email, "team": team.Name}).Fatalf(
				"Code or API error: Duplicate user membership in team")
		}
	}
	return teamMembershipMap
}

type TeamDefinitionList []TeamDefinition
type TeamDefinitionNameMap map[string]TeamDefinition

type TeamDefinition struct {
	Name    string   `json:"name" yaml:"name"`
	Members []string `json:"members" yaml:"members"`
}

func parseTeamDefinitionFile(file string) (TeamDefinitionList, TeamDefinitionNameMap) {
	// we have to re-marshal the payload into our own data structure
	payload, _ := ioutil.ReadFile(file)

	// unmarshall into the data structure
	teamDefinitionList := TeamDefinitionList{}
	var err error
	if strings.HasSuffix(file, ".json") {
		err = json.Unmarshal([]byte(payload), &teamDefinitionList)
	} else if strings.HasSuffix(file, ".yaml") {
		err = yaml.Unmarshal([]byte(payload), &teamDefinitionList)
	} else {
		log.WithFields(logrus.Fields{"file": file}).Fatalf("File must end in .json or .yaml")
	}
	if err != nil {
		log.WithFields(logrus.Fields{"err": err, "file": file}).Fatalf("Failed to read data from file")
	}

	teamDefinitionNameMap := TeamDefinitionNameMap{}
	for _, e := range teamDefinitionList {
		teamDefinitionNameMap[e.Name] = e
	}

	return teamDefinitionList, teamDefinitionNameMap
}

func recurseEnsureTeamExists(teamNames []string, teamIdMap TeamIdMap,
	teamNameMap TeamNameMap, teamDescriptionTag string, client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) *Team {

	if len(teamNames) == 0 {
		return nil
	} else {
		currentTeamName := teamNames[len(teamNames)-1]
		ancestorTeamNames := teamNames[:len(teamNames)-1]
		parentTeamObj := recurseEnsureTeamExists(ancestorTeamNames, teamIdMap, teamNameMap, teamDescriptionTag, client, authInfo)

		if _, present := teamNameMap[currentTeamName]; present {
			log.WithFields(logrus.Fields{"team": currentTeamName}).Tracef("Team already Exists")
		} else {
			var parentTeamId int64
			var parentTeamName string
			if parentTeamObj != nil {
				parentTeamId = parentTeamObj.Id
				parentTeamName = parentTeamObj.Name
			}

			log.WithFields(logrus.Fields{"team": currentTeamName, "parentTeam": parentTeamName, "parentTeamId": parentTeamId}).Infof("Creating new Team")

			params := api_teams.NewTeamsCreateParams()
			org_helper := int64(1)
			depth_helper := "inherit"
			show_helper := "SHOW"
			params.Data = api_teams.TeamsCreateBody{
				Name:        &currentTeamName,
				Parent:      parentTeamId,
				Description: teamDescriptionTag,
				Org:         &org_helper,
				Depth:       &depth_helper,
				Visibility:  &show_helper,
			}
			resp, err := client.Teams.TeamsCreate(params, authInfo)
			if err != nil {
				failMsg := "Operation Failed: To debug, run with environment variable DEBUG=1 set"
				log.WithFields(logrus.Fields{"err": err}).Fatalf(failMsg)
			}

			// we have to re-marshal the payload into our own data structure

			// turn the payload back into a json string
			payload := resp.Payload
			jsonbytes, err := json.MarshalIndent(payload, "", "  ")
			if err != nil {
				log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
			}

			// unmarshall into the data structure
			t := Team{}
			err = json.Unmarshal(jsonbytes, &t)
			if err != nil {
				log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
			}
			teamIdMap[t.Id] = t
			teamNameMap[t.Name] = t

		}
		t := teamNameMap[currentTeamName]
		return &t
	}
}

func iterateEnsureUsersInTeam(teamDefinition TeamDefinition, userEmailMap UserEmailMap, teamNameMap TeamNameMap,
	teamMembershipMap TeamMembershipMap, client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) {

	// member is an email address
	for _, member := range teamDefinition.Members {
		userObj, present := userEmailMap[member]
		if !present {
			// This member has never been seen in GitPrime before.
			//   They have not checked in code
			log.WithFields(logrus.Fields{"user": member}).Tracef("Skipping add of user not seen in GitPrime")
			continue
		} // The user is indeed in GitPrime

		segments := strings.Split(teamDefinition.Name, ".")
		teamDefinitionName := segments[len(segments)-1]
		teamObj, present := teamNameMap[teamDefinitionName]
		if !present {
			log.WithFields(logrus.Fields{"team": teamDefinitionName}).Fatalf("Code or API error: Team must exist in GitPrime")
			continue
		} // The team is indeed in GitPrime

		if _, present := teamMembershipMap[member]; !present {
			teamMembershipMap[member] = make(map[string]TeamMembershipItem)
		}

		if _, present := teamMembershipMap[member][teamDefinitionName]; present {
			log.WithFields(logrus.Fields{"user": member, "team": teamDefinitionName}).Tracef("User team membership already exists")
		} else {
			log.WithFields(logrus.Fields{"user": member, "team": teamDefinitionName}).Infof("Adding user to team")

			// make the call to create the teamMembership user->team
			params := api_teamMembership.NewTeamMembershipCreateParams()
			depth_helper := "inherit"
			params.Data = api_teamMembership.TeamMembershipCreateBody{
				ApexUserID:     &userObj.Id,
				TeamID:         &teamObj.Id,
				Depth:          &depth_helper,
				MembershipType: "contributor",
			}

			resp, err := client.TeamMembership.TeamMembershipCreate(params, authInfo)
			if err != nil {
				failMsg := "Operation Failed: To debug, run with environment variable DEBUG=1 set"
				log.WithFields(logrus.Fields{"err": err}).Fatalf(failMsg)
			}

			// we have to re-marshal the payload into our own data structure

			// turn the payload back into a json string
			payload := resp.Payload
			jsonbytes, err := json.MarshalIndent(payload, "", "  ")
			if err != nil {
				log.WithFields(logrus.Fields{"err": err, "payload": payload}).Fatalf("failed")
			}

			// unmarshall into the data structure
			t := TeamMembership{}
			err = json.Unmarshal(jsonbytes, &t)
			if err != nil {
				log.WithFields(logrus.Fields{"err": err}).Fatalf("failed")
			}

			teamMembershipMap[member][teamDefinitionName] = TeamMembershipItem{
				User:           userObj,
				UserEmail:      member,
				Team:           teamObj,
				TeamName:       teamObj.Name,
				TeamMembership: t,
			}
		}
	}
}
