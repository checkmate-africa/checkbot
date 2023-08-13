.PHONY: build

build:
	sam build
api:
	sam local start-api
start:
	sam local start-lambda