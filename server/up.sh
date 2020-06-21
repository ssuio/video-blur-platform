#!/bin/bash
export PORT=9000 
export HOST=0.0.0.0
export HOST_DATA_DIR=$HOME/workspace/video-processing/server/data
export DATA_DIR=$HOME/workspace/video-processing/server/data
export SQLITE_FILE=$HOME/workspace/video-processing/server/data/vp.db
export VIDEO_DIR=$HOME/workspace/video-processing/server/data/videos/
export TMP_DIR=/$HOME/workspace/video-processing/server/data/tmp/
export SESSION_KEY=ssuio
go run main.go