package main

import (
	"fmt"
	"net/http"
	"kpi/lab1/handlers"
	"log"
)
  
  
func main(){
	http.HandleFunc("/time", handlers.TimeHandler)
	fmt.Print("Starting server at port 8795")

	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}