#!/bin/sh

tgt="printbox.local"

cd src
GOOS=linux GOARCH=arm GOARM=7 go build -o ../start.arm7 .
cd ..
scp start.arm7 ${tgt}:
echo ""
echo ""
ssh ${tgt} "./start.arm7; echo "-------------"; cat /printbox/docker-compose.yml"
