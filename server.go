package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // string convert processes
	"strconv"

	"github.com/gorilla/mux" // For router
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

// create a post and append posts variable
func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost post
	// get request body and error from r.Body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If has a error show last user.
		fmt.Fprintf(w, "Please send post title and description!")
	}
	// json unmarshall parses the json-encoded data and stores result in value point to by &newPost
	json.Unmarshal(reqBody, &newPost)
	// append to posts
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}

// get all posts from posts variable
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	// json encode and send header application/json ?
	json.NewEncoder(w).Encode(posts)
}

// get single post
func getOnePost(w http.ResponseWriter, r *http.Request) {
	// get id
	postID := mux.Vars(r)["id"]
	// postID var is string but ID integer my post struct
	// convert str to int
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Fprintf(w, "String to int casting operation failed :(")
	}
	isFind := 0
	// iterate all posts
	for _, singlePost := range posts {
		if singlePost.ID == postIDInt {
			isFind = 1
			json.NewEncoder(w).Encode(singlePost)
		}
	}
	if isFind == 0 {
		fmt.Fprintf(w, "Post does not found :(")
	}
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
	// for get a single post call getOnePost func in get method
	router.HandleFunc("/posts/{id}", getOnePost).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

}
