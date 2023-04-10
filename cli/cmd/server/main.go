package main

import (
	"log"
	"net"

	"github.com/finallly/srp_protocol/cli/internal/db"
	"github.com/finallly/srp_protocol/cli/internal/srp_net"
	"github.com/finallly/srp_protocol/cli/pkg/parser"
)

const configName = "config_server"

func main() {
	if err := db.PingDB(); err != nil {
		log.Fatalln(err.Error())
	}
	defer db.CloseDB()

	if err := parser.ParseConfig(configName); err != nil {
		log.Fatalln(err.Error())
	}

	listener, err := net.Listen("tcp", getAddress())
	if err != nil {
		log.Fatalln(err.Error())
	}
	srp_net.ListenerHandler(listener)
}

func getAddress() string {
	port := parser.GetDataFromConfig("port")
	return ":" + port
}
