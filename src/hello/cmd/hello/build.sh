#!/bin/sh
echo "$TAG"
CGO_ENABLED=0 go build -v
docker build -t ${TAG} .
