package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type App struct {
}

func main() {
	app := App{}

	log.Println("Starting broker service of port %s", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
