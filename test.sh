#!/usr/bin/env sh

set -euo pipefail

docker build -f build/Dockerfile.test -t web-api-tester-test .

docker rmi web-api-tester-test
