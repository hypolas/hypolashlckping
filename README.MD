# Description

healthcheck for ping. This module is a part of:

[Universal healthcheck](https://github.com/hypolas/healthcheck/releases)


# Configuration

Export this environnement variables:

| Variable | Default value |  Description|   Required |   Set on | 
|---|---|---|---|---|
| `HYPOLAS_HEALTHCHECK_TYPE`  | *""*  | Type of healthcheck  | **X** | **ping** |
| `HYPOLAS_HEALTHCHECK_PING_COUNT`  | *4*  | How many ping will be send  |
| `HYPOLAS_HEALTHCHECK_PING_HOST`  | *""*  | Hostname/IP of target   | **X**
| `HYPOLAS_HEALTHCHECK_PING_MAX_PERCENT_LOST`  | *0*  | Max percent packet lost  | 
| `HYPOLAS_LOGS_FILE`  | *""*  | Path of logs file  | **X**

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