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

func setupForm(form *ui.Form) {
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

	form.Append("Strikes", StrikesEntry, false)
	form.Append("Balls", BallsEntry, false)
	form.Append("Pitch Type", PitchTypeSelection, false)
	form.Append("Pitch Location", PitchLocationSelection, false)
	form.Append("Outcome", OutcomeSelection, false)
	form.Append("Hit Type", HitTypeSelection, false)
	form.Append("Hit Direction", HitLocationsSelection, false)
	form.Append("Pitcher", PitcherNameEntry, false)
	form.Append("Pitcher Handedness", PitcherHandSelection, false)
	form.Append("Opponent Name", OpponentTeamName, false)
}

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
	filePlayerGrid.SetPadded(true)
	excelGroup.SetChild(filePlayerGrid)

	AddPlayer = ui.NewButton("New Player")
	AddData = ui.NewButton("Append Data")
	SelectFile = ui.NewButton("Open Excel File")
	FileSelectionOutput = ui.NewEntry()
	PlayerSelection = ui.NewCombobox()


	FileSelectionOutput.SetReadOnly(true)
	filePlayerGrid.Append(FileSelectionOutput,0, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignCenter)

	filePlayerGrid.Append(SelectFile,1, 0, 1, 1,
		false, ui.AlignEnd, false, ui.AlignCenter)

	filePlayerGrid.Append(PlayerSelection,0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignCenter)

	filePlayerGrid.Append(AddPlayer,1, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignCenter)

	filePlayerGrid.Append(ui.NewLabel("Add Data"),0, 2, 1, 1,
		false, ui.AlignStart, false, ui.AlignCenter)

	filePlayerGrid.Append(AddData,1, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignCenter)


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
