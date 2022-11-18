package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi joe\n")
	})
	fmt.Println("Listening and serving...")
	http.ListenAndServe(":8090", nil)
}
