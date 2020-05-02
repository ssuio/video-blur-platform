package model

import (
	"time"
)

type User struct {
	id string `json:"id"`
	name string `json:"name"`
	email string `json:"email"`
	password string 
}

type Video struct {
	id string `json:"id"`
	name string `json:"name"`
	perm Perm `json: "perm"`
	size int
	lengthInSecond int
	date time
}

type PermRule struct {

}

type Perm struct {
	ownerID string
	rules []PermRule
}

type DataProvider interface {
	GetUser() User
	CreateUser()(User, error)
	GetVideo() Video(User, error)
	GetVideos() ([]Video, error)
}

type SqlliteProvider struct {

}

func (s SqlliteProvider) GetUser(){

}

func (s SqlliteProvider) CreateUser(){

}

func (s SqlliteProvider) GetVideo(){

}

func (s SqlliteProvider) GetVideos(){

}

func GetDataManager() DataProvider{
	p := &SqlliteProvider{}
	return p
}