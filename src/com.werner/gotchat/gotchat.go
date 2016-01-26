package main

import (
	"net/http"
	"os"
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
)

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "version": "1.0.0" })
}

func getLogger() *logger.Logger {
	log, err := logger.New("LOG", 1, os.Stdout)
	if err != nil {
		panic(err) // TODO Check for error
	}
	return log
}

func main() {
	router := gin.Default()
	router.GET("/version", version)
	router.Run(":8080")
}
