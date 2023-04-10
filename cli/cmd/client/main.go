package main

import (
	"log"

	"github.com/finallly/srp_protocol/cli/internal/srp_net"
	"github.com/finallly/srp_protocol/cli/pkg/parser"
)

const configName = "config"

func main() {
	if err := parser.ParseConfig(configName); err != nil {
		log.Fatalln(err.Error())
	}

	if err := srp_net.StartClientConnection(); err != nil {
		log.Fatalln(err.Error())
	}
}
