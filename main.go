package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":7070", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, my name is JSONpea"))
}
