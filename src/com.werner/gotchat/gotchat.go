package main

import (
	"net/http"
	"os"
	"github.com/apsdehal/go-logger"
)

func handleConnection (w http.ResponseWriter, r *http.Request) {
	log := getLogger()
	log.Debug("Connection")
	w.Header().Set("Content-Type", "plain/text; charset=utf-8")
	w.Write([]byte("Hello World"))
	return
}

func getLogger () *logger.Logger{
	log, err := logger.New("LOG", 1, os.Stdout)
	if err != nil {
		panic(err) // TODO Check for error
	}
	return log
}

func main () {
	http.HandleFunc("/", handleConnection)
  http.ListenAndServe(":8080", nil)
}
