#!/usr/bin/env bash

# Clean up from the prior run
SILENCED=`killall -9 gitprime-cli 2>&1 > /dev/null`

# Do some checks
if [[ -z ${GITPRIME_AUTH_TOKEN:-""} ]]; then
    echo "Please set environment variable for GITPRIME_AUTH_TOKEN"
    exit 0
fi
if [[ -z ${GITPRIME_SERVER_HOST:-""} ]]; then
    echo "Please set environment variable for GITPRIME_SERVER_HOST"
    exit 0
fi

tmux new-session -n goapi -s test "tmux set-option -t test status off; yes | GITPRIME_AUTH_TOKEN=${GITPRIME_AUTH_TOKEN} GITPRIME_SERVER_HOST=${GITPRIME_SERVER_HOST} ./demo/demo_client.sh; tmux kill-session -t test"
