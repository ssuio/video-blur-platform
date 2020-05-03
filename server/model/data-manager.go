package model

import (
	"database/sql"
	"time"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"os"

)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string 
	CreatedTime time.Time
}

type Video struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Perm Perm `json: "perm"`
	Size int
	LengthInSecond int
	Date time.Time
}

type PermRule struct {

}

type Perm struct {
	ownerID string
	rules []PermRule
}

type DataProvider interface {
	GetUser(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser() (User, error)
	GetVideo() (Video, error)
	GetVideos() ([]Video, error)
}

type SqlliteProvider struct {

}

func (s SqlliteProvider) GetUser(id string) (User, error){
	var user User
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, name, email, password, created_time from user where id = ?")
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stmt.Close()
	user = User{} 
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedTime)
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}

func (s SqlliteProvider) GetUserByEmail(email string) (User, error){
	var user User
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer db.Close()
	
	stmt, err := db.Prepare("select id, name, email, password, created_time from user where email = ?")
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stmt.Close()
	user = User{}
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedTime)
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}

func (s SqlliteProvider) CreateUser() (User, error){
	t, _ := time.Parse("2006-01-02", "2020-05-01")
	return User{
		"111",
		"noah",
		"noah.chou@eztable.com",
		"12345678",
		t,
	}, nil
}

func (s SqlliteProvider) GetVideo() (Video, error){
	return Video{

	}, nil
}

func (s SqlliteProvider) GetVideos() ([]Video, error){
	return []Video{}, nil
}

func GetDataManager() DataProvider{
	p := SqlliteProvider{}
	return p
}