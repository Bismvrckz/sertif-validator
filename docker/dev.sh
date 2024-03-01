#!/bin/bash
docker volume create sv_mysql_keycloak_data
docker volume create sv_keycloak_data
docker volume create sv_mysql_data
docker volume create sv_go


docker compose -f dev.yml $1 $2 $3 --remove-orphans
