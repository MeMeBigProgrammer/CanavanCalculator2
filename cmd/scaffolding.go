package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func setupForm(form *ui.Form) {
	form.SetPadded(true)

	// init Entries
	StrikesEntry = ui.NewEntry()
	BallsEntry = ui.NewEntry()
	PitcherNameEntry = ui.NewEntry()
	OpponentTeamName = ui.NewEntry()

	//init Dropdowns
	PitchTypeSelection = ui.NewCombobox()
	for _, item := range PitchTypes {
		PitchTypeSelection.Append(item)
	}

	PitchLocationSelection = ui.NewCombobox()
	for _, item := range PitchLocations {
		PitchLocationSelection.Append(item)
	}

	OutcomeSelection = ui.NewCombobox()
	for _, item := range Outcomes {
		OutcomeSelection.Append(item)
	}

	HitTypeSelection = ui.NewCombobox()
	for _, item := range HitTypes {
		HitTypeSelection.Append(item)
	}

	HitLocationsSelection = ui.NewCombobox()
	for _, item := range HitLocations {
		HitLocationsSelection.Append(item)
	}

	PitcherHandSelection = ui.NewCombobox()
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
	grid.SetPadded(true)
	SelectFile = ui.NewButton("Open Excel File")
	FileSelectionOutput = ui.NewEntry()

	AddPlayer = ui.NewButton("New Player")
	PlayerSelection = ui.NewCombobox()

	AddData = ui.NewButton("Append Data")

	ShowHitLocations = ui.NewButton("Show Guide")

	AddPlayer.Disable()
	AddData.Disable()
	PlayerSelection.Disable()
	ShowHitLocations.Disable()

	grid.Append(ShowHitLocations, 1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("Toggle"), 0, 0, 1, 1,
		false, ui.AlignStart, false, ui.AlignFill)

	FileSelectionOutput.SetReadOnly(true)
	grid.Append(FileSelectionOutput, 0, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	grid.Append(SelectFile, 1, 1, 1, 1,
		false, ui.AlignEnd, false, ui.AlignFill)

	grid.Append(PlayerSelection, 0, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(AddPlayer, 1, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("Add Data"), 0, 3, 1, 1,
		false, ui.AlignStart, false, ui.AlignFill)

	grid.Append(AddData, 1, 3, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	ShowHitLocations.OnClicked(func(*ui.Button) {
		// Get image from excel, if there is One
		// save copy of jpg next to .exe
		// open using run
		// delete on close
		showHitLocationGuideImage(FileSelection)
	})
}

func setupDialogBox(box *ui.Window, entry *ui.Entry, button *ui.Button) {
	hBox := ui.NewHorizontalBox()
	group := ui.NewGroup("")

	group.SetChild(hBox)
	box.SetChild(group)
	hBox.SetPadded(true)

	hBox.Append(entry, true)
	hBox.Append(button, false)

	button.OnClicked(func(*ui.Button) {
		err := addNewPlayerSheet(FileSelection, entry.Text())
		if err == nil {
			ui.MsgBox(box, "Sucess!", "A a new player has been added.")
			box.Destroy()
			refreshPlayerSelection()
		} else {
			ui.MsgBoxError(box, "Error!", err.Error())
		}
	})

}
