package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// Constants for Excel
var headers = []string{"Date", "Plate Appearence", "Balls", "Strikes", "Total Pitches",
	"Pitch Type", "Pitch Location", "Outcome", "Hit Type", "Hit Direction", "Pitcher",
	"Pitcher Orientation", "Opponent", "Pitcher"}

// Q: Why open file in every function?
// A: Reduces chance of IO conflict

//take in path to file, check valid Excel and if it has Player sheets
func isValidExcel(filepath string) (isValid bool, error error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return false, err
	}
	for _, name := range f.GetSheetMap() {
		if f.GetCellValue(name, "A1") == "NOT" {
			fmt.Println("Valid!")
			return true, nil
		}
	}
	return false, nil
}

func getPlayerNames(filepath string) (names []string, err error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}
	var playerList []string
	for _, name := range f.GetSheetMap() {
		if f.GetCellValue(name, "A1") != "NOT" {
			playerList = append(playerList, name)
		}
	}
	return playerList, nil

}

func addNewPlayerSheet(filepath string, sheetname string) (error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return err
	}

	f.NewSheet(sheetname)
	f.SetSheetRow(sheetname, "A1", &headers)
	err = f.Save()
	if err != nil {
		return err
	}

	return nil

}

func appendDataRow(sheetname string, data DataInput) (err error) {

	return nil

}
