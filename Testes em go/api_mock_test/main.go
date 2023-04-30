package main

import (
	"log"
	"net/http"

	"github.com/diegocassandri/livrogolang/api-test-mock/api"
)

func main() {
	service := api.Service{
		HTTPClient: http.DefaultClient,
	}

	http.HandleFunc("/", service.GetRandomUser)
	log.Println("Listening...")
	panic(http.ListenAndServe(":8000", nil))
}
