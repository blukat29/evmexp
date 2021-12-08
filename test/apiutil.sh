#!/bin/bash

echo "{\"format\":\"evm_generic\", \"binary\":\"$(cat vectors/eth_usdt.txt)\"}" \
    | curl -v -H "Content-Type: application/json" -X POST --data @- \
      http://localhost:8000/api/code/upload

curl -v -o out.json http://localhost:8000/api/deco/evm_generic-6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb
