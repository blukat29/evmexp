#!/bin/bash

echo "{\"code\":\"$(cat vectors/eth_usdt.txt)\"}" | curl -v -H "Content-Type: application/json" -X POST --data @- http://localhost:7000/deco -o out.json
