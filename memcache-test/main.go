package main

import (
	"fmt"

	"github.com/KristinaEtc/config"
	_ "github.com/KristinaEtc/slflog"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/ventu-io/slf"
)

var log = slf.WithContext("memcache.go")

// ConfFile is a file with all program options
type ConfFile struct {
	Name     string
	Host     string
	Enabled  bool
	Instance string
}

var globalOpt = ConfFile{
	Name:     "config",
	Host:     "localhost:1111",
	Enabled:  true,
	Instance: "test-instance",
}

func main() {

	config.ReadGlobalConfig(&globalOpt, "template options")

	log.Infof("%s", globalOpt.Name)

	log.Error("----------------------------------------------")
	log.Info("Starting working...")
	mc := memcache.New(globalOpt.Host)
	err := mc.Set(&memcache.Item{Key: "key_one", Value: []byte("kovalevskayakv")})
	if err != nil {
		log.Errorf("setting a value: %s", err.Error())
	}

	// Get a single value
	val, err := mc.Get("key_one")
	if err != nil {
		log.Errorf("getting a value: %s", err.Error())
		return
	}

	fmt.Printf("-- %s\n", val.Value)

	// Get multiple values
	it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {
		log.Errorf("getMulti: %s", err)
		return
	}

	// It's important to note here that `range` iterates in a *random* order
	for k, v := range it {
		log.Infof("## %s => %s\n", k, v.Value)
	}

}
