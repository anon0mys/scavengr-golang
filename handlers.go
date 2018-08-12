package main

import (
  "net/http"
  "encoding/json"
  "fmt"
)

func (app *Api) AllPoints(writer http.ResponseWriter, req *http.Request) {
  var point Point

  points := point.AllPoints()
  respondWithJSON(writer, http.StatusOK, points)
}

func respondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  fmt.Printf("%v", response)

  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(code)
  writer.Write(response)
}
