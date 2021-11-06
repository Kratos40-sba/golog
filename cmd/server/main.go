package main

import (
	"github.com/Kratos40-sba/golog/internal/server"
	"log"
)

func main() {
	svc := server.NewHTTPServer(":8088")
	log.Fatal(svc.ListenAndServe())
}
