#!/bin/zsh
set -e

kubectl delete -k ./kubernetes || true