package main

import (
	"errors"
)

var dataInputMappings = map[string]string{
	" ":                  "",
	"Change Up":          "CH",
	"Fastball":           "FB",
	"Curveball":          "CB",
	"Slider":             "SL",
	"Hard Hit":           "H",
	"Routine Hit":        "R",
	"Weak Hit":           "W",
	"Standing Strikeout": "uK",
	"Strikeout":          "K",
	"Walk":               "BB",
	"Hit By Pitch":       "HBP",
	"Grounder":           "G",
	"Line-Drive":         "L",
	"Fly":                "F",
	"Pitcher":            "1",
	"Catcher":            "2",
	"First Base":         "3",
	"Second Base":        "4",
	"Shortstop":          "6",
	"Third Base":         "5",
	"Right Field":        "9",
	"Center Field":       "8",
	"Left Field":         "7",
	"Left Handed":        "LHP",
	"Right Handed":       "RHP"}

type DataInput struct {
	Strikes          int
	Balls            int
	PitchType        string
	PitchLocation    string
	Outcome          string
	HitType          string
	HitLocations     string
	PitcherName      string
	PitcherHands     string
	OpponentTeamName string
	sheetName        string
}

func (data DataInput) IsValidInput() (isvalid bool, error error) {
	/* RULES:
	1.) Strikes must be 0-2
	2.) Balls must be 0-3
	3.) Hit Direction and type must be blank if there is a HBP, BB, or Strikeout
	4.) MUST HAVE A PLAYER INPUT
	*/
	error = nil
	isvalid = true

	if data.sheetName == " " {
		isvalid = false
		error = errors.New("You must select a player!")
	}

	if data.Strikes < 0 || data.Strikes > 2 {
		isvalid = false
		error = errors.New("Strikes must be between 0 and 2!")
	} else if data.Balls < 0 || data.Balls > 3 {
		isvalid = false
		error = errors.New("Balls must be between 0 and 3!")
	}

	shouldBeBlank := false
	if data.Outcome == "Standing Strikeout" ||
		data.Outcome == "Strikeout" ||
		data.Outcome == "Walk" ||
		data.Outcome == "Hit By Pitch" {
		shouldBeBlank = true
	}

	if shouldBeBlank && data.HitLocations != " " {
		isvalid = false
		error = errors.New("Impossible Input!")
	} else if shouldBeBlank && data.HitType != " " {
		isvalid = false
		error = errors.New("Impossible Input!")
	}

	return isvalid, error
}
