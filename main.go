package main

import (
	"fmt"
	"github.com/DMLayMan/PeterBot/db"
	"log"
	"net/http"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	initRouters()

	log.Fatal(http.ListenAndServe(":80", nil))
}
