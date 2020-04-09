package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/sirupsen/logrus"
)

type UserIdMap map[int]User
type UserEmailMap map[string]User

type User struct {
	Id    int    `json:"id"`
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

	// unmarshall into the data strcture
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

type TeamIdMap map[int]Team
type TeamNameMap map[string]Team

type Team struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func parseTeamListToTeamMaps(payload interface{}) (TeamIdMap, TeamNameMap) {
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
		Results  []Team
	}

	// unmarshall into the data strcture
	t := tmpStruct{}
	err = json.Unmarshal(jsonbytes, &t)
	if err != nil {
		fmt.Println(err)
	}

	// load into our own data struture
	teamIdMap := TeamIdMap{}
	teamNameMap := TeamNameMap{}
	for _, team := range t.Results {
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
	Id     int `json:"id"`
	UserId int `json:"apex_user_id"`
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

	// unmarshall into the data strcture
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
			User:           user,
			UserEmail:      user.Email,
			Team:           team,
			TeamName:       team.Name,
			TeamMembership: teamMembership,
		}
	}
	return teamMembershipMap
}

type OrgTeamList []struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func parseOrgTeamList(payloadFile string) OrgTeamList {
	// we have to re-marshal the payload into our own data structure
	payload, _ := ioutil.ReadFile(payloadFile)

	// unmarshall into the data strcture
	t := OrgTeamList{}
	err := json.Unmarshal([]byte(payload), &t)
	if err != nil {
		fmt.Println(err)
	}

	return t
}
