#!/usr/bin/env bash

# Load the common demo lib
demo_dir="$(dirname "$0")"
. "$demo_dir/demo_magic_common.sh"
. "$demo_dir/common.sh"

#####################################################################
# Start the good stuff

SLEEPTIME=2

pe "# Show CLI usage."
pe "${CLI} --help"
sleep ${SLEEPTIME}

pe "# Show CLI version."
pe "${CLI} version"
sleep ${SLEEPTIME}

pe "# Export your GITPRIME_AUTH_TOKEN as an environment variable."
p  "export GITPRIME_AUTH_TOKEN=<secret_token>"
sleep 1

pe "# Export your GITPRIME_SERVER_HOST as an environment variable."
p  "export GITPRIME_SERVER_HOST=flow.pluralsight.com  # This is the default"
sleep 1

FILTERS=$(cat <<EOF
sed 's@https://.*/v3@https://flow.pluralsight.com/v3@g' \
| jq '(..|objects|select(has("name"))).name |= "REDACTED"' \
  | jq '(..|objects|select(has("email"))).email |= "REDACTED"' \
  | jq '(..|objects|select(has("login_email"))).login_email |= "REDACTED"' \
  | jq '(..|objects|select(has("login_preferred_name"))).login_preferred_name |= "REDACTED"' \
  | jq '(..|objects|select(has("ancestors"))).ancestors |= "REDACTED"' \
  | head -n 25
EOF
)

pe "# List all users"
p  "${CLI} user list | head -n 25"
echo "${CLI} user list | $FILTERS" | sh
sleep ${SLEEPTIME}

pe "# Get a specific user with user-id=1"
p  "${CLI} user get 1 | head -n 25"
echo "${CLI} user get 1 | $FILTERS" | sh
sleep ${SLEEPTIME}

pe "# List all teams"
p "${CLI} team list | head -n 25"
echo "${CLI} team list | $FILTERS" | sh
sleep ${SLEEPTIME}

pe "# Get a specific team with team-id=1"
p  "${CLI} team get 1 | head -n 25"
echo "${CLI} team get 1 | $FILTERS" | sh
sleep ${SLEEPTIME}

pe "# That's it for the demo.  Thanks for watching!"
sleep ${SLEEPTIME}

clear
