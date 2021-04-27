package controllers

import (
	"encoding/json"
	"github.com/khdoba2000/REST-API-bookstore/models"
	"net/http"
	"strconv"


)


func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello World",
		},
	}

	fmt.Println("Endpoint hint: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "testPost Article EndPoint Mit")

}



func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage EndPoint Mit")
}