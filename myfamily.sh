#!/bin/bash

# Check if HERO_ID is set
if [ -z "$HERO_ID" ]; then
    echo "Please set the HERO_ID environment variable."
    exit 1
fi

# Fetch superhero data and print relatives in one line
curl -s "https://01.gritlab.ax/assets/superhero/all.json" | \
jq -r --arg id "$HERO_ID" '
  .[] 
  | select(.id == ($id | tonumber)) 
  | (.connections.relatives // "") 
  | gsub("\n"; ", ")
'