package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Addr: ":8089",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 1000 * time.Millisecond,
		ReadTimeout:  20 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	//logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
