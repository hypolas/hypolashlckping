# Description

healthcheck for ping.

# Configuration

Export this environnement variables:

- HYPOLAS_HEALTHCHECK_TYPE=ping
- HYPOLAS_HEALTHCHECK_PING_COUNT=4
- HYPOLAS_HEALTHCHECK_PING_HOST=www.google.fr
- HYPOLAS_HEALTHCHECK_PING_MAX_PERCENT_LOST=0
- HYPOLAS_LOGS_FILE=test/logs.log

# Default value

| Variable | Default value |  Description|   Required |
|---|---|---|---|
| `HYPOLAS_HEALTHCHECK_TYPE`  | *"ping"*  | Type of healthcheck  | false
| `HYPOLAS_HEALTHCHECK_PING_COUNT`  | *4*  | How many ping will be send  | false
| `HYPOLAS_HEALTHCHECK_PING_HOST`  | *""*  | Hostname/IP of target   | **true**
| `HYPOLAS_HEALTHCHECK_PING_MAX_PERCENT_LOST`  | *0*  | Max percent packet lost  | false
| `HYPOLAS_LOGS_FILE`  | *""*  | Path of logs file  | **true**

# Exemple

## Dockerfile
```Dockerfile
FROM myimage:1.0.0

ENV HYPOLAS_HEALTHCHECK_TYPE="ping"
ENV HYPOLAS_HEALTHCHECK_PING_COUNT="5"
ENV HYPOLAS_HEALTHCHECK_PING_HOST="wantpingup.lan"
ENV HYPOLAS_LOGS_FILE="log/logs.log"

RUN xxxxxx

HEALTHCHECK --interval=1m --timeout=30s --retries=3 CMD healthcheck
```

## docker-compose.yml

```yaml
version: '3'

services:
  app:
    image: myimage:1.0.0
    environment:
      HYPOLAS_HEALTHCHECK_TYPE: "ping"
      HYPOLAS_HEALTHCHECK_PING_COUNT: 5
      HYPOLAS_HEALTHCHECK_PING_HOST: wantpingup.lan
      HYPOLAS_LOGS_FILE: log/logs.log
    healthcheck:
      test: ["CMD", "healthcheck"]
      timeout: 30s
      interval: 1m
      retries: 3
```

## Directly on Linux

```bash
...launcher.sh
#!/usr/bin/env bash

export HYPOLAS_HEALTHCHECK_TYPE="ping"
export HYPOLAS_HEALTHCHECK_PING_COUNT="5"
export HYPOLAS_HEALTHCHECK_PING_HOST="wantpingup.lan"
export HYPOLAS_LOGS_FILE="log/logs.log"
./healthcheck
```

## Directly on Windows

```bash
...launcher.bat
set HYPOLAS_HEALTHCHECK_TYPE="ping"
set HYPOLAS_HEALTHCHECK_PING_COUNT="5"
set HYPOLAS_HEALTHCHECK_PING_HOST="wantpingup.lan"
set HYPOLAS_LOGS_FILE="log/logs.log"
./healthcheck.exe
```