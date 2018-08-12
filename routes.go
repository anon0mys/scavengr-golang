package main

func (app *Api) InitializeRoutes(*Api) {
  app.Router.HandleFunc("/points/", app.AllPoints).Methods("GET")
  app.Router.HandleFunc("/points", app.AllPoints).Methods("GET")
}
