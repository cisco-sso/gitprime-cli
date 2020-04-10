package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	runtime "github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"

	api_client "github.com/cisco-sso/gitprime-cli/client"
	api_teamMembership "github.com/cisco-sso/gitprime-cli/client/team_membership"
	api_teams "github.com/cisco-sso/gitprime-cli/client/teams"
)

type UserIdMap map[int64]User
type UserEmailMap map[string]User

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func parseUserListToUserMaps(payload interface{}) (UserIdMap, UserEmailMap) {
	// we have to re-marshal the payload into our own data structure

	// turn the payload back into a json string
	jsonbytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
	}

	// define the data structure into which to unmarshall
	type tmpStruct struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []User
	}

	// unmarshall into the data structure
	t := tmpStruct{}
	err = json.Unmarshal(jsonbytes, &t)
	if err != nil {
		fmt.Println(err)
	}

	// load into our own data struture
	userIdMap := UserIdMap{}
	userEmailMap := UserEmailMap{}
	for _, user := range t.Results {
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

func getAllTeams(client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) []Team {
	numToRequest := int64(1000)
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
		teamResp, _ := client.Teams.TeamsList(teamParams, authInfo) // TODO: handle err
		payload := teamResp.Payload

		// Re-Marshal into our own data structure
		////  Marshall and turn the payload back into a json string
		jsonbytes, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
		}
		////  Unmarshall into our own data structure
		err = json.Unmarshal(jsonbytes, &t)
		if err != nil {
			log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
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

// map[user.email][team.name]
type TeamMembershipMap map[string]map[string]TeamMembershipItem

type TeamMembershipItem struct {
	User           User
	UserEmail      string
	Team           Team
	TeamName       string
	TeamMembership TeamMembership
}

type TeamMembership struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"apex_user_id"`
	Team   Team
}

func parseTeamMembershipListToTeamMembershipMap(payload interface{}, userIdMap UserIdMap, teamIdMap TeamIdMap, emailFilterRegexp *regexp.Regexp, teamDescriptionTag string) TeamMembershipMap {
	// we have to re-marshal the payload into our own data structure

	// turn the payload back into a json string
	jsonbytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
	}

	// define the data structure into which to unmarshall
	type tmpStruct struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []TeamMembership
	}

	// unmarshall into the data structure
	t := tmpStruct{}
	err = json.Unmarshal(jsonbytes, &t)
	if err != nil {
		fmt.Println(err)
	}

	// load into our own data struture
	teamMembershipMap := TeamMembershipMap{}
	for _, teamMembership := range t.Results {
		user := userIdMap[teamMembership.UserId]
		team := teamMembership.Team

		if !emailFilterRegexp.MatchString(user.Email) {
			fmt.Println("skipping add of user [%v]", user.Email)
			continue
		}

		if _, ok := teamMembershipMap[user.Email]; !ok {
			teamMembershipMap[user.Email] = make(map[string]TeamMembershipItem)
		}

		teamMembershipMap[user.Email][team.Name] = TeamMembershipItem{
			UserEmail:      user.Email,
			TeamName:       team.Name,
			User:           user,
			Team:           team,
			TeamMembership: teamMembership,
		}
	}
	return teamMembershipMap
}

type OrgTeamList []OrgTeam
type OrgTeamNameMap map[string]OrgTeam

type OrgTeam struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func parseOrgTeamList(payloadFile string) (OrgTeamList, OrgTeamNameMap) {
	// we have to re-marshal the payload into our own data structure
	payload, _ := ioutil.ReadFile(payloadFile)

	// unmarshall into the data structure
	orgTeamList := OrgTeamList{}
	err := json.Unmarshal([]byte(payload), &orgTeamList)
	if err != nil {
		fmt.Println(err)
	}

	orgTeamNameMap := OrgTeamNameMap{}
	for _, e := range orgTeamList {
		orgTeamNameMap[e.Name] = e
	}

	return orgTeamList, orgTeamNameMap
}

func recurseEnsureTeamExists(teamNames []string, teamIdMap TeamIdMap,
	teamNameMap TeamNameMap, teamDescriptionTag string, client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) *Team {

	if len(teamNames) == 0 {
		return nil
	} else {
		currentTeamName := teamNames[len(teamNames)-1]
		ancestorTeamNames := teamNames[:len(teamNames)-1]
		parentTeamObj := recurseEnsureTeamExists(ancestorTeamNames, teamIdMap, teamNameMap, teamDescriptionTag, client, authInfo)

		if _, present := teamNameMap[currentTeamName]; !present {
			var parentTeamId int64
			var parentTeamName string
			if parentTeamObj != nil {
				parentTeamId = parentTeamObj.Id
				parentTeamName = parentTeamObj.Name
			}
			fmt.Printf("Creating New Team %s with Parent Team Name %s Id %d\n", currentTeamName, parentTeamName, parentTeamId)

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
				log.WithFields(logrus.Fields{"err": err}).Errorf(failMsg)
			}

			// we have to re-marshal the payload into our own data structure

			// turn the payload back into a json string
			payload := resp.Payload
			jsonbytes, err := json.MarshalIndent(payload, "", "  ")
			if err != nil {
				log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
			}

			// unmarshall into the data structure
			t := Team{}
			err = json.Unmarshal(jsonbytes, &t)
			if err != nil {
				fmt.Println(err)
			}
			teamIdMap[t.Id] = t
			teamNameMap[t.Name] = t

		} else {
			fmt.Printf("Team %s already exists\n", currentTeamName)
		}
		t := teamNameMap[currentTeamName]
		return &t
	}
}

func iterateEnsureUsersInTeam(orgTeam OrgTeam, userEmailMap UserEmailMap, teamNameMap TeamNameMap,
	teamMembershipMap TeamMembershipMap, client *api_client.GitprimeCli, authInfo runtime.ClientAuthInfoWriter) {

	// member is an email address
	for _, member := range orgTeam.Members {
		userObj, present := userEmailMap[member]
		if !present {
			// This member has never been seen in GitPrime before.
			//   They have not checked in code
			fmt.Printf("Warning: user %s not found in GitPrime\n", member)
			continue
		} // The user is indeed in GitPrime

		segments := strings.Split(orgTeam.Name, ".")
		orgTeamName := segments[len(segments)-1]
		teamObj, present := teamNameMap[orgTeamName]
		if !present {
			// This team currently does not exist in GitPrime
			//   They have not checked in code
			fmt.Printf("Warning: team %s does not exist\n", orgTeamName)
			continue
		} // The team is indeed in GitPrime

		if _, present := teamMembershipMap[member]; !present {
			teamMembershipMap[member] = make(map[string]TeamMembershipItem)
		}

		if _, present := teamMembershipMap[member][orgTeamName]; !present {
			fmt.Printf("Adding user %s to team %s\n", member, orgTeamName)

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
				log.WithFields(logrus.Fields{"err": err}).Errorf(failMsg)
			}

			// we have to re-marshal the payload into our own data structure

			// turn the payload back into a json string
			payload := resp.Payload
			jsonbytes, err := json.MarshalIndent(payload, "", "  ")
			if err != nil {
				log.WithFields(logrus.Fields{"err": err, "payload": payload}).Errorf("failed")
			}

			// unmarshall into the data structure
			t := TeamMembership{}
			err = json.Unmarshal(jsonbytes, &t)
			if err != nil {
				fmt.Println(err)
			}

			teamMembershipMap[member][orgTeamName] = TeamMembershipItem{
				User:           userObj,
				UserEmail:      member,
				Team:           teamObj,
				TeamName:       teamObj.Name,
				TeamMembership: t,
			}
		} else {
			fmt.Printf("Team Membership Already Exists: User %s to team %s\n", member, orgTeamName)
		}
	}
}
