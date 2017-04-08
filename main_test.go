package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestArticlesHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	//Make the handler function satisfy http.Handler
	//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
	http.HandlerFunc(articlesHandler).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	//Confirm the returned json is what we expected
	//Manually build up the expected json string
	expected := string(`[{"id":1,"title":"New blog resolution","content":"I have decided to give my blog a new life and would hence forth try to write as often"},{"id":2,"title":"Go is cool","content":"Yeah i have been told that multiple times"},{"id":3,"title":"Interminttent fasting","content":"You should try this out, it helps clear the brain and tons of health benefits"},{"id":4,"title":"Yet another blog post","content":"I made a resolution earlier to keep on writing. Here is an affirmation of that"},{"id":5,"title":"Backpacking","content":"Yup, i did just that"}]`)

	//The assert package checks if both JSON string are equal and for a plus, it actually confirms if our manually built JSON string is valid
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}

func TestArticleHandlerWithValidPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/2", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/posts/{id:[0-9]+}", articleHandler).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

	expected := string(`{"id":2,"title":"Go is cool","content":"Yeah i have been told that multiple times"}`)

	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}

func TestArticleHandlerWithAnInvalidPost(t *testing.T) {

	req, err := http.NewRequest("POST", "/posts/42", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/posts/{id:[0-9]+}", articleHandler).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusNotFound, status)
	}

	expected := "404 page not found\n"

	assert.Equal(t, expected, rr.Body.String(), "Response body differs")
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}
