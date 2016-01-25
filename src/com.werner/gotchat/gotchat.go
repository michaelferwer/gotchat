package main

import (
	"net/http"
	"os"
	"github.com/apsdehal/go-logger"
	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Version 1.0.0\n"))
	return
}

func getLogger() *logger.Logger {
	log, err := logger.New("LOG", 1, os.Stdout)
	if err != nil {
		panic(err) // TODO Check for error
	}
	return log
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ping).Methods("GET")
	http.ListenAndServe(":8000", router)
}
