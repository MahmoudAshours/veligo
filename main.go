package main

import (
	"net/http"
)

func main() {
	setupRoutesAndHandlers()
	http.ListenAndServe(":3000", nil)
}
