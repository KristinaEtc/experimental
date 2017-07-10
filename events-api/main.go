package main

import _ "github.com/KristinaEtc/slflog"

import (
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/ventu-io/slf"
)

var log = slf.WithContext("main.go")

// MemcacheConf is a struct with memcache configs.
type MemcacheConf struct {
	Host     string
	Enabled  bool
	Instance string
}

// ConfFile is a file with all program options
type ConfFile struct {
	Name           string
	Host           string
	MemcacheConfig MemcacheConf
}

var globalOpt = ConfFile{
	Name: "config",
	Host: "localhost:7778",
	MemcacheConfig: MemcacheConf{
		Host:     "localhost:1111",
		Enabled:  true,
		Instance: "test-instance",
	},
}

func runProcess(conn *ws.Conn) {
	//select
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go runProcess(conn)
}

func main() {
	log.Info("Starting work")
	http.HandleFunc("/", wsHandler)
	panic(http.ListenAndServe(":8080", nil))
}
