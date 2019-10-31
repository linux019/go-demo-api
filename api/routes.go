package api

import (
	"api-demo/api/controllers"
	log "api-demo/apilogger"
	c "api-demo/constants"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func AddRoutes(router *mux.Router) {
	if workDir, err := os.Getwd(); err == nil {
		webInterfaceHandler := http.StripPrefix(c.WebInterfaceURI, http.FileServer(http.Dir(workDir+c.WebInterfaceURI)))
		router.PathPrefix(c.WebInterfaceURI).Handler(webInterfaceHandler)
	} else {
		log.Logger.Fatal(err)
	}

	router.Use(loggingMiddleware, jsonMiddleware)
	router.HandleFunc("/create_account", controllers.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/authenticate", controllers.AuthenticateAccount).Methods(http.MethodPost)
	//router.HandleFunc("/demo", controllers.Demo).Methods(http.MethodPost)
}
