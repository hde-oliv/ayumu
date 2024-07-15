#!/usr/bin/env sh

git pull origin master

docker build -t ayumu-discord .

docker stop ayumu && docker rm ayumu

docker run -d --name=ayumu --restart=always --env-file=./.env ayumu-discord
