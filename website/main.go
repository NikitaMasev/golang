// website project main.go
package main

import (
	
	"github.com/NikitaMasev/golang/website/controllers"
	"net/http"
	"log"
)

func main() {
	controller:=controllers.MainController
	http.HandleFunc("/",controller.ControllerHTMLResponse)
	log.Fatal(http.ListenAndServe(":8085",nil))
}
