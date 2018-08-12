package main_test

import (
  "testing"
  "os"
  "net/http"
  "net/http/httptest"

  "github.com/anon0mys/scavengr-golang"
)

var app main.Api

func TestMain(m *testing.M) {
  app = main.Api{}
  os.Setenv("ELASTIC_SEARCH_URL", "localhost:9200")
  os.Setenv("ENV", "test")
  app.Initialize()

  code := m.Run()

  resetIndicies()

  os.Exit(code)
}

func TestEmptyIndex(t *testing.T) {
  resetIndicies()

  req, _ := http.NewRequest("GET", "/points", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  app.Router.ServeHTTP(rr, req)

  return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
}

func resetIndicies() {
  // Migrations for Elasticsearch go here!!
}
