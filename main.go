package main

import (
	"Best_trash_API/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"
	manyBooksResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env")
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server ON:", port)
	router := mux.NewRouter()
	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManybooksResourcePrefix(router, manyBooksResourcePrefix)
	log.Println("Router initalizing succesfully!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
