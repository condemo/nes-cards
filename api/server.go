package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condemo/nes-cards/api/handlers"
	"github.com/condemo/nes-cards/api/middlewares"
	"github.com/condemo/nes-cards/service"
	"github.com/condemo/nes-cards/store"
)

type ApiServer struct {
	addr  string
	store store.Store
}

func NewApiServer(addr string, store store.Store) *ApiServer {
	return &ApiServer{
		addr:  addr,
		store: store,
	}
}

func (s *ApiServer) Run() {
	router := http.NewServeMux()
	views := http.NewServeMux()
	game := http.NewServeMux()
	player := http.NewServeMux()
	fs := http.FileServer(http.Dir("public/static"))

	// Middlewares
	basicMDstack := middlewares.MiddlewareStack(
		middlewares.Recover,
		middlewares.Logger,
	)

	router.Handle("/", basicMDstack(views))
	router.Handle("/game/", http.StripPrefix("/game", basicMDstack(game)))
	router.Handle("/player/", http.StripPrefix("/player", basicMDstack(player)))
	router.Handle("/static/", http.StripPrefix("/static", fs))

	server := http.Server{
		Addr:         s.addr,
		Handler:      router,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 5,
	}

	// GameService
	gc := service.NewGameController()

	viewsHandler := handlers.NewViewsHandler(s.store, gc)
	viewsHandler.RegisterRoutes(views)
	gameHandler := handlers.NewGameHandler(s.store, gc)
	gameHandler.RegisterRoutes(game)
	playerHandler := handlers.NewPlayerHandler(s.store)
	playerHandler.RegisterRoutes(player)

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigC

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// server.Shutdown ends the execution of the program
	// after waiting for all active connections to finish or 30 seconds to pass
	server.Shutdown(ctx)
}
