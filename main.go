package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func handle(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, "Unable to read body.")
		return
	}

	fmt.Printf("%s %s\n", request.Method, request.URL)
	for header, value := range request.Header {
		fmt.Printf("%s: %s\n", header, value[0])
	}
	fmt.Printf("\n%s\n\n\n", body)

	response.WriteHeader(200)
}

func main() {
	portNumber := os.Getenv("LISTEN_PORT")
	if portNumber == "" {
		portNumber = "80"
	}
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
