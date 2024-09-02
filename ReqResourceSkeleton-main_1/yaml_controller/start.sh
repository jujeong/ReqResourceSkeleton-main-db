#!/bin/sh
export $(grep -v '^#' .env | xargs)

# Kill processes
fuser -k $KETI_SERVER_PORT/tcp
fuser -k $MYSQL_PORT/tcp
docker compose -f ./dockerCompose/docker-compose.yaml down
docker compose -f ./dockerCompose/docker-compose.yaml up
echo "start.sh :: Running Database server"