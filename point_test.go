package main_test

import (
  "testing"

  "github.com/anon0mys/scavengr-golang"
)

var point main.Point

func TestPointHasAttributes(t *testing.T) {

  point = main.Point{ Message: "Test adding point",
                      Clue: "test clue",
                      Address: "123 test address",
                    	ScavengerHuntID: 2,
                    	Location: []string{"40.12", "-70.34"}}

  if point.Message != "Test adding point" {
    t.Errorf("Expected a message. Got %s", point.Message)
  }

}
