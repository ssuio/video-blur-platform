package main

import (
	"fmt"
	// "time"
	// "log"
	"net/http"
	// "github.com/gorilla/sessions"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("hello"))
}

func userHandler(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	fmt.Fprintf(w, string("user"))
	w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Something bad happened!"))
}

func videoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("video"))
}

func httpStart(){
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":9000", nil)
}

func main() {
	httpStart()
}
