package config

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
)

//TrackServerConfig is a config data...
type TrackServerConfig struct {
	Port int
}

//GetTrackServerConfig make getting config param...
func GetTrackServerConfig() TrackServerConfig {
	var tsc TrackServerConfig

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
	kv := client.KV()
	port, _, err := kv.Get("TrackServerConfig/port", nil)
	if err != nil {
		fmt.Println(err)
	}
	portTmp, err := strconv.Atoi(string(port.Value))
	if err != nil {
		fmt.Println(err)
	}
	tsc.Port = portTmp
	return tsc
}
