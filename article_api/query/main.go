package main

import (
	"article_api/query/config/pgadapter"
	"article_api/query/models"
	"article_api/query/pipeline"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var Adapter = pgadapter.Adapter{}.New()
	Adapter.Table.AutoMigrate(&models.Article{})
	Adapter.Connection.Close()
	r := mux.NewRouter()
	r.HandleFunc("/", pipeline.GetArticle).Methods("GET")
	r.HandleFunc("/create", pipeline.CreateArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
