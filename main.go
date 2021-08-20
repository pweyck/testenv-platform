package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shopwareLabs/testenv-platform/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	go handler.PullImageUpdatesTask()

	router := httprouter.New()

	router.GET("/", handler.Info)

	// New Routes
	router.GET("/environments", handler.ListContainer)
	router.POST("/environments", handler.CreateEnvironment)
	router.DELETE("/environments", handler.DeleteContainer)

	log.Println("Go!")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
