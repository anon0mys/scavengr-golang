package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Api struct {
  Router *mux.Router
  Server *http.Server
}

func (app *Api) Initialize() {
  app.Router = mux.NewRouter()
  app.InitializeRoutes(app)
}

func (app *Api) Run() {
  fmt.Printf("Server starting, listening on port: %v\n", app.Server.Addr)
  log.Fatal(app.Server.ListenAndServe())
}
