#!/bin/sh
mkdir -p /tkbai-dashboard/src/tmp/build
ln -s /tkbai-dashboard/src /usr/local/go/src/tkbai-dashboard
cd /tkbai-dashboard/src
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest&
go install github.com/cosmtrek/air@latest&
migrate -database "mysql://root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/tkbai" -path /tkbai-dashboard/migration up
while true; do sleep 1; done