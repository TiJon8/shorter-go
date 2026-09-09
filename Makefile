include .env.local
export

export PROJECT_ROOT=$(shell pwd)

run:
	@go run ${PROJECT_ROOT}/cmd/shorter/main.go -use-env