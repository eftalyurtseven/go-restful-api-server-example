package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux" // For router?
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type post struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allPosts []post

var posts = allPosts{
	{
		ID:          1,
		Title:       "Test",
		Description: "Deneme bir açıklama",
	},
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost post
	reqBody, err := ioutil.ReadAll(r.Body) // get request body and error from r.Body
	if err != nil {
		fmt.Fprintf(w, "Please send post title and description!") // If has a error show last user.
	}
	json.Unmarshal(reqBody, &newPost) // json unmarshall parses the json-encoded data and stores result in value point to by &newPost
	posts = append(posts, newPost)    // append to posts
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	// json encode and send header application/json ?
	json.NewEncoder(w).Encode(posts)
}

func main() {

	// create a router
	router := mux.NewRouter().StrictSlash(true)
	// for / path call homeRoute func
	router.HandleFunc("/", homeRoute)
	// for posts/ path call create func in post method
	router.HandleFunc("/posts", createPost).Methods("POST")
	// for posts/ path call getAllPosts func in get method
	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))

}
