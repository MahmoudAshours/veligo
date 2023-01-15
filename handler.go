package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, users!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, user!")
}
