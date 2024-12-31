#!/bin/sh

# Read the JSON payload from stdin
payload=$(cat)

# Extract the key using jq
# key=$(echo "$payload" | jq -r '.[0].Key')

# Log the event
# echo "Change detected in key: $key" >>/consul/data/watch.log
echo "Change detected:" >>/consul/data/watch-key.log

# Log the full payload for reference (optional)
echo "Payload: $payload" >>/consul/data/watch-key.log
