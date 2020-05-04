package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	CreatedTime time.Time
}

type Video struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Size        int64     `json: "size"`
	CreatedTime time.Time `json: "createdTime"`
	ImageUrl    string
	OwnerID     int
	Perm        Perm
}

type PermRule struct {
}

type Perm struct {
	ownerID string
	rules   []PermRule
}

type DataProvider interface {
	GetUser(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(name, email, password, createdTime string) error
	GetVideo(id string) (Video, error)
	ListVideos(ownerID int) ([]Video, error)
	CreateVideo(id string, ownerID int, name string, description string, size int64, createdTime string) error
}

type SqliteProvider struct{}

func (s SqliteProvider) GetUser(id int) (User, error) {
	var user User
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return user, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, name, email, password, created_time from user where id = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	user = User{}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedTime)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s SqliteProvider) GetUserByEmail(email string) (User, error) {

	var user User
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return user, err
	}
	defer db.Close()

	fmt.Println(email)
	stmt, err := db.Prepare("select id, name, email, password, created_time from user where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	user = User{}
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedTime)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s SqliteProvider) CreateUser(name, email, password, createdTime string) error {
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user (name, email, password, created_time ) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, password, createdTime)
	if err != nil {
		return err
	}

	return nil
}

func (s SqliteProvider) GetVideo(id string) (Video, error) {
	var video Video
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return video, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, owner_id, name, description, size, created_time from video where id = ?")
	if err != nil {
		return video, err
	}
	defer stmt.Close()

	video = Video{}
	err = stmt.QueryRow(id).Scan(&video.ID, &video.OwnerID, &video.Name, &video.Description, &video.Size, &video.CreatedTime)
	if err != nil {
		return video, err
	}
	return video, nil
}

func (s SqliteProvider) ListVideos(ownerID int) ([]Video, error) {
	list := make([]Video, 0)
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return list, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, owner_id, name, description, size, created_time FROM video WHERE owner_id=?", ownerID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		video := Video{}
		if err := rows.Scan(&video.ID, &video.OwnerID, &video.Name, &video.Description, &video.Size, &video.CreatedTime); err != nil {
			return list, err
		}
		list = append(list, video)
	}

	return list, nil
}

func (s SqliteProvider) CreateVideo(id string, ownerID int, name string, description string, size int64, createdTime string) error {
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_FILE"))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO video (id, owner_id, name, description, size, created_time ) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, ownerID, name, description, size, createdTime)
	if err != nil {
		return err
	}
	return nil
}

func GetDataManager() DataProvider {
	p := SqliteProvider{}
	return p
}
