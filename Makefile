BIN_OUTPUT := $(if $(filter $(shell go env GOOS), windows), _output/http-server.exe, _output/http-server)

default: build

clean:
	rm _output/*

build:
	go build -o $(BIN_OUTPUT)