#!/bin/bash

curl -s  https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -r --arg hero_id "$HERO_ID" ' .[] | select(.id == ($hero_id | tonumber)) | .connections.relatives | sub("\n";"\\n")'