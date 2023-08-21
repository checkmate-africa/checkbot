.PHONY: build

build:
	sam build
api:
	sam local start-api --env-vars env.json
start:
	sam local start-lambda --env-vars env.json