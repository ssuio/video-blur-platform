package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("hello"))
}

func userHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("user"))
}

func videoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("video"))
}

func httpStart(){
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

func main() {
	httpStart()
}
