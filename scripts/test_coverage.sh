#!/bin/bash

go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
rm -rf coverage.out
