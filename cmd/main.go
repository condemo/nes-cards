package main

import (
	"flag"
	"fmt"

	"github.com/condemo/nes-cards/api"
)

func main() {
	addr := flag.String("p", ":4000", "addr")
	flag.Parse()

	apiServer := api.NewApiServer(*addr)
	fmt.Println("Server starting at port: ", *addr)
	apiServer.Run()
}
