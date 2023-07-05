package main

import "fmt"

/*
	This is the client. The service that is being called by our users from the outside.
	This client will then implement the circuit breaker pattern and will try to call the service
*/

// Run go mod init <project name> first. This goes into a remote repository so don't forget to initialise a new repository

type CircuitBreaker interface{}

func main() {
	fmt.Println("Set up a basic server to implement the circuit breaker pattern")
}
