package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var responseBody *string

func handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Unable to read body.")
		return
	}

	fmt.Printf("%s %s\n", r.Method, r.URL)
	for header, value := range r.Header {
		fmt.Printf("%s: %s\n", header, value[0])
	}
	fmt.Printf("\n%s\n\n\n", body)

	io.WriteString(w, *responseBody)
}

func main() {
	portNumber := flag.String("port", "80", "Port to listen on")
	responseBody = flag.String("response", "", "String response")
	flag.Parse()

	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port", *portNumber)
	if *responseBody != "" {
		fmt.Println("Response:", *responseBody)
	}

	http.ListenAndServe(":"+*portNumber, nil)
}
