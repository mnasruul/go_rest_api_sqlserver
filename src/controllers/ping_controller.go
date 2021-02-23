package controllers

import (
	"log"
	"net/http"
)


const (
	pong = "pong2"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	log.Println(pong)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
