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
    t.Errorf("Expected a message. Got: %s", point.Message)
  }

  if point.Clue != "test clue" {
    t.Errorf("Expected a clue. Got: %s", point.Clue)
  }

  if point.Address != "123 test address" {
    t.Errorf("Expected an address. Got: %s", point.Address)
  }

  if point.ScavengerHuntID != 2 {
    t.Errorf("Expected a scavenger hunt id. Got: %v", point.ScavengerHuntID)
  }

  if point.Location == nil {
    t.Errorf("Expected a location. Got: %s", point.Location)
  }
}

func TestPointReturnsAllPoints(t *testing.T) {
  
}
