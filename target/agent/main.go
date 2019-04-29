package main

import (
	"log"
	"net"

	"github.com/caarlos0/env"
)

type config struct {
	Title   string `env:"TITLE"   envDefault:"CI test"`
	Address string `env:"ADDRESS" envDefault:":1000"`
	Network string `env:"NETWORK" envDefault:"tcp"`
}

var CFG config

func main() {
	err := env.Parse(&CFG)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen(CFG.Network, CFG.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(connection net.Conn) {
			_, err := connection.Write([]byte{})
			if err != nil {
				log.Fatal(err)
			}
			connection.Close()
		}(connection)
	}
}
