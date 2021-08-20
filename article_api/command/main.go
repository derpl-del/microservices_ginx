package main

import (
	"article_api/command/config/pgadapter"
	"article_api/command/models"
	"article_api/command/pipeline"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var Adapter = pgadapter.Adapter{}.New()
	Adapter.Table.AutoMigrate(&models.Article{})
	Adapter.Connection.Close()
	r := mux.NewRouter()
	r.HandleFunc("/", pipeline.CreateArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
