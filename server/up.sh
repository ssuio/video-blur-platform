#!/bin/bash
export PORT=9000 
export HOST=127.0.0.1
export HOST_DATA_DIR=$HOME/workspace/vysioneer-assignment/server/data
export DATA_DIR=$HOME/workspace/vysioneer-assignment/server/data
export SQLITE_FILE=$HOME/workspace/vysioneer-assignment/server/data/vysioneer.db
export VIDEO_DIR=$HOME/workspace/vysioneer-assignment/server/data/videos/
export TMP_DIR=/$HOME/workspace/vysioneer-assignment/server/data/tmp/
export SESSION_KEY=ssuio
go run main.go