#!/bin/bash

docker build --no-cache -t aenocmartinez/threads-container:qa -f Dockerfile.prod . --platform linux/amd64

docker push aenocmartinez/threads-container:qa

docker rmi aenocmartinez/threads-container:qa