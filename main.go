package main

import (
	"net/http"
)

func main() {
	setupRoutesAndHandlers()
	http.ListenAndServe(":8080", nil)
}
