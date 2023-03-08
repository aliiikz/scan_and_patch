#!/bin/bash

while read p; do
    OUTPUT=$(grep -w 'fixed\|OWN\|IR' /home/ali/Projects/MyGithub/scan_and_patch_servers/results/$p\_short.txt | cut -d "|" -f 2,3,7,8,9,10 | awk '{print$1,$3,$7}' | column -t)
    curl --location --request POST 'https://api.telegram.org/[BOT]:[TOKEN]/sendMessage' \
        --header 'Content-Type: application/json' \
        --data-raw "{\"chat_id\": \"[ID]\", \"text\": \"ServerIP is: $p\nThe Result:\n$OUTPUT\", \"disable_notification\": true}"
    sleep 3
done </home/ali/Projects/MyGithub/scan_and_patch_servers/configs/hosts.txt