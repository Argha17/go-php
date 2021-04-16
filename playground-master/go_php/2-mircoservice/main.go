package main

import (
	"github.com/Argha17/playground/go_php/2-microservice/http"
	"github.com/Argha17/playground/go_php/2-microservice/handler"
)

func main() {
	handler.DbInit()
	handler.DbStoreData()

	r := httphandler.New()
	r.GET("/ping", handler.Ping)
	r.GET("/search", handler.SearchMoviews)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
