#!/bin/bash

if [ ! -d bin ]; then
  echo mkdir bin
fi

cd ./LoadBalancer/WebServerBalancer/

go build -o "../../bin/BalancerWebServer" ./BalancerWebServer.go

./../../bin/BalancerWebServer
