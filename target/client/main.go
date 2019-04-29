package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

type config struct {
	Address string `env:"ADDRESS" envDefault:":1111"`
}

var CFG config

func main() {
	err := env.Parse(&CFG)
	if err != nil {
		log.Fatal(err)
	}
	rtr := router.New()
	root := rtr.Root()
	root.GET(HandleHome, "renders dashboard")
	err = http.ListenAndServe(CFG.Address, rtr)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request, p *params.Params) {

}
