# gitprime-cli

Golang-based CLI for making API calls against Plauralsight's GitPrime (Flow) product.

Also contains subcommands for syncing users and teams from a CSV spreadsheet to the GitPrime (Flow) product.

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

  team get <team-id>
    Get a specific team

  teammembership list
    List all teams

  teammembership get <teammembership-id>
    Get a specific teammembership

  user list
    List all users

  user get <user-id>
    Get a specific user

  useralias list
    List all users

  useralias get <useralias-id>
    Get a specific useralias

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

## How this project was created

* Manually downloaded the GitPrime (Flow) swagger file from `https://flow.pluralsight.com/v3/customer/core/docs/?format=openapi` (requires manual download via Firefox web inspector because of auth issues).
* Manually modified the swagger file with additional details to facilite code generation.
* Ran [go-swagger](https://github.com/go-swagger/go-swagger) against GitPrime (Flow) swagger file to generate Golang API client libraries.
* Wrote code to create CLI client using the [Kingpin](https://github.com/alecthomas/kingpin) library by connecting CLI actions to Gitprime API calls
