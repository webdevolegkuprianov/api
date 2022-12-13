all: local-up

#build:
#	go build -v ./cmd/app

#start:
#	./app

local-up:
	docker-compose -f deploy/docker-compose.yml up --build -d