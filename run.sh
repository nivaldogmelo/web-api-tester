#!/usr/bin/env sh

set -eo pipefail

if [[ -z $1 ]]; then
    echo "usage: ./run.sh <config-file> <database-file>"
    exit 1
fi

if [[ -z $2 ]]; then
    echo "usage: ./run.sh <config-file> <database-file>"
    exit 1
fi

docker build -f build/RunDockerfile -t web-api-tester .

clear

docker run --rm -p 3000:3000 -v "$PWD/$1":/config.yaml -v "$PWD/$2:/$2" web-api-tester
