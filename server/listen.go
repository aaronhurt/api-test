package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("-- New Request --\nRequest URL: %s\n", r.URL.String())
		for k, v := range r.Header {
			fmt.Printf("Request Header: %s => %s\n", k, v)
		}
		fmt.Printf("-- End-Of-Request --\n\n")
	})

	fmt.Printf("Listening on localhost:9900...\n")
	http.ListenAndServe("localhost:9900", nil)
}
