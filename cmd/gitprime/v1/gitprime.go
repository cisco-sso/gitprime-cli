package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/cisco-sso/gitprime-cli/info"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	pkg "github.com/cisco-sso/gitprime-cli/pkg/v1"
	logger "github.com/cisco-sso/gitprime-cli/wrap/logrus/v1"
	profile "github.com/cisco-sso/gitprime-cli/wrap/profile/v1"
	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	api_client "github.com/cisco-sso/gitprime-cli/client"
	api_teamMembership "github.com/cisco-sso/gitprime-cli/client/team_membership"
	api_teams "github.com/cisco-sso/gitprime-cli/client/teams"
	api_userAlias "github.com/cisco-sso/gitprime-cli/client/user_alias"
	api_users "github.com/cisco-sso/gitprime-cli/client/users"
)

const (
	TimeFormat = time.RFC3339
)

var log *logger.Logger

func init() {
	log = logger.New()
}

func Execute() {

	// Start the profiler and defer stopping it until the program exits.
	defer profile.Start().Stop()

	var (
		// gitprime
		app          = kingpin.New("gitprime", "The Gitprime CLI command.")
		appAuthToken = app.Flag("auth-token", "API auth token. Export env var GITPRIME_AUTH_TOKEN as an alternative").
				Default("").OverrideDefaultFromEnvar("GITPRIME_AUTH_TOKEN").String()
		appLogLevel = app.Flag("log-level", "Set log-level (trace|debug|info|warn|error|fatal|panic).").
				Default("info").OverrideDefaultFromEnvar("GITPRIME_LOG_LEVEL").String()
		appServerHost = app.Flag("server-host", "Server host.").
				Default(api_client.DefaultHost).OverrideDefaultFromEnvar("GITPRIME_SERVER_HOST").String()
		appServerPort = app.Flag("server-port", "Server port.").
				Default("443").OverrideDefaultFromEnvar("GITPRIME_SERVER_PORT").Int()

		///////////////////////////////////////
		// gitprime teams
		app_Team                         = app.Command("team", "Team commands")
		app_Team_List                    = app_Team.Command("list", "List all teams")
		app_Team_Get                     = app_Team.Command("get", "Get a specific team")
		app_Team_Get__TeamId             = app_Team_Get.Arg("id", "ID of the team").Required().Int64()
		app_Team_Create                  = app_Team.Command("create", "Create a new team")
		app_Team_Create__TeamName        = app_Team_Create.Arg("name", "Name of the team").Required().String()
		app_Team_Create__TeamDescription = app_Team_Create.Flag("description", "Description of the team").Default("").String()
		app_Team_Create__TeamParentId    = app_Team_Create.Flag("parent-id", "ID of the parent team").Int64()
		app_Team_Create__TeamOrgId       = app_Team_Create.Flag("org-id", "ID of the org").Default("1").Int64()
		app_Team_Create__TeamDepth       = app_Team_Create.Flag("depth", "Depth of the team").Default("inherit").String()
		app_Team_Create__TeamVisibility  = app_Team_Create.Flag("visibility", "Visibility of the team").Default("SHOW").String()
		app_Team_Delete                  = app_Team.Command("delete", "Delete a specific team")
		app_Team_Delete__TeamId          = app_Team_Delete.Arg("id", "ID of the team").Required().Int64()

		///////////////////////////////////////
		// gitprime teammembership
		app_TeamMembership                                  = app.Command("teammembership", "Team membership commands")
		app_TeamMembership_List                             = app_TeamMembership.Command("list", "List all teams")
		app_TeamMembership_Get                              = app_TeamMembership.Command("get", "Get a specific teammembership")
		app_TeamMembership_Get__TeamMembershipId            = app_TeamMembership_Get.Arg("id", "ID of the teammembership").Required().Int64()
		app_TeamMembership_Create                           = app_TeamMembership.Command("create", "Create a new teammembership")
		app_TeamMembership_Create__TeamMembershipApexUserId = app_TeamMembership_Create.Arg("user-id", "ID of the apex user").Required().Int64()
		app_TeamMembership_Create__TeamMembershipTeamId     = app_TeamMembership_Create.Arg("team-id", "ID of the team").Required().Int64()
		app_TeamMembership_Create__TeamMembershipDepth      = app_TeamMembership_Create.Flag("depth", "Depth of the teammembership").Default("inherit").String()
		app_TeamMembership_Create__TeamMembershipType       = app_TeamMembership_Create.Flag("type", "Type of the teammembership").Default("contributor").String()
		app_TeamMembership_Delete                           = app_TeamMembership.Command("delete", "Delete a specific teammembership")
		app_TeamMembership_Delete__TeamMembershipId         = app_TeamMembership_Delete.Arg("id", "ID of the teammembership").Required().Int64()

		///////////////////////////////////////
		// gitprime users
		app_User             = app.Command("user", "User commands")
		app_User_List        = app_User.Command("list", "List all users")
		app_User_Get         = app_User.Command("get", "Get a specific user")
		app_User_Get__UserId = app_User_Get.Arg("user-id", "ID of the user").Required().Int64()

		///////////////////////////////////////
		// gitprime useraliass
		app_UserAlias                  = app.Command("useralias", "User alias commands")
		app_UserAlias_List             = app_UserAlias.Command("list", "List all users")
		app_UserAlias_Get              = app_UserAlias.Command("get", "Get a specific useralias")
		app_UserAlias_Get__UserAliasId = app_UserAlias_Get.Arg("useralias-id", "ID of the useralias").Required().Int64()

		///////////////////////////////////////
		// gitprime sync
		app_Sync                              = app.Command("sync", "Sync commands")
		app_Sync_Team                         = app_Sync.Command("team", "Team sync command")
		app_Sync_Team__TeamJsonFile           = app_Sync_Team.Arg("team-json-file", "TeamJson File").Required().ExistingFile()
		app_Sync_Team__RegexFilterEmailDomain = app_Sync_Team.Flag("regex-filter-email-domain", "Valid users will match this regex [e.g. '.*@domain.com']").Default(".*").Regexp()
		app_Sync_Team__TeamDescriptionTag     = app_Sync_Team.Flag("team-description-tag", "Tag to be added to Team description to track that a team is managed by this tool").Default("gitprime-cli(donotedit,autogen)").String()

		///////////////////////////////////////
		// gitprime version
		appVersion = app.Command("version", "Display version information.")
	)

	a := &pkg.App{
		Config:  &pkg.AppConfig{},
		Secrets: &pkg.AppSecrets{},
	}

	p := kingpin.MustParse(app.Parse(os.Args[1:]))

	// gitprime
	log.SetLevel(*appLogLevel)

	// TODO: Implement actual YAML files here. (rank: Config < Env Var < Flag)

	// Populate Secrets
	if *appAuthToken != "" {
		a.Secrets.AuthToken = *appAuthToken
	} else {
		// Check here that it is Required, so we don't affect help
		log.Fatalf(" Auth Token must be provided either by arg '--auth-token' or envvar 'GITPRIME_AUTH_TOKEN'")
	}

	// Populate Config
	if *appLogLevel != "" {
		a.Config.LogLevel = *appLogLevel
	}
	if *appServerHost != "" {
		a.Config.ServerHost = *appServerHost
	}
	if *appServerPort != 0 {
		a.Config.ServerPort = *appServerPort
	}

	// create the API client
	transport := httptransport.New(a.Config.ServerHost, api_client.DefaultBasePath, api_client.DefaultSchemes)
	client := api_client.New(transport, strfmt.Default)
	authInfo := httptransport.APIKeyAuth("Authorization", "header", "Bearer "+a.Secrets.AuthToken)

	limitHelper := int64(1000000)
	switch p {

	case app_Team_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsListParams()
			params.Limit = &limitHelper
			return client.Teams.TeamsList(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_Team_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsReadParams()
			params.ID = *app_Team_Get__TeamId
			return client.Teams.TeamsRead(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_Team_Create.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsCreateParams()
			params.Data = api_teams.TeamsCreateBody{
				Name:        app_Team_Create__TeamName,
				Parent:      *app_Team_Create__TeamParentId,
				Description: *app_Team_Create__TeamDescription,
				Org:         app_Team_Create__TeamOrgId,
				Depth:       app_Team_Create__TeamDepth,
				Visibility:  app_Team_Create__TeamVisibility,
			}
			return client.Teams.TeamsCreate(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_Team_Delete.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsDeleteParams()
			params.ID = *app_Team_Delete__TeamId
			return client.Teams.TeamsDelete(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_TeamMembership_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipListParams()
			params.Limit = &limitHelper
			return client.TeamMembership.TeamMembershipList(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_TeamMembership_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipReadParams()
			params.ID = *app_TeamMembership_Get__TeamMembershipId
			return client.TeamMembership.TeamMembershipRead(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_TeamMembership_Create.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipCreateParams()
			params.Data = api_teamMembership.TeamMembershipCreateBody{
				ApexUserID:     app_TeamMembership_Create__TeamMembershipApexUserId,
				TeamID:         app_TeamMembership_Create__TeamMembershipTeamId,
				Depth:          app_TeamMembership_Create__TeamMembershipDepth,
				MembershipType: *app_TeamMembership_Create__TeamMembershipType,
			}
			return client.TeamMembership.TeamMembershipCreate(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_TeamMembership_Delete.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipDeleteParams()
			params.ID = *app_TeamMembership_Delete__TeamMembershipId
			return client.TeamMembership.TeamMembershipDelete(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_User_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_users.NewUsersListParams()
			params.Limit = &limitHelper
			return client.Users.UsersList(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_User_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_users.NewUsersReadParams()
			params.ID = *app_User_Get__UserId
			return client.Users.UsersRead(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_UserAlias_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_userAlias.NewUserAliasListParams()
			params.Limit = &limitHelper
			return client.UserAlias.UserAliasList(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_UserAlias_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_userAlias.NewUserAliasReadParams()
			params.ID = *app_UserAlias_Get__UserAliasId
			return client.UserAlias.UserAliasRead(params, authInfo)
		}
		requestAndPrint(f, log)

	case app_Sync_Team.FullCommand():
		description_tag := *app_Sync_Team__TeamDescriptionTag
		orgTeamList, orgTeamNameMap := parseOrgTeamList(*app_Sync_Team__TeamJsonFile)

		// Read the current teams
		teams := getAllTeams(client, authInfo)
		teamIdMap, teamNameMap := parseTeamListToTeamMaps(teams)

		// Ensure teams
		for orgTeamName := range orgTeamNameMap {
			segments := strings.Split(orgTeamName, ".")
			recurseEnsureTeamExists(segments,
				teamIdMap,
				teamNameMap,
				description_tag,
				client,
				authInfo)
		}

		userParams := api_users.NewUsersListParams()
		userParams.Limit = &limitHelper
		userResp, _ := client.Users.UsersList(userParams, authInfo)
		userIdMap, userEmailMap := parseUserListToUserMaps(userResp.Payload)

		teamMembershipParams := api_teamMembership.NewTeamMembershipListParams()
		teamMembershipParams.Limit = &limitHelper
		teamMembershipResp, _ := client.TeamMembership.TeamMembershipList(teamMembershipParams, authInfo)
		teamMembershipMap := parseTeamMembershipListToTeamMembershipMap(
			teamMembershipResp.Payload,
			userIdMap,
			teamIdMap,
			*app_Sync_Team__RegexFilterEmailDomain,
			description_tag)

		// Ensure users are associated with teams
		for _, orgTeam := range orgTeamList {
			iterateEnsureUsersInTeam(orgTeam, userEmailMap, teamNameMap, teamMembershipMap, client, authInfo)
		}

	case appVersion.FullCommand():
		type Version struct {
			Program         string `json:"program"`
			License         string `json:"license"`
			URL             string `json:"url"`
			BuildUser       string `json:"build_user"`
			BuildDate       string `json:"build_date"`
			Language        string `json:"language"`
			LanguageVersion string `json:"language_version"`
			Version         string `json:"version"`
			Revision        string `json:"revision"`
			Branch          string `json:"branch"`
		}
		version := Version{
			Program:         info.Program,
			License:         info.License,
			URL:             info.URL,
			BuildUser:       info.BuildUser,
			BuildDate:       info.BuildDate,
			Language:        info.Language,
			LanguageVersion: info.LanguageVersion,
			Version:         info.Version,
			Revision:        info.Revision,
			Branch:          info.Branch,
		}
		versionBytes, _ := json.MarshalIndent(version, "", "  ")
		fmt.Println(string(versionBytes))
	}
}

func requestAndPrint(f func() (interface{}, error), log *logger.Logger) {
	log.WithFields(logrus.Fields{"args": os.Args}).Tracef("called")

	type PayloadInterface interface {
		GetPayload() interface{}
	}

	resp, err := f()
	if err != nil {
		failMsg := "Operation Failed: To debug, run with environment variable DEBUG=1 set"

		apiError, ok := err.(*runtime.APIError)
		if !ok {
			log.WithFields(logrus.Fields{"err": err}).Errorf(failMsg)
		}
		response, ok := apiError.Response.(runtime.ClientResponse)
		if !ok {
			log.WithFields(logrus.Fields{"err": err}).Errorf(failMsg)
		}
		body, err := ioutil.ReadAll(response.Body())
		if err != nil {
			log.WithFields(logrus.Fields{"err": err}).Errorf(failMsg)
		}
		// Upon error 400, we cannot seem to read the body of the http response for printing.
		//    It errors out in the ioutil.ReadAll statement above.
		//    TODO: Fix this by tuning the swagger file.
		//    https://github.com/go-openapi/runtime/issues/121
		fmt.Println(string(body))
		os.Exit(1)
	}

	var str interface{}
	p, ok := resp.(PayloadInterface)
	if ok {
		str = p.GetPayload()
	} else {
		str = resp
	}

	json, err := json.MarshalIndent(str, "", "  ")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err, "str": str}).Errorf("failed")
	}
	fmt.Println(string(json))
}

func printAsJson(o interface{}) {
	out, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(out))
}
