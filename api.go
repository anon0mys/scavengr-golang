package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

type Api struct {
  Router *mux.Router
  Server *http.Server
}

func (app *Api) Initialize() {
  app.Router = mux.NewRouter()
  app.InitializeRoutes(app)

  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

  app.Server = &http.Server{Addr: ":" + os.Getenv("PORT"), Handler: handlers.CORS(headersOk, originsOk, methodsOk)(app.Router)}
}

func (app *Api) Run() {
  fmt.Printf("Server starting, listening on port: %v\n", app.Server.Addr)
  log.Fatal(app.Server.ListenAndServe())
}
