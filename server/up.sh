#!/bin/bash
export PORT=9000 
export DATA_DIR=/home/ssuio/workspace/vysioneer-assignment/server/data
export SQLITE_FILE=/home/ssuio/workspace/vysioneer-assignment/server/data/vysioneer.db
export SESSION_KEY=ssuio
go run main.go