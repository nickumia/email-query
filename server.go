//go:build OMIT
// +build OMIT

// The server program issues Google search requests and demonstrates the use of
// the go.net Context API. It serves on port 8080.
//
// The /search endpoint accepts these query params:
//
//	q=the Google search query
//	timeout=a timeout for the request, in time.Duration format
//
// For example, http://localhost:8080/search?q=golang&timeout=1s serves the
// first few Google search results for "golang" or a "deadline exceeded" error
// if the timeout expires.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Email struct {
	Email string
}

func main() {
	// http.HandleFunc("/", handleEmail)
	// err := http.ListenAndServe(":8080", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/email", handleEmail)
	err := http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleEmail(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Email
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(t)
	fmt.Fprintf(w, "Email: %+v", t.Email)
}
