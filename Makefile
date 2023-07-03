#!make
include .env

cmgr:
	migrate create -ext sql -dir db/migrations -seq ${name}

migup:
	migrate -path db/migrations -database "${MYSQL_DBCONN}" -verbose up

migdown:
	migrate -path db/migrations -database "${MYSQL_DBCONN}" -verbose down

migupx:
	migrate -path db/migrations -database "${MYSQL_DBCONN}" -verbose up ${num}

migdownx:
	migrate -path db/migrations -database "${MYSQL_DBCONN}" -verbose down ${num}

setupair:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

run:
	bin/air
