package services

import (
	"vysioneer-assignment/model"
	"vysioneer-assignment/job"
	"time"
)

var userService *UserService = &UserService{
	model.GetDataManager(),
	job.TaskManager{},
}

type UserService struct{
	dataManager model.DataProvider
	taskManager job.TaskManager
}

func (us UserService) GetUser (id string) (model.User, error){
	return us.dataManager.GetUser(id)
}

func (us UserService) Register () error{
	return nil
}

func (us UserService) Login () (model.User, error){
	t, _ := time.Parse("2006-01-02", "2020-05-01")
	return model.User{
		"111",
		"noah",
		"noah.chou@eztable.com",
		"12345678",
		t,
	}, nil
}

func GetUserService() *UserService{
	return userService
}
