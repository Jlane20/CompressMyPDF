package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildLayout(
	sliderLabel *widget.Label, 
	slider *widget.Slider, 
	outputPath *widget.Entry, 
	outputBtn *widget.Button, 
	compressBtn *widget.Button,
	progress *widget.ProgressBar,
	statusLabel *widget.Label) *fyne.Container{
		return container.NewVBox(
			container.NewVBox(
				widget.NewLabelWithStyle("PDF Compressor", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
				widget.NewSeparator(),
			),
			widget.NewLabel("(100 = High Compression, 300 = Low Compression)"),
			container.NewVBox(sliderLabel, slider),
			container.NewPadded(
				container.NewVBox(
					createFileSection("Input PDF 1: ", inputEntries[0], inputButtons[0]),
					createFileSection("Input PDF 2: ", inputEntries[1], inputButtons[1]),
					createFileSection("Input PDF 3: ", inputEntries[2], inputButtons[3]),
					createFileSection("Input PDF 4: ", inputEntries[3], inputButtons[4]),
					createFileSection("Input PDF 5: ", inputEntries[4], inputButtons[4]),
					createFileSection("Input PDF 6: ", inputEntries[5], inputButtons[5]),
					createFileSection("Output PDF:", outputPath, outputBtn),
					widget.NewSeparator(),
					compressBtn,
					progress,
					statusLabel,
				),
			),
		)
}