package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"log"
	"vysioneer-assignment/auth"
	"vysioneer-assignment/job"
	"vysioneer-assignment/model"
	"vysioneer-assignment/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type ViewFunc func(http.ResponseWriter, *http.Request)

func init() {
	gob.Register(&model.User{})

	//Check folders
	os.MkdirAll(os.Getenv("TMP_DIR"), os.ModePerm)
	os.MkdirAll(os.Getenv("VIDEO_DIR"), os.ModePerm)

	job.GetJobManager().Run()
}

func generalHandler(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Check UP")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Content-Type", "application/json; charset=UTF-8;")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			fmt.Println("Check OPTIONS")
			fmt.Fprintf(w, "ok")
			return
		}
		f(w, r)
		return
	}
}

func authHandler(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("ok"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Auth user
	user, err = auth.AuthUser(w, r)
	if err != nil {
		// w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session.Values["user"] = &user
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ok")
	return
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["user"] = nil

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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	us := services.GetUserService()
	err = us.CreateUser(user.Name, user.Email, user.Password, time.Now().String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	fmt.Fprintf(w, string("ok"))
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
	newName := r.FormValue("name")
	description := r.FormValue("description")
	perm, err := strconv.ParseBool(r.FormValue("perm"))

	// funcType := r.FormValue("type")
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	defer file.Close()

	videoID := uuid.NewV4().String()
	fmt.Println("Uploaded " + videoID)
	f, err := os.OpenFile(os.Getenv("TMP_DIR")+videoID+".mp4", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Write tmp file failed!"))
		return
	}

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
		w.Write([]byte("500 - User error!"))
		return
	}

	vs := services.GetVideoService()
	fmt.Printf("filesize %d\n", fi.Size())
	fmt.Printf("status %s\n", model.VideoStatusUpload)
	err = vs.CreateVideo(videoID, model.VideoStatusUpload, user.ID, newName, description, perm, fi.Size(), time.Now().String())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Create video failed."))
		return
	}

	// Add job
	video, err := vs.GetVideo(videoID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Get video failed."))
		return
	}

	jm := job.GetJobManager()
	err = jm.AddFaceBlurVideoJob(video)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Add video job failed."))
		return
	}

	fmt.Fprintf(w, string("ok"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoID := vars["id"]

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
	video, err := vs.GetVideo(videoID)
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

	Openfile, err := os.Open(os.Getenv("VIDEO_DIR") + videoID + ".mp4")
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

	switch r.Method {
	case "GET":
		jsonBytes, err := json.Marshal(video)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		fmt.Fprintf(w, string(jsonBytes))
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}

		paramVideo := model.Video{}
		err = json.Unmarshal(body, &paramVideo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Video unmarshal failed.!"))
			return
		}

		video.Perm = paramVideo.Perm
		vs := services.GetVideoService()
		err = vs.UpdateVideo(video)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Update video failed."))
			return
		}
		fmt.Fprintf(w, "")

	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
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
		fmt.Println(err)
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

// func processVideoHandler(w http.ResponseWriter, r *http.Request) {
// 	if job.InUse {
// 		fmt.Fprintf(w, "Video processing is in used.")
// 		return
// 	}

// 	r.ParseMultipartForm(32 << 20)
// 	newName := r.FormValue("name")
// 	description := r.FormValue("description")

// 	// funcType := r.FormValue("type")
// 	file, _, err := r.FormFile("file")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("500 - Something bad happened!"))
// 		return
// 	}
// 	defer file.Close()

// 	videoID := uuid.NewV4().String()
// 	fmt.Println("Processing " + videoID)
// 	f, err := os.OpenFile(os.Getenv("TMP_DIR")+videoID+".mp4", os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("500 - Write tmp file failed!"))
// 		return
// 	}

// 	io.Copy(f, file)
// 	fi, _ := f.Stat()

// 	// Processing video
// 	err = job.FaceBlurHandler(videoID)
// 	if err == job.ErrorOnUsed {
// 		fmt.Fprintf(w, string("Video processing is in used."))
// 		return
// 	} else if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("500 - Face blur failed"))
// 		return
// 	}

// 	//Record video meta
// 	session, err := store.Get(r, "vysioneer-assignment")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	val := session.Values["user"]
// 	var user = &model.User{}
// 	var ok bool
// 	if user, ok = val.(*model.User); !ok {
// 		// Handle the case that it's not an expected type
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("500 - User error!"))
// 		return
// 	}

// 	vs := services.GetVideoService()
// 	fmt.Printf("filesize %d\n", fi.Size())
// 	err = vs.CreateVideo(videoID, user.ID, newName, description, fi.Size(), time.Now().String())
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("500 - Create video failed."))
// 		return
// 	}

// 	fmt.Fprintf(w, string("ok"))
// }

func browseVideoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoID := params["id"]

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
	video, err := vs.GetVideo(videoID)
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

	http.ServeFile(w, r, os.Getenv("VIDEO_DIR")+videoID+".mp4")
}

func sharelinkHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vs := services.GetVideoService()
	video, _ := vs.GetVideo(id)

	if video.Perm {
		Openfile, err := os.Open(os.Getenv("VIDEO_DIR") + id + ".mp4")
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

	http.Error(w, "File not found.", 404)
	return
}

func httpStart() {

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./dist"))
	r.Handle("/web/", http.StripPrefix("/web/", fs))
	r.HandleFunc("/health", healthHandler)

	//User
	r.HandleFunc("/user", authHandler(userHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/user-service/register", registerHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/user-service/login", loginHandler)
	r.HandleFunc("/user-service/logout", logoutHandler).Methods("POST", "OPTIONS")

	//Video
	r.HandleFunc("/video/{id}", authHandler(videoHandler)).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/videos", authHandler(videosHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/video-service/upload", authHandler(uploadHandler)).Methods("POST", "OPTIONS")
	r.HandleFunc("/video-service/download/{id}", authHandler(downloadHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/video-service/browse/{id}", authHandler(browseVideoHandler)).Methods("GET", "OPTIONS")
	// r.HandleFunc("/video-service/process", authHandler(processVideoHandler)).Methods("POST", "OPTIONS")

	//Sharelink
	r.HandleFunc("/sharelink/{id}", sharelinkHandler).Methods("GET", "OPTIONS")

	headers := handlers.AllowedHeaders([]string{"Set-Cookie", "X-Requested-With", "Content-Type", "authorization", "access-control-allow-origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"})
	origins := handlers.AllowedOrigins([]string{"https://localhost:8080", "https://localhost:9000", "https://web.ezsofa.com", "https://web.ezsofa.com"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	cred := handlers.AllowCredentials()
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), handlers.CORS(origins, headers, methods, cred)(r)))
}

func main() {
	fmt.Println("VYSIONEER assignment running...")
	httpStart()
}
