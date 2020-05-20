package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// Maps for converting selections into Excel data

//take in path to file, check valid Excel and has Player sheets
func IsValidExcel(filepath string) (isValid bool, error error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
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

func GetPlayerNames(filepath string) (names []string) {
	f, _ := excelize.OpenFile(filepath)
	var playerList []string
	for _, name := range f.GetSheetMap() {
		if f.GetCellValue(name, "A1") != "NOT" {
			playerList = append(playerList, name)
		}
	}
	return playerList

}

func AddNewPlayerSheet(filepath string, sheetname string) {
	//f, _ := excelize.OpenFile(filepath)
}

func AppendDataRow(sheetname string) {

}
