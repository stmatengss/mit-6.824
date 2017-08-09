package main

import (
		"fmt"
		"net/http"
)

func main() {

		hello := func(w http.ResponseWriter, r *http.Request) {
				res := r.URL.Path[len("/hello/"):]
				fmt.Fprintf(w, "Hello %s!\n", res)
		}

		http.HandleFunc("/hello/", hello)
		http.ListenAndServe("localhost:8888", nil)
}
