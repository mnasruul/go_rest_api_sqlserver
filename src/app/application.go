package app

import (
	"github.com/gorilla/mux"
	"github.com/mnasruul/go_rest_api_sqlserver/src/datasources/sqlserver/db"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	sqlserver.Init()

	mapUrls()

	srv := &http.Server{
		Addr: ":8089",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	//logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
