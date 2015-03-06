package main

import (
	"fmt"
	"net/http"
)

const listeningAddress = ":7070"

func main() {
	fmt.Printf("Listening on %s ...\n", listeningAddress)
	http.HandleFunc("/", hello)
	http.ListenAndServe(listeningAddress, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, my name is JSONpea"))
}
