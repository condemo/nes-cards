package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/condemo/nes-cards/api"
	"github.com/condemo/nes-cards/store"
)

func main() {
	addr := flag.String("p", ":4000", "addr")
	flag.Parse()

	sqliteStorage := store.NewSqliteStore()
	db, err := sqliteStorage.Init()
	if err != nil {
		log.Fatal("database error: ", err)
	}

	store := store.NewStorage(db)

	apiServer := api.NewApiServer(*addr, store)
	fmt.Println("Server starting at port", *addr)
	apiServer.Run()
}
