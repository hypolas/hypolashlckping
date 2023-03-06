#!/usr/bin/env bash

export HYPOLAS_HEALTHCHECK_TYPE=ping
export HYPOLAS_HEALTHCHECK_PING_COUNT=4
export HYPOLAS_HEALTHCHECK_PING_HOST=www.google.fr
export HYPOLAS_LOGS_FILE=test/logs.log

go test .

export HYPOLAS_HEALTHCHECK_ID=MYID # Simulate healthcheck -id MYID
export HYPOLAS_HEALTHCHECK_MYID_TYPE=ping
export HYPOLAS_HEALTHCHECK_MYID_PING_COUNT=4
export HYPOLAS_HEALTHCHECK_MYID_PING_HOST=www.google.fr
export HYPOLAS_LOGS_FILE=test/myid.log

go test .
