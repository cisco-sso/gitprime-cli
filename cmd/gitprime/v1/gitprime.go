package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
		app_Team             = app.Command("team", "Team commands")
		app_Team_List        = app_Team.Command("list", "List all teams")
		app_Team_Get         = app_Team.Command("get", "Get a specific team")
		app_Team_Get__TeamId = app_Team_Get.Arg("team-id", "ID of the team").Required().Int64()

		///////////////////////////////////////
		// gitprime teammemberships
		app_TeamMembership                       = app.Command("teammembership", "Team membership commands")
		app_TeamMembership_List                  = app_TeamMembership.Command("list", "List all teams")
		app_TeamMembership_Get                   = app_TeamMembership.Command("get", "Get a specific teammembership")
		app_TeamMembership_Get__TeamMembershipId = app_TeamMembership_Get.Arg("teammembership-id", "ID of the teammembership").Required().Int64()

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
		// gitprime version
		appVersion = app.Command("version", "Display version information.")
	)
	log := logger.New()

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
	// authInfo := httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", a.Secrets.AuthToken)
	// authInfo := httptransport.BearerToken(a.Secrets.AuthToken)
	authInfo := httptransport.APIKeyAuth("Authorization", "header", "Bearer "+a.Secrets.AuthToken)

	switch p {

	case app_Team_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsListParams()
			return client.Teams.TeamsList(params, authInfo)
		}
		printPayload(f, log)

	case app_Team_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_teams.NewTeamsReadParams()
			params.ID = *app_Team_Get__TeamId
			return client.Teams.TeamsRead(params, authInfo)
		}
		printPayload(f, log)

	case app_TeamMembership_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipListParams()
			return client.TeamMembership.TeamMembershipList(params, authInfo)
		}
		printPayload(f, log)

	case app_TeamMembership_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_teamMembership.NewTeamMembershipReadParams()
			params.ID = *app_TeamMembership_Get__TeamMembershipId
			return client.TeamMembership.TeamMembershipRead(params, authInfo)
		}
		printPayload(f, log)

	case app_User_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_users.NewUsersListParams()
			return client.Users.UsersList(params, authInfo)
		}
		printPayload(f, log)

	case app_User_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_users.NewUsersReadParams()
			params.ID = *app_User_Get__UserId
			return client.Users.UsersRead(params, authInfo)
		}
		printPayload(f, log)

	case app_UserAlias_List.FullCommand():
		f := func() (interface{}, error) {
			params := api_userAlias.NewUserAliasListParams()
			return client.UserAlias.UserAliasList(params, authInfo)
		}
		printPayload(f, log)

	case app_UserAlias_Get.FullCommand():
		f := func() (interface{}, error) {
			params := api_userAlias.NewUserAliasReadParams()
			params.ID = *app_UserAlias_Get__UserAliasId
			return client.UserAlias.UserAliasRead(params, authInfo)
		}
		printPayload(f, log)

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

func printPayload(f func() (interface{}, error), log *logger.Logger) {
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
