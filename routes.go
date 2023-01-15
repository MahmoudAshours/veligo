package main

import "net/http"

func setupRoutesAndHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/:id", userHandler)
}
