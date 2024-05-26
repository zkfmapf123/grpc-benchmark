package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Println("Connect To Sub Server : ", port)

	// http
	r := mux.NewRouter()
	r.HandleFunc("/", health).Methods("GET")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/sub", sub).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success")
	w.Write([]byte("success"))
}

func sub(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("HTTP >> ", formattedTime)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
