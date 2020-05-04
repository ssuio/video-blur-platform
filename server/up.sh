#!/bin/bash
export PORT=9000 
export HOST=127.0.0.1
export DATA_DIR=/Users/noahchou/Workspace/vysioneer-assignment/server/data
export SQLITE_FILE=/Users/noahchou/Workspace/vysioneer-assignment/server/data/vysioneer.db
export VIDEO_DIR=/Users/noahchou/Workspace/vysioneer-assignment/server/data/videos/
export SESSION_KEY=ssuio
go run main.go