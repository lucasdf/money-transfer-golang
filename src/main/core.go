package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/google/uuid"
  "github.com/shopspring/decimal"
  "github.com/gorilla/mux"
)

type UserBalance struct {
  Balance decimal.Decimal `json:"balance"`
  Id uuid.UUID `json:"id"`
}

func balance(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  log.Print(vars)
  balance := UserBalance{Balance: decimal.NewFromFloat(10.0), Id: uuid.New()}
  response(w, http.StatusOK, balance)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/balance/{id}", balance).Methods("GET")
  http.Handle("/", r)
  fmt.Println("Listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
