package main

import (
	"fmt"
	"kpi/lab1/handlers"
	"log"
	"net/http"
)

const port = ":8795"

func main() {
	http.HandleFunc("/time", handlers.TimeHandler)
	fmt.Print("Starting server at port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
