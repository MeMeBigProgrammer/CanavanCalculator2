package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

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

func setupSettingsGrid(grid *ui.Grid) {
	SelectFile = ui.NewButton("Open Excel File")
	FileSelectionOutput = ui.NewEntry()

	AddPlayer = ui.NewButton("New Player")
	PlayerSelection = ui.NewCombobox()

	AddData = ui.NewButton("Append Data")

	AddPlayer.Disable()
	AddData.Disable()
	PlayerSelection.Disable()

	FileSelectionOutput.SetReadOnly(true)
	grid.Append(FileSelectionOutput, 0, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	grid.Append(SelectFile, 1, 0, 1, 1,
		false, ui.AlignEnd, false, ui.AlignFill)

	grid.Append(PlayerSelection, 0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(AddPlayer, 1, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("Add Data"), 0, 2, 1, 1,
		false, ui.AlignStart, false, ui.AlignFill)

	grid.Append(AddData, 1, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
}

func setupDialogBox(box *ui.Window, entry *ui.Entry, button *ui.Button) {
	hbox := ui.NewHorizontalBox()
	group := ui.NewGroup("")

	group.SetChild(hbox)
	box.SetChild(group)
	hbox.SetPadded(true)

	hbox.Append(entry, true)
	hbox.Append(button, false)

}
