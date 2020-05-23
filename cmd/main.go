package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
	"os"
)

var mainWin *ui.Window
var filePlayerGrid *ui.Grid

// Entries
var StrikesEntry *ui.Entry
var BallsEntry *ui.Entry
var PitcherNameEntry *ui.Entry
var OpponentTeamName *ui.Entry
var FilePathLabel *ui.Entry

// Dropdowns
var PitchTypeSelection *ui.Combobox
var PitchLocationSelection *ui.Combobox
var OutcomeSelection *ui.Combobox
var HitTypeSelection *ui.Combobox
var HitLocationsSelection *ui.Combobox
var PitcherHandSelection *ui.Combobox

//Settings
var AddPlayer *ui.Button
var SelectFile *ui.Button
var AddData *ui.Button
var ShowHitLocations *ui.Button
var FileSelectionOutput *ui.Entry
var PlayerSelection *ui.Combobox

//Dropdown selections
var PitchTypes = []string{"Change Up", "Fastball", "Curveball", "Slider", " "}

var PitchLocations = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"11", "12", " "}

var Outcomes = []string{
	"Hard Hit", "Routine Hit", "Weak Hit", "Standing Strikeout", "Strikeout", "Walk",
	"Hit By Pitch", " "}

var HitTypes = []string{
	"Grounder", "Line-Drive", "Fly", " "}

var HitLocations = []string{
	"Pitcher", "Catcher", "First Base", "Second Base", "Shortstop", "Third Base",
	"Right Field", "Center Field", "Left Field", " "}

var PitcherHands = []string{
	"Left Handed", "Right Handed", " "}

// Global Veriables
var FileSelection string
var PlayerNames = []string{}

func getComboBoxValue(box *ui.Combobox, options []string) string {
	inputIndex := strconv.Itoa(box.Selected())
	if inputIndex == "-1" {
		return " "
	}
	index, _ := strconv.ParseInt(inputIndex, 0, 64)
	return options[index]
}

func refreshPlayerSelection() {
	PlayerSelection = ui.NewCombobox()
	filePlayerGrid.Append(PlayerSelection, 0, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	names, err := getPlayerNames(FileSelection)
	if err != nil {
		ui.MsgBoxError(mainWin, "Error!", err.Error())
		return
	} else {
		PlayerNames = names
	}

	for _, value := range PlayerNames {
		PlayerSelection.Append(value)
	}
}

func setupUI() {
	mainWin = ui.NewWindow("Canavan Calculator", 370, 480, true)
	mainWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		os.Remove("pic.jpg")
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
		os.Remove("pic.jpg")
		return true
	})

	// init Form Layout
	inputGroup := ui.NewGroup("Inputs")
	mainForm := ui.NewForm()
	setupForm(mainForm)
	inputGroup.SetChild(mainForm)

	//Excel,Player management
	excelGroup := ui.NewGroup("Settings")
	filePlayerGrid = ui.NewGrid()
	setupSettingsGrid(filePlayerGrid)
	excelGroup.SetChild(filePlayerGrid)

	// button actions
	SelectFile.OnClicked(func(*ui.Button) {
		FileSelection = ui.OpenFile(mainWin)
		FileSelectionOutput.SetText(FileSelection)
		isValid, _ := isValidExcel(FileSelection)
		if isValid {
			refreshPlayerSelection()
			AddPlayer.Enable()
			AddData.Enable()
			PlayerSelection.Enable()
		} else {
			ui.MsgBoxError(mainWin, "Error!", "Not a valid Excel file!")
			AddPlayer.Disable()
			AddData.Disable()
			PlayerSelection.Disable()
		}
	})

	AddPlayer.OnClicked(func(*ui.Button) {
		dialogBox := ui.NewWindow("New Player", 150, 75, false)
		dialogBox.OnClosing(func(*ui.Window) bool {
			return true
		})
		nameEntry := ui.NewEntry()
		submitButton := ui.NewButton("Add")
		setupDialogBox(dialogBox, nameEntry, submitButton)
		dialogBox.Show()
	})

	AddData.OnClicked(func(*ui.Button) {
		// convert entries to ints
		strikesNumber, StrikeErr := strconv.ParseInt(StrikesEntry.Text(), 0, 64)
		ballsNumber, BallErr := strconv.ParseInt(BallsEntry.Text(), 0, 64)

		if StrikeErr != nil || BallErr != nil {
			ui.MsgBoxError(mainWin, "Error!", "Balls and Strikes must be a number!")
			return
		}

		newData := DataInput{
			Strikes:          int(strikesNumber),
			Balls:            int(ballsNumber),
			PitchType:        getComboBoxValue(PitchTypeSelection, PitchTypes),
			PitchLocation:    getComboBoxValue(PitchLocationSelection, PitchLocations),
			Outcome:          getComboBoxValue(OutcomeSelection, Outcomes),
			HitType:          getComboBoxValue(HitTypeSelection, HitTypes),
			HitLocations:     getComboBoxValue(HitLocationsSelection, HitLocations),
			PitcherName:      PitcherNameEntry.Text(),
			PitcherHands:     getComboBoxValue(PitcherHandSelection, PitcherHands),
			OpponentTeamName: OpponentTeamName.Text(),
			sheetName:        getComboBoxValue(PlayerSelection, PlayerNames)}

		isValid, err := newData.IsValidInput()

		if !isValid {
			ui.MsgBoxError(mainWin, "Error!", err.Error())
		} else {
			appendDataRow(FileSelection, newData)
			// send data to excel and save
		}
	})

	// append all groups
	mainVBox := ui.NewVerticalBox()
	mainVBox.SetPadded(true)
	mainVBox.Append(inputGroup, true)
	mainVBox.Append(excelGroup, false)

	// main window
	mainWin.SetChild(mainVBox)
	mainWin.SetMargined(true)

	mainWin.Show()
}

func main() {
	ui.Main(setupUI)
	/* TODO:
	1. Add Data Append and verify
	2. Add image display
	*/
}
