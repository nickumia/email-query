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
	"net/http"
	"os"
	"strings"
	"time"
)

type Email struct {
	Email string
}

var now time.Time
var replacer = strings.NewReplacer(" ", "_")

func main() {
	// http.HandleFunc("/", handleEmail)
	// err := http.ListenAndServe(":8080", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/email", handleEmail)
	err := http.ListenAndServe(":8080", mux)

	// Timestamp for file
	now = time.Now()

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
	// TODO: Save to file
	// saveFile, err := os.OpenFile(replacer.Replace(now.Format("Mon Jan 2 15:04:05 MST 2006")), os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer saveFile.Close()
	// f := bufio.NewWriter(saveFile)
	// _, err = f.Write([]byte(t.Email + "\n"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// f.Flush()
	fmt.Println(t.Email)
	fmt.Fprintf(w, "Email: %+v", t.Email)
}
