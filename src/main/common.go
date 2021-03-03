package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, status int, payload interface{}) {
  response, err := json.Marshal(payload)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    // log error
    return
  }

  log.Print(response, "test")

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write([]byte(response))
}
