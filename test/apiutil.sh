#!/bin/bash

echo "{\"code\":\"$(cat vectors/eth_usdt.txt)\"}" \
    | curl -v -H "Content-Type: application/json" -X POST --data @- \
      http://localhost:8080/api/deco -o out.json
