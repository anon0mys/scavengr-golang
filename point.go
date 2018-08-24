package main

import (
)


type Point struct {
  ID int              `json:"id"`
  PointID int         `json:"point_id"`
  UserId int          `json:"user_id"`
  ScavengerHuntID int `json:"scavenger_hunt_id"`
  Message string      `json:"message"`
  Location []string   `json:"location"`
  Clue string         `json:"clue"`
  Address string      `json:"address"`
  Found bool          `json:"found"`
}

func (point *Point) AllPoints() []Point {
  var points []Point

  return points
}
