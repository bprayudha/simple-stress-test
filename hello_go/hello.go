package main

import (
  "log"
  "net/http"
  "github.com/bitly/go-simplejson"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", index).Methods("GET")

  http.Handle("/", router)

  log.Println("Listening on :3000")
  log.Fatal(http.ListenAndServe(":3000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
  json := simplejson.New()
  json.Set("hello", "world")

  payload, err := json.MarshalJSON()
  if err != nil {
    log.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(payload)
}
