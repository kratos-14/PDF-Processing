#!/bin/bash
build=$1
load=$2
DIR_PATH=$(pwd)

echo $DIR_PATH
if [ "$build" = "build" ]; then
    docker build -t suhail12/producer-service:latest $DIR_PATH/producer-service/
    docker build -t suhail12/frontend-service:latest $DIR_PATH/frontend-service/
    docker build -t suhail12/compressor-service:latest $DIR_PATH/compressor-service/
    docker build -t suhail12/dbclean-service:latest $DIR_PATH/DBClean-service/
fi

if [ "$load" = "load" ]; then
    kind load docker-image suhail12/frontend-service:latest -n kubernetes
    kind load docker-image suhail12/producer-service:latest -n kubernetes
    kind load docker-image suhail12/compressor-service:latest -n kubernetes
    kind load docker-image suhail12/dbclean-service:latest -n kubernetes
fi