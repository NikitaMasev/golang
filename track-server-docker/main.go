package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/NikitaMasev/golang/track-server-docker/config"
	"github.com/NikitaMasev/golang/track-server-docker/handlersFuncs"
)

func main() {
	var tsc = &config.TrackServerConfig{}
	err := config.Fill(tsc)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handlersFuncs.Handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(tsc.Port), nil))
}

/* Instruction for docker

1. Run Consul : $ ./consul agent -server -bind 127.0.0.1 -ui-dir ~/consul/ui/ -data-dir ~/consul/data/
2. Build Docker container: $ sudo docker build -t app .
3. Run Docker container: $ sudo docker run -e CONSUL_HTTP_ADDR=8500 -p 8085:8085 app

*/
