#!/bin/bash
# go-ws: webserver implementation in go

AIR:=air

bench:
	$(AIR) bench

run: 
	$(AIR) server --port 8080
