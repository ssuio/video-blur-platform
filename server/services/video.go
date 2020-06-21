package services

import (
	"video-processing/job"
	"video-processing/model"
)

var videoService *VideoService

type VideoService struct {
	dataManager model.DataProvider
	jobManager  job.TaskManager
}

func (vs VideoService) UploadVideo() {}

func (vs VideoService) CreateVideo(id string, status string, ownerID int, name string, description string, perm bool, size int64, createdTime string) error {
	return vs.dataManager.CreateVideo(id, status, ownerID, name, description, perm, size, createdTime)
}

func (vs VideoService) UpdateVideo(video model.Video) error {
	return vs.dataManager.UpdateVideo(video)
}

func (vs VideoService) DeleteVideo(video string) error {
	return vs.dataManager.DeleteVideo(video)
}

func (vs VideoService) GetVideo(id string) (model.Video, error) {
	return vs.dataManager.GetVideo(id)
}

func (vs VideoService) ListVideos(ownerID int) ([]model.Video, error) {
	return vs.dataManager.ListVideos(ownerID)
}

func (vs VideoService) ProcessVideo(plan string) {

}

func (vs VideoService) GenerateSharelink(videoID string) {

}

func (vs VideoService) DownloadVideo(videoID string) {
	//Check permission

	//Download

}

func GetVideoService() *VideoService {
	if videoService == nil {
		videoService = &VideoService{
			model.GetDataManager(),
			job.TaskManager{},
		}
	}
	return videoService
}
