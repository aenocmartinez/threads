#!/bin/bash

docker build --no-cache -t pulzo/threads-container:latest -f Dockerfile.prod . --platform linux/amd64

docker push aenocmartinez/threads-container:latest

docker rmi aenocmartinez/threads-container:latest