package controllers

import (
	"net/http"
	"github.com/NikitaMasev/golang/website/models"
	"github.com/NikitaMasev/golang/website/views"
)
//Controller struct...
type Controller struct{}

//MainController is a global variable of controller...
var MainController Controller

//ControllerHTMLResponse func of controller for rendering html view...
func (c Controller) ControllerHTMLResponse(w http.ResponseWriter, req *http.Request)  {
    data:=models.ModelData
    views.CreateHTMLResponseView(w,data)
}