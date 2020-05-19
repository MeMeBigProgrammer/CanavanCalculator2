package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// STRUCTURE
/*
main is the GUI with buttons and event handlers, and selections

backend.go takes the input, has input mappings to Excel inputs
interacts with Excel file

*/

var mainwin *ui.Window

//entries
var StrikesEntry *ui.Entry
var BallsEntry *ui.Entry
var PitcherNameEntry *ui.Entry
var OpponentTeamName *ui.Entry
var FilePathLabel *ui.Entry

var FileSelection string

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

//Dropdowns
var PitchTypeSelection *ui.Combobox
var PitchLocationSelection *ui.Combobox
var OutcomeSelection *ui.Combobox
var HitTypeSelection *ui.Combobox
var HitLocationsSelection *ui.Combobox
var PitcherHandSelection *ui.Combobox

func setupUI() {
	mainwin = ui.NewWindow("Canavan Calculator", 420, 440, true)
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

	// init Entries
	StrikesEntry := ui.NewEntry()
	BallsEntry := ui.NewEntry()
	PitcherNameEntry := ui.NewEntry()
	OpponentTeamName := ui.NewEntry()

	//init Dropdowns
	var PitchTypeSelection = ui.NewCombobox()
	for _, item := range PitchTypes {
		PitchTypeSelection.Append(item)
	}

	var PitchLocationSelection = ui.NewCombobox()
	for _, item := range PitchLocations {
		PitchLocationSelection.Append(item)
	}

	var OutcomeSelection = ui.NewCombobox()
	for _, item := range Outcomes {
		OutcomeSelection.Append(item)
	}

	var HitTypeSelection = ui.NewCombobox()
	for _, item := range HitTypes {
		HitTypeSelection.Append(item)
	}

	var HitLocationsSelection = ui.NewCombobox()
	for _, item := range HitLocations {
		HitLocationsSelection.Append(item)
	}

	var PitcherHandSelection = ui.NewCombobox()
	for _, item := range PitcherHands {
		PitcherHandSelection.Append(item)
	}

	mainForm.Append("Strikes", StrikesEntry, false)
	mainForm.Append("Balls", BallsEntry, false)
	mainForm.Append("Pitch Type", PitchTypeSelection, false)
	mainForm.Append("Pitch Location", PitchLocationSelection, false)
	mainForm.Append("Outcome", OutcomeSelection, false)
	mainForm.Append("Hit Type", HitTypeSelection, false)
	mainForm.Append("Hit Direction", HitLocationsSelection, false)
	mainForm.Append("Pitcher", PitcherNameEntry, false)
	mainForm.Append("Pitcher Handedness", PitcherHandSelection, false)
	mainForm.Append("Opponent Name", OpponentTeamName, false)

	inputGroup.SetChild(mainForm)

	//Excel,Player management
	excelGroup := ui.NewGroup("Destination Selection")
	filePlayerGrid := ui.NewGrid()
	excelGroup.SetChild(filePlayerGrid)

	FileSelection := ui.NewEntry()
	FileSelection.SetReadOnly(true)
	filePlayerGrid.Append(FileSelection,0, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignCenter)

	SelectFile := ui.NewButton("Open Excel File")
	filePlayerGrid.Append(SelectFile,1, 0, 1, 1,
		true, ui.AlignEnd, false, ui.AlignCenter)

	PlayerSelection := ui.NewCombobox()
	filePlayerGrid.Append(PlayerSelection,0, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignCenter)

	AddPlayer := ui.NewButton("New Player")
	filePlayerGrid.Append(AddPlayer,1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignCenter)

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
