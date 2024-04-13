#!/bin/bash
docker volume create tkbai_mysql_keycloak_data
docker volume create tkbai_keycloak_data
docker volume create tkbai_mysql_data
docker volume create tkbai_go


docker compose -f dev.yml $1 $2 $3 --remove-orphans
