package main

import (
	"api-demo/api"
	log "api-demo/apilogger"
	"api-demo/constants"
	"api-demo/db"
	"context"
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	if os.Getuid() == 0 {
		log.Logger.Fatal("Don't run me under root!")
	}

	port := flag.Int("p", constants.ServerPort, "Server port")
	dbUser := flag.String("user", constants.DbUser, "Database user")
	dbPassw := flag.String("password", constants.DbPassw, "Database password")
	flag.Parse()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, os.Kill)
	shutdownContext, cancel := context.WithTimeout(context.Background(), constants.GracefulTimeout)

	defer cancel()

	router := mux.NewRouter().StrictSlash(true)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(*port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Logger.Info("Staring HTTP server on port ", *port)
		log.Logger.Errorln(server.ListenAndServe())
		os.Exit(1)
	}()

	api.AddRoutes(router)
	db.ConnectDB(dbUser, dbPassw)
	defer db.DisconnectDB()

	<-signalChannel

	_ = server.Shutdown(shutdownContext)
	log.Logger.Info("shutting down")
	os.Exit(0)
}
