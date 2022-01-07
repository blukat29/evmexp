#!/bin/bash

set -e

mkdir -p data

docker run --rm -p 8000:8000 \
    --user $(id -u):$(id -g) \
    -v /etc/passwd:/etc/passwd:ro \
    -v /etc/group:/etc/group:ro \
    -v $(pwd)/data:/app/data \
    -e DB_PATH="/app/data/data.db" \
    -e KAS_KEY_ID \
    -e KAS_SECRET \
    -it evmexp "$@"
