package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {  
  if r.URL.Path != "/time" {
    http.Error(w, "Error 404 not found.", http.StatusNotFound)
    return
  }
  
  if r.Method != "GET" {
    http.Error(w, "This method is not supported.", http.StatusNotFound)
    return
  }

  fmt.Fprintf(w, "Your request: ")

  realtime := time.Now().Format(time.RFC3339)
  response := map[string]string{"time": realtime}

  w.Header().Set("Content-Type", "application/json")
  err := json.NewEncoder(w).Encode(response)
  if err != nil {
    log.Fatalf("Error while encoding: %s", err)
    return
  }
}