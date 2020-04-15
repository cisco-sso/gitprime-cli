# gitprime-cli

Golang-based CLI for making API calls against Plauralsight's GitPrime (Flow) product.

Also contains subcommands for syncing users and teams from a structured YAML or JSON file to the GitPrime (Flow) product.

## Usage

```
$ ./gitprime-cli --help
usage: gitprime [<flags>] <command> [<args> ...]

The Gitprime CLI command.

Flags:
  --help              Show context-sensitive help (also try --help-long and --help-man).
  --auth-token=""     API auth token. Export env var GITPRIME_AUTH_TOKEN as an alternative
  --log-level="info"  Set log-level (trace|debug|info|warn|error|fatal|panic).
  --server-host="flow.pluralsight.com"
                      Server host.
  --server-port=443   Server port.

Commands:
  help [<command>...]
    Show help.

  team list
    List all teams

  team get <id>
    Get a specific team

  team create [<flags>] <name>
    Create a new team

  team delete <id>
    Delete a specific team

  teammembership list
    List all teams

  teammembership get <id>
    Get a specific teammembership

  teammembership create [<flags>] <user-id> <team-id>
    Create a new teammembership

  teammembership delete <id>
    Delete a specific teammembership

  user list
    List all users

  user get <user-id>
    Get a specific user

  useralias list
    List all users

  useralias get <useralias-id>
    Get a specific useralias

  sync team [<flags>] <team-definition-file>
    Team sync command

  version
    Display version information.
```

## Demo

![Animated Image of Terminal](https://github.com/cisco-sso/gitprime-cli/raw/master/demo/demo.gif)

## Installing

```
go get github.com/cisco-sso/gitprime-cli
```


## Building

```
make all
```


## Running

```
# Set your authentication token and endpoint
export GITPRIME_AUTH_TOKEN=<secret_token>
export GITPRIME_SERVER_HOST=<https://your-gitprime-instance>


# Execute a command
gitprime-cli team list
```

## Debugging

```
# Enable debugging via environment variable
export DEBUG=1

# Execute the command with a flag
gitprime-cli --log-level=debug ...
```

## Team Syncing from File

This tool may be used to sync a file which contains a mapping of team names to users.

In GitPrime, to create a team and membership structure that looks like this:

```
org/
└── business-unit
    ├── team1
    │   ├── user1@domain.com
    │   ├── user2@domain.com
    │   └── user3@domain.com
    └── team2
        ├── user4@domain.com
        └── user5@domain.com
```

Create a YAML file like this:

```
# teams.yaml
#
# This is a list of teams
#   Each team has a team name.
#   Each team has a list of members identified by a list of email addresses.
#   Team names support parent child relationships specified with dot separation.
#   The tool will skip members listed here but not encounted by GitPrime's code scanning.
- name: org.business-unit.team1
  members:
    - user1@domain.com
    - user2@domain.com
    - user3@domain.com
- name: org.business-unit.team2
  members:
    - user4@domain.com
    - user5@domain.com
```

Then run:

```
gitprime-cli sync team teams.yaml --log-level=trace
```

Re-running the tool is safe.  If new users or teams are added to the teams.json file, these entities will be added when the tool is run again.  However, removing users or teams from the teams.json file will not remove these entities when subsequently re-running the tool.  This may be worked around by manually deleting the top level teams from git prime (in this case `org`) which will recursively delete all teams and memberships and create a fresh empty state, and then re-running the tool again to sync all state as new.


## How this project was created

* Manually downloaded the GitPrime (Flow) swagger file from `https://flow.pluralsight.com/v3/customer/core/docs/?format=openapi` (requires manual download via Firefox web inspector because of auth issues).
  * The swagger file is not published here because Pluralsight will not allow it.
  * Thus complete end-to-end code generation is not available unless you download the swagger file, patch it with `./api/swagger/patch.txt`, and save the result as `./api/swagger/swagger.patched.json`.  After doing this, `make all` will work.
* Manually modified the swagger file with additional details to facilite code generation.
* Ran [go-swagger](https://github.com/go-swagger/go-swagger) against GitPrime (Flow) swagger file to generate Golang API client libraries.
* Wrote code to create CLI client using the [Kingpin](https://github.com/alecthomas/kingpin) library by connecting CLI actions to Gitprime API calls
