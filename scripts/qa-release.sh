#!/bin/bash

docker build --no-cache -t pulzo/threads-container:qa -f Dockerfile.prod . --platform linux/amd64

docker push pulzo/threads-container:qa

docker rmi pulzo/threads-container:qa