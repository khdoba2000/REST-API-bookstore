package urls

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khdoba2000/REST-API-bookstore/controllers"
)

func HandleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controllers.HomePage)
	myRouter.HandleFunc("/articles", controllers.AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", controllers.TestPostArticles).Methods("POST")
	log.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
