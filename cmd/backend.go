package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"github.com/skratchdot/open-golang/open"
	"time"
	"strconv"
)

// Constants for Excel
var headers = []string{"Date", "Plate Appearence", "Balls", "Strikes", "Total Pitches",
	"Pitch Type", "Pitch Location", "Outcome", "Hit Type", "Hit Direction", "Pitcher",
	"Pitcher Orientation", "Opponent"}

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

func addNewPlayerSheet(filepath string, sheetname string) error {
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

func showHitLocationGuideImage(filepath string) error {
	// Because the andlabs/ui doesn't have any image display function,
	// I have to do this shit
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return err
	}

	_, raw := f.GetPicture("Guide", "A2")
	ioutil.WriteFile("pic.jpg", raw, 0)
	open.Run("pic.jpg")
	return nil
}

func appendDataRow(filepath string, data DataInput) (err error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return
	}
	// get Plate Appearence index
	rows, err := f.Rows(data.sheetName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	plateAppearenceIndex := 0
	for rows.Next() {
		row := rows.Columns()
		if err != nil {
			fmt.Println(err.Error())
		}
		value, _ := strconv.ParseInt(row[1], 0, 64)
		if int(value) > plateAppearenceIndex {
			plateAppearenceIndex = int(value) + 1
			fmt.Println("PA", row[1], plateAppearenceIndex)
		}
	}

	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")
	balls := data.Balls
	strikes := data.Strikes

	dataToInsert := []int{
		plateAppearenceIndex,
		balls,
		strikes,
		balls + strikes }

	stringToInsert := []string{
		dataInputMappings[data.PitchType],
		data.PitchLocation,
		dataInputMappings[data.Outcome],
		dataInputMappings[data.HitType],
		dataInputMappings[data.HitLocations],
		data.PitcherName,
		dataInputMappings[data.PitcherHands],
		data.OpponentTeamName }

	f.InsertRow(data.sheetName, 1)
	f.SetCellValue(data.sheetName, "A2", currentTime)
	f.SetSheetRow(data.sheetName, "B2", &dataToInsert)
	f.SetSheetRow(data.sheetName, "F2", &stringToInsert)
	err = f.Save()
	if err != nil {
		return err
	}

	return nil
}
