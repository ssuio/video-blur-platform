package services

import (
	"vysioneer-assignment/job"
)

var instance *VideoService

type VideoService struct {
	jobManager job.TaskManager
}

func (vs VideoService) UploadVideo(){

}

func (vs VideoService) ListVideos(){

}

func (vs VideoService) ProcessVideo (plan string){

}

func (vs VideoService) GenerateSharelink (videoID string){

}

func (vs VideoService) DownloadVideo (videoID string){

}

func GetVideoService () *VideoService{
	if instance == nil {
		instance = &VideoService{}
	}
	return instance
}