package handlers

import (
	"Best_trash_API/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBookId(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parametr ID"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindBookId(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book...")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookId int = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parametr ID"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	oldBook, ok := models.FindBookId(id)
	if !ok {
		log.Println("book not found in database:", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {

	}
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
}
