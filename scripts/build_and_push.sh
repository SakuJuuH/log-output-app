#!/bin/zsh

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <tag>"
  exit 1
fi

TAG=$1


docker build --platform=linux/amd64 -t sakuheinonen/log-output:${TAG} ./log-output
docker build --platform=linux/amd64 -t sakuheinonen/ping-pong:${TAG} ./ping-pong
docker build --platform=linux/amd64 -t sakuheinonen/greeter:${TAG} ./greeter

docker push sakuheinonen/log-output:${TAG}
docker push sakuheinonen/ping-pong:${TAG}
docker push sakuheinonen/greeter:${TAG}
