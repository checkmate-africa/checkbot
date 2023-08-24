#!/bin/bash

export GOFLAGS="-ldflags=-buildvcs=false"
sam build --use-container
