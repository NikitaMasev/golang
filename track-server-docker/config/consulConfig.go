package config

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/hashicorp/consul/api"
)

//TrackServerConfig is a config data...
type TrackServerConfig struct {
	Port int
}

//Fill filling struct data...
func Fill(cfg interface{}) error {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return errors.New("Can't connect to consul")
	}
	kv := client.KV()
	// Folder for searching inner keys
	folder := reflect.TypeOf(cfg).Elem().Name()

	elem := reflect.ValueOf(cfg).Elem()

	for i := 0; i < elem.NumField(); i++ {
		// Use field name as a key
		key := elem.Type().Field(i).Name
		val, _, err := kv.Get(folder+"/"+key, nil)
		if err != nil {
			return errors.New("Error of getting value from consul")
		}
		switch elem.Field(i).Interface().(type) {
		case int, int32, int64:
			v, err := strconv.Atoi(string(val.Value))
			if err != nil {
				return errors.New("Error of transformation from string to int")
			}
			elem.Field(i).SetInt(int64(v))

		case string:
			v := string(val.Value)
			elem.Field(i).SetString(v)
		}
	}
	return nil
}
