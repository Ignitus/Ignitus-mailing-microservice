.PHONY: run_prod

dev:
	clear && go run main.go

build:
	clear && go build

run_prod: build
	GIN_MODE=release ./ignitus-mailing-microservice