package main

import (
	"github.com/DMLayMan/PeterBot/service"
	"net/http"
)

func initRouters() {
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/gpt", service.GPTHandler)
}
