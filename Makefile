.PHONY: build

include lambda.mk

build:
	sam build --use-container
start-api:
	sam local start-api --env-vars env.json
start-lambda:
	sam local start-lambda --env-vars env.json