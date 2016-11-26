package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/NikitaMasev/golang/track-server-consul/config"
	"github.com/NikitaMasev/golang/track-server-consul/handlersFuncs"
)

func main() {
	var cfgStruct = config.GetTrackServerConfig()
	cfgport := ":" + strconv.Itoa(cfgStruct.Port)
	http.HandleFunc("/", handlersFuncs.Handler)
	log.Fatal(http.ListenAndServe(cfgport, nil))
}
