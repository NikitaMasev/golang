package handlersFuncs

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

const urlHead = "http://"

//Handler is a func for http.HandleFunc...
func Handler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	url, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, req, urlHead+string(url), http.StatusFound)
}