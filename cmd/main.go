package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
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
var FileSelectionOutput *ui.Entry
var PlayerSelection *ui.Combobox

//Dropdown selections
var PitchTypes = []string{"Change Up", "Fastball", "Curveball", "Slider", " "}

var PitchLocations = []string{ // TODO fix this mess
	"Ball-Inside-High", "Ball-Inside-low", "Ball-Outside-High",
	"Ball-Outside-Low", "Ball-Middle Inside-High", "Ball-Middle Outside-High", "Ball-Inside-Low",
	"Ball-Outside-Low", "Strike-Inside-High", "Strike-Inside-Low",
	"Strike-Outside-Low", "Strike-Outside-High", "Strike-Center",
	"Strike-Center-High", "Strike-Center-Low", " "}

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

func refreshPlayerSelection() {
	PlayerSelection = ui.NewCombobox()
	filePlayerGrid.Append(PlayerSelection, 0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	PlayerNames = getPlayerNames(FileSelection)
	for _, value := range PlayerNames {
		PlayerSelection.Append(value)
	}
}

func setupUI() {
	mainWin = ui.NewWindow("Canavan Calculator", 370, 480, true)
	mainWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
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
		// verify valid input
		// send data to excel and save
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
}
