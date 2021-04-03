package main

import (
	"log"
	"net/http"

	"github.com/IamNator/thealth/internal/router"
)

func main() {
	r, _ := router.Router()
	log.Println("server running on localhost:9000")
	http.ListenAndServe(":9000", r)
}
