package main

import (
	"bewell-backend-challenge/internal/adapter/router"
	"bewell-backend-challenge/internal/config"
	"bewell-backend-challenge/util/helpers/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/graceful"
)

func main() {
	appconfig := config.Read()
	ginDefault := gin.Default()
	
	router := router.NewRouter(ginDefault, appconfig)

	serTimeout := common.Timeout() * 2 // nolint:gomnd,mnd
	gracefulServer := graceful.WithDefaults(&http.Server{
		Addr:         ":" + appconfig.AppPort,
		Handler:      router,
		ReadTimeout:  serTimeout,
		WriteTimeout: serTimeout,
	})

	log.Println("Starting the server on port...", appconfig.AppPort)
	if err := graceful.Graceful(gracefulServer.ListenAndServe, gracefulServer.Shutdown); err != nil {
		log.Println("Failed to gracefully shutdown")
	}
}