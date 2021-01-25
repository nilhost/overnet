#!/usr/bin/env bash
case "$1" in

client)
    ./overnet --config="./config/config-client.json"
;;

client-prod)
    ./overnet --config="./config/config-client-prod.json"
;;

server)
    ./overnet --config="./config/config-server.json"
;;

build-mac)
    cd ../ && make macos && cd -
;;
*)
echo "usage: ./run.sh client/server"

esac
