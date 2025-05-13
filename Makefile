#!/bin/bash
# go-ws: webserver implementation in go

go-ws:
	go build -o bin/$@ main.go

run: go-ws
	bin/$<
