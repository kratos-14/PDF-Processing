#!/bin/bash

DIR_PATH=$(pwd)

echo $DIR_PATH
docker build -t suhail12/producer-service:latest $DIR_PATH/producer-service/

docker build -t suhail12/frontend-service:latest $DIR_PATH/frontend-service/

docker build -t suhail12/compressor-service:latest $DIR_PATH/compressor-service/

docker build -t suhail12/dbclean-service:latest $DIR_PATH/DBClean-service/
