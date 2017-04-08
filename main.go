package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var posts []*post

type post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {

	posts = []*post{
		{1, "New blog resolution", "I have decided to give my blog a new life and would hence forth try to write as often"},
		{2, "Go is cool", "Yeah i have been told that multiple times"},
		{3, "Interminttent fasting", "You should try this out, it helps clear the brain and tons of health benefits"},
		{4, "Yet another blog post", "I made a resolution earlier to keep on writing. Here is an affirmation of that"},
		{5, "Backpacking", "Yup, i did just that"},
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world..\n Visit the '/posts' (GET) route to get all posts. " +
			"\n '/posts/id' (GET) to get a specific post." +
			"\n '/posts/id' (POST) to create a new post . \n " +
			"'/posts/delete' (PUT) to delete a post  "))
	}).Methods("GET")

	r.HandleFunc("/posts", articlesHandler).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", articleHandler).Methods("GET", "POST")
	r.HandleFunc("/posts/delete", deleteArticleHandler).Methods("PUT")

	log.Println("Starting server at port 4000")

	http.ListenAndServe(":4000", r)
}

//Fetches all posts
func articlesHandler(w http.ResponseWriter, r *http.Request) {

	users, _ := json.Marshal(posts) //Handle errors in real life

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(users))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	//Gorilla mux stores url mappings as strings,
	//We would have to convert them to an int in other to use it as an index for fetchong the specified post
	postId, _ := strconv.Atoi(id)

	var postFound bool
	var p *post

	for _, v := range posts {
		if v.ID == postId {
			postFound = true
			p = v
			break
		}
	}

	if postFound {
		w.WriteHeader(http.StatusOK)
		requestedPost, _ := json.Marshal(p)
		fmt.Fprintf(w, string(requestedPost))
		return
	}

	//Throw a 404
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {

}