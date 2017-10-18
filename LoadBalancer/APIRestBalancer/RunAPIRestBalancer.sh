#!/bin/bash

if [ ! -d bin ]; then
  echo mkdir bin
fi

cd ./LoadBalancer/APIRestBalancer/

go build -o "../../bin/BalancerAPIRest" ./BalancerAPIRest.go

./../../bin/BalancerAPIRest
