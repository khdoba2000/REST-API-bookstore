package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khdoba2000/REST-API-bookstore/models"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := models.Articles{
		models.Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello World",
		},
	}

	fmt.Println("Endpoint hint: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func TestPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "testPost Article EndPoint Mit")

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage EndPoint Mit")
}
