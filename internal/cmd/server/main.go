package main

import (
	"log"

	"github.com/foreverLoveWisdom/commit-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
