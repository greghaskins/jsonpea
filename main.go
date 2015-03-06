package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const listeningAddress = ":7070"

func main() {
	fmt.Printf("Listening on %s ...\n", listeningAddress)
	http.HandleFunc("/get", handleGet)
	http.ListenAndServe(listeningAddress, nil)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	encodedUrl := r.URL.Query().Get("url")

	remoteUrl, err := url.QueryUnescape(encodedUrl)
	if err != nil {
		log.Printf("Error parsing URL: %s\n", remoteUrl)
		return
	}

	data, err := fetchBody(remoteUrl)
	if err != nil {
		log.Printf("Error fetching <%s>: %s\n", remoteUrl, err)
		return
	}

	callback := r.URL.Query().Get("callback")
	if len(callback) > 0 {
		w.Write([]byte(callback))
		w.Write([]byte("("))
		w.Write(data)
		w.Write([]byte(");"))
	} else {
		w.Write(data)
	}
}

func fetchBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
