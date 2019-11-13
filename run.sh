#!/usr/bin/env bash

docker run --rm -d -p 3306:3306 --env-file .env mysql
echo "Waiting for Dabase conatainer to be ready"
sleep 30
GO111MODULE=on go run main.go