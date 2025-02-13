package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const numInputs = 6
var inputEntries = make([]*widget.Entry, numInputs)
var inputButtons = make([]*widget.Button, numInputs)


func main() {
	a := app.NewWithID("234409dijosidf")
	w := a.NewWindow("PDF Compressor")
	w.SetContent(mainUI(w))
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}

func mainUI(w fyne.Window) fyne.CanvasObject {
	
	// Output entries
	outputPath := widget.NewEntry()
	outputBtn := createOutputButton(w, outputPath)
	
	// Create buttons and entries for input
	for i := 0; i < numInputs; i++ {
		inputEntries[i] = widget.NewEntry()
		inputButtons[i] = createInputButton(w, i)
	}

	// Compression Setting Slider
	sliderLabel := widget.NewLabel("Compression Value: 0.0")
	slider := createSlider(sliderLabel)

	// Progress and status
	progress := widget.NewProgressBar()
	statusLabel := widget.NewLabel("Ready")

	// Compress button
	compressBtn := createCompressButton(outputPath, w, statusLabel, progress)

	// Layout
	layout := buildLayout(sliderLabel, &slider, outputPath, outputBtn, compressBtn, progress, statusLabel) 
	return layout
}

//-----------------------------------Helper Functions------------------------------//
func createFileSection(label string, entry *widget.Entry, btn *widget.Button) fyne.CanvasObject {
	return container.NewBorder(nil, nil, widget.NewLabel(label), btn, entry)
}
func jobStarted(statusLabel *widget.Label, progress *widget.ProgressBar) {
	statusLabel.SetText("Compressing...")
	progress.SetValue(25.0)
}
func jobDone(progress *widget.ProgressBar, statusLabel *widget.Label, w fyne.Window, out[]byte, err error) {
	if err != nil {
		statusLabel.SetText("Error")
		dialog.ShowError(
			fmt.Errorf("compression failed: %v\n%s", err, string(out)),
			w)
		return
	}
	progress.SetValue(1.0) // "100%"
	statusLabel.SetText("Done!")
}