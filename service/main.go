package main

import (
	"fmt"
	"net/http"
)

/*
This is the service server.
The client will try to call this service and expects a response.
This service will expose a health check API for clients to check for its status
*/

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/operation", operation)

	fmt.Println("Service server started")
	http.ListenAndServe(":8000", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "online")
}

func operation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "operation success")
}
