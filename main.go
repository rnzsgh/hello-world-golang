package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Set up logging format to JSON and output to standard output
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Define the HTTP handler for the root ("/") URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Generate a Fibonacci function
		f := fib()

		// Write a "Hello World" message to the response
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World - new!\n")

		// Iterate through environment variables and write to the response
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			io.WriteString(w, pair[0]+"="+pair[1]+"\n")
		}

		// Generate and write the first 90 Fibonacci numbers to the response
		for i := 1; i <= 90; i++ {
			io.WriteString(w, strconv.Itoa(f())+"\n")
		}

		// Log an informational message
		log.Info("Hello world called - this is the log message")

	})

	// Start the HTTP server on port 80
	http.ListenAndServe(":80", nil)
}

// Function to generate Fibonacci numbers
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
