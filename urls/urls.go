package urls

import (
	"log"
	"net/http"
	"github.com/khdoba2000/REST-API-bookstore/controllers"
	"github.com/gorilla/mux"
)

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controllers.homePage)
	myRouter.HandleFunc("/articles", controllers.allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", controllers.testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
