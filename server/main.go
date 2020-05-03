package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
	"io"
	"vysioneer-assignment/services"
	"vysioneer-assignment/auth"
	"vysioneer-assignment/model"
	"encoding/json"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
type ViewFunc func(http.ResponseWriter, *http.Request)

func generalHandler(f ViewFunc) ViewFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
		return
	}
}

func authHandler(f ViewFunc) ViewFunc{
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

func loginHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("user"))
	w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Something bad happened!"))
}

func registerHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("user"))
	w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Something bad happened!"))
}

func userHandler(w http.ResponseWriter, r *http.Request){
	session, err := store.Get(r, "vysioneer-assignment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["user"]
	var user = &model.User{}
    if _, ok := val.(*model.User); !ok {
		// Handle the case that it's not an expected type
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	
	us := services.GetUserService()
	u,_ := us.GetUser(user.ID)

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}else{
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func videoHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	  varID := vars["id"]
	  println(varID)
	// fmt.Fprintf(w, "video " + varID)
	Openfile, err := os.Open("fun.mp4")
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename=fun.mp4")
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}

func httpStart(){
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./dist"))
	r.Handle("/web/", http.StripPrefix("/web/", fs))
	r.HandleFunc("/user", authHandler(userHandler)).Methods("GET")
	r.HandleFunc("/video",  authHandler(videoHandler)).Methods("GET")
	r.HandleFunc("/videos",  generalHandler(userHandler))
	r.HandleFunc("/login",  generalHandler(userHandler))
	r.HandleFunc("/register",  generalHandler(userHandler))

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:" + os.Getenv("PORT"),
    }

    log.Fatal(srv.ListenAndServe())
}

func main() {
	httpStart()
}
