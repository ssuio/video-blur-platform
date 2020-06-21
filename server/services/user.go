package services

import (
	"video-processing/job"
	"video-processing/model"
)

var userService *UserService = &UserService{
	model.GetDataManager(),
	job.TaskManager{},
}

type UserService struct {
	dataManager model.DataProvider
	taskManager job.TaskManager
}

func (us UserService) CreateUser(name, email, password, createdTime string) error {
	return us.dataManager.CreateUser(name, email, password, createdTime)
}

func (us UserService) GetUser(id int) (model.User, error) {
	return us.dataManager.GetUser(id)
}

func (us UserService) GetUserByEmail(email string) (model.User, error) {
	return us.dataManager.GetUserByEmail(email)
}

func (us UserService) Register() error {
	return nil
}

func GetUserService() *UserService {
	return userService
}
