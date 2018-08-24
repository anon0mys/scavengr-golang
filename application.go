package main

var api Api

func main() {
  api = Api{}
  api.Initialize()

  api.Run()
}
