// https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	staticFileDirectory := http.Dir("./assets/")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	fmt.Println("vim-go")
	r := newRouter()
	http.ListenAndServe(":8181", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
