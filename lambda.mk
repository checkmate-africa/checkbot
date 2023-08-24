build-EventGatewayFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o event-gateway lambda/event/gateway/main.go
	mv ./event-gateway $(ARTIFACTS_DIR)/

build-EventTaskFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o event-task lambda/event/task/main.go
	mv ./event-task $(ARTIFACTS_DIR)/

build-InteractionGatewayFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o interaction-gateway lambda/interaction/gateway/main.go
	mv ./interaction-gateway $(ARTIFACTS_DIR)/

build-InteractionTaskFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o interaction-task lambda/interaction/task/main.go
	mv ./interaction-task $(ARTIFACTS_DIR)/

build-ShufflerFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o shuffler lambda/shuffler/main.go
	mv ./shuffler $(ARTIFACTS_DIR)/