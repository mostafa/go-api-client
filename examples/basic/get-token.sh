#!/bin/bash

read -e -p "Logto Endpoint: " -i "http://localhost:3001/" logtoEndpoint
read -e -p "App ID: " appId
read -e -p "App Secret: " appSecret

basicAuthString=$(echo -n "$appId:$appSecret" | base64 -w 0 -)

curl -sS --location \
    --request POST $logtoEndpoint'oidc/token' \
    --header 'Authorization: Basic '$basicAuthString \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'grant_type=client_credentials' \
    --data-urlencode 'resource=https://default.logto.app/api' \
    --data-urlencode 'scope=all' -o token.json

echo "Token saved to token.json"
cat token.json
