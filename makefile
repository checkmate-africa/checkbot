.PHONY: build

include lambda.mk

build:
	sam build --use-container
api:
	sam local start-api --env-vars env.json
start:
	sam local start-lambda --env-vars env.json