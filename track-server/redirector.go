// track-server project main.go
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

const urlHead = "http://"

func Handler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	url, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, req, urlHead+string(url), http.StatusFound)
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
