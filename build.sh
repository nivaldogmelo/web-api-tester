#!/usr/bin/env sh

set -euo pipefail

docker build -f build/Dockerfile.build -t web-api-tester .
