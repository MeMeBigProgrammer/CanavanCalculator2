package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

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

var PitchLocations = []string{
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

func setupUI() {
	mainwin = ui.NewWindow("Canavan Calculator", 400, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	// init Form Layout
	inputGroup := ui.NewGroup("Inputs")
	mainForm := ui.NewForm()
	mainForm.SetPadded(true)
	setupForm(mainForm)
	inputGroup.SetChild(mainForm)

	//Excel,Player management
	excelGroup := ui.NewGroup("Settings")
	filePlayerGrid := ui.NewGrid()
	setupSettingsGrid(filePlayerGrid)
	filePlayerGrid.SetPadded(true)
	excelGroup.SetChild(filePlayerGrid)

	SelectFile.OnClicked(func(*ui.Button) {
		FileSelection = ui.OpenFile(mainwin)
		FileSelectionOutput.SetText(FileSelection)
		isValid, _ := IsValidExcel(FileSelection)
		if isValid {
			// I have to replace the combobox, buh moment
			PlayerSelection = ui.NewCombobox()
			filePlayerGrid.Append(PlayerSelection, 0, 1, 1, 1,
				false, ui.AlignFill, false, ui.AlignFill)

			PlayerNames = GetPlayerNames(FileSelection)
			for _, value := range PlayerNames {
				PlayerSelection.Append(value)
			}
			AddPlayer.Enable()
			AddData.Enable()
			PlayerSelection.Enable()

		} else {
			AddPlayer.Disable()
			AddData.Disable()
			PlayerSelection.Disable()
		}
	})

	AddPlayer.OnClicked(func(*ui.Button) {
		dialogBox := ui.NewWindow("New Player", 240, 75, false)
		dialogBox.OnClosing(func(*ui.Window) bool {
			return true
		})
		nameEntry := ui.NewEntry()
		submitButton := ui.NewButton("Add")
		setupDialogBox(dialogBox, nameEntry, submitButton)
		dialogBox.Show()
	})

	// append all groups
	mainVBox := ui.NewVerticalBox()
	mainVBox.SetPadded(true)
	mainVBox.Append(inputGroup, true)
	mainVBox.Append(excelGroup, false)

	// main window
	mainwin.SetChild(mainVBox)
	mainwin.SetMargined(true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
