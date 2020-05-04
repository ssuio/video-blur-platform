package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"vysioneer-assignment/auth"
	"vysioneer-assignment/job"
	"vysioneer-assignment/model"
	"vysioneer-assignment/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type ViewFunc func(http.ResponseWriter, *http.Request)

func init() {
	gob.Register(&model.User{})
}

func generalHandler(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
		return
	}
}

func authHandler(f ViewFunc) ViewFunc {
	return generalHandler(func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		session, err := store.Get(r, "vysioneer-assignment")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Check session
		val := session.Values["user"]
		if val != nil {
			f(w, r)
			return
		}

		// Auth user
		user, err = auth.AuthUser(w, r)
		fmt.Println(user)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			session.Values["user"] = &user
			f(w, r)
			return
		}
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("ok"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("logout")
	session.Values["user"] = &model.User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string("ok"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("failed1")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("failed")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	us := services.GetUserService()
	err = us.CreateUser(user.Name, user.Email, user.Password, time.Now().String())
	if err != nil {
		fmt.Println("failed2")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	fmt.Fprintf(w, string("user"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["user"]
	var user = &model.User{}
	var ok bool
	if user, ok = val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	fmt.Println(user)
	us := services.GetUserService()
	u, _ := us.GetUser(user.ID)

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	} else {
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, fileHandler, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
	defer file.Close()

	videoID := uuid.NewV4().String()
	f, err := os.OpenFile(os.Getenv("VIDEO_DIR")+videoID+".mp4", os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	fi, _ := f.Stat()

	//Record video meta
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := session.Values["user"]
	var user = &model.User{}
	var ok bool
	if user, ok = val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	vs := services.GetVideoService()
	fmt.Printf("filesize %d\n", fi.Size())
	err = vs.CreateVideo(videoID, user.ID, fileHandler.Filename, "", fi.Size(), time.Now().String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}

	fmt.Fprintf(w, string("ok"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varID := vars["id"]

	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["user"]
	var user = &model.User{}
	var ok bool
	if user, ok = val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	//Check video meta and permission
	vs := services.GetVideoService()
	video, err := vs.GetVideo(varID)
	fmt.Println(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if user.ID != video.OwnerID {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
		return
	}

	Openfile, err := os.Open(os.Getenv("VIDEO_DIR") + varID + ".mp4")
	defer Openfile.Close()
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+video.Name)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)
	return
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vs := services.GetVideoService()
	video, _ := vs.GetVideo(id)

	jsonBytes, err := json.Marshal(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	} else {
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func videosHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["user"]
	var user = &model.User{}
	var ok bool
	if user, ok = val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	vs := services.GetVideoService()
	videos, err := vs.ListVideos(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	jsonBytes, err := json.Marshal(videos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	} else {
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func processVideoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ssuio In")
	r.ParseMultipartForm(32 << 20)
	newName := r.FormValue("name")
	description := r.FormValue("description")

	fmt.Println("ssuio ")
	fmt.Println("newName " + newName)
	fmt.Println("description " + description)
	// funcType := r.FormValue("type")
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	defer file.Close()

	videoID := uuid.NewV4().String()
	f, err := os.OpenFile(os.Getenv("TMP_DIR")+videoID+".mp4", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.Copy(f, file)
	fi, _ := f.Stat()

	//Processing video
	err = job.FaceBlurHandler(videoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	//Record video meta
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := session.Values["user"]
	var user = &model.User{}
	var ok bool
	if user, ok = val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	vs := services.GetVideoService()
	fmt.Printf("filesize %d\n", fi.Size())
	err = vs.CreateVideo(videoID, user.ID, newName, description, fi.Size(), time.Now().String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}

	fmt.Fprintf(w, string("ok"))
}

func httpStart() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./dist"))
	r.Handle("/web/", http.StripPrefix("/web/", fs))

	//User
	r.HandleFunc("/user", authHandler(userHandler)).Methods("GET")
	r.HandleFunc("/user-service/register", generalHandler(registerHandler))
	r.HandleFunc("/user-service/login", authHandler(loginHandler)).Methods("POST")
	r.HandleFunc("/user-service/logout", generalHandler(logoutHandler)).Methods("POST")

	//Video
	r.HandleFunc("/video", authHandler(videoHandler))
	r.HandleFunc("/video/{id}", authHandler(videoHandler)).Methods("GET")
	r.HandleFunc("/videos", authHandler(videosHandler)).Methods("GET")
	r.HandleFunc("/video-service/upload", authHandler(uploadHandler)).Methods("PUT")
	r.HandleFunc("/video-service/download/{id}", authHandler(downloadHandler)).Methods("GET")
	r.HandleFunc("/video-service/sharelink", authHandler(downloadHandler)).Methods("POST")
	r.HandleFunc("/video-service/process", authHandler(processVideoHandler)).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("VYSIONEER assignment running...")
	httpStart()
}
