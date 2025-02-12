package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var inputPath = widget.NewEntry()


func main() {
	a := app.NewWithID("234409dijosidf")
	w := a.NewWindow("PDF Compressor")
	w.SetContent(mainUI(w))
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}

func mainUI(w fyne.Window) fyne.CanvasObject {
	// Input/Output entries
	outputPath := widget.NewEntry()

	// Button to browse for input PDF (File Open)
	inputBtn := widget.NewButton("Browse...", func() {
		fileOpen := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, err error) {
				if err != nil || r == nil {
					return
				}
				inputPath.SetText(r.URI().Path())
			}, w)

		// Limit to only .pdf files (optional)
		fileOpen.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fileOpen.Show()
	})
	valueLabel := widget.NewLabel("Compression Value: 0.0")
	slider := widget.Slider{
		Value: 100.0,
		Min: 100,
		Max: 300,
		Step: 20,

		OnChanged: func(val float64) {
			valueLabel.SetText(fmt.Sprintf("Compression Value: %.1f", val))
		},
		OnChangeEnded: func(val float64) {
			fmt.Println("Final Value Selected:", val)
		},
	}
	dropZone := newDropZone(inputPath)	// Button to browse for output PDF (File Save)

	outputBtn := widget.NewButton("Browse...", func() {
		if inputPath.Text == "" {
			dialog.ShowInformation("No Input Selected",
				"Please choose an input file first.", w)
			return
		}

		// Suggest an output filename based on the input
		inputExt := filepath.Ext(inputPath.Text)
		baseName := strings.TrimSuffix(filepath.Base(inputPath.Text), inputExt)
		suggestedName := baseName + "_compressed" + inputExt

		// Use the same directory as input for convenience
		suggestedDir := filepath.Dir(inputPath.Text)
		dirURI := storage.NewFileURI(suggestedDir)
		listableDir, err := storage.ListerForURI(dirURI)
		if err != nil {
			dialog.ShowError(fmt.Errorf("invalid directory: %v", err), w)
			return
		}



		fileSave := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, err error) {
				if err != nil || uc == nil {
					return
				}
				outputPath.SetText(uc.URI().Path())
			}, w)

		// Pre-fill suggestions
		fileSave.SetFileName(suggestedName)
		fileSave.SetLocation(listableDir)
		fileSave.Show()
	})

	// Compression Settings
	qualityPreset := widget.NewRadioGroup([]string{
		"Screen (Smallest)",
		// "Ebook (Recommended)",
		// "Printer",
		// "Prepress (Best Quality)",
	}, func(s string) {
		// handle changes if needed
	})
	qualityPreset.SetSelected("Screen (Smallest)")

	// Progress and status
	progress := widget.NewProgressBar()
	statusLabel := widget.NewLabel("Ready")

	// Map from our radio labels to Ghostscript settings
	// gsPresets := map[string]string{
	// 	"Screen (Smallest)":       "/screen",
	// 	// "Ebook (Recommended)":     "/ebook",
	// 	// "Printer":                 "/printer",
	// 	// "Prepress (Best Quality)": "/prepress",
	// }

	// Compress action
	compressBtn := widget.NewButton("Compress PDF", func() {
		inPath := inputPath.Text
		outPath := outputPath.Text
		if inPath == "" || outPath == "" {
			dialog.ShowInformation("Missing Paths",
				"Please specify both input and output files before compressing.",
				w)
			return
		}

		// Map the radio selection to a Ghostscript preset
		// selectedPreset := gsPresets[qualityPreset.Selected]
		// if selectedPreset == "" {
		// 	// fallback in case user hasn't selected anything
		// 	selectedPreset = "/ebook"
		// }

		// Let’s do some basic validation for PDF extension
		if strings.ToLower(filepath.Ext(inPath)) != ".pdf" {
			dialog.ShowError(
				fmt.Errorf("input file must be a PDF: %s", inPath), w)
			return
		}

		// Start a goroutine so the UI won't freeze
		go func() {
			// // Update status to "Compressing..."
			// fyne.CurrentApp().Driver().CreateWindow()
			// (func() {
			// 	statusLabel.SetText("Compressing...")
			// 	progress.SetValue(0.0) // if you want to show "0%" at start
			// })

			// Construct the Ghostscript command
			cmd := exec.Command("gs",
				"-sDEVICE=pdfwrite",
			 	"-dCompatibilityLevel=1.7",
			  	// "-dPDFSETTINGS=/screen", // This seems pointless with all the other params to compresss
   				"-dDownsampleColorImages=true",
  				"-dColorImageResolution=200",
   				"-dDownsampleGrayImages=true", 
				"-dGrayImageResolution=200",
				"-dDownsampleMonoImages=true", 
				"-dMonoImageResolution=200",
				// "-dCompressFonts=true", 
				"-dDetectDuplicateImages=true",
				"-dNOPAUSE",
				"-dBATCH",
				fmt.Sprintf("-sOutputFile=%s", outPath),
				inPath,
			)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(out))
			// Use CallOnMainThread again to update UI after command completes
			// fyne.CurrentApp().Driver().CallOnMainThread(func() {
			// 	if err != nil {
			// 		statusLabel.SetText("Error")
			// 		dialog.ShowError(
			// 			fmt.Errorf("compression failed: %v\n%s", err, string(out)),
			// 			w)
			// 		return
			// 	}
			// 	progress.SetValue(1.0) // "100%"
			// 	statusLabel.SetText("Done!")
			// })
		}()
	})

	// Layout
	return container.NewVBox(
		container.NewCenter(dropZone),
		container.NewVBox(
			widget.NewLabelWithStyle("PDF Compressor", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewSeparator(),
		),
		widget.NewLabel("(100 = High Compression, 300 = Low Compression)"),
		container.NewVBox(valueLabel, &slider),
		container.NewPadded(
			container.NewVBox(
				createFileSection("Input PDF:", inputPath, inputBtn),
				createFileSection("Output PDF:", outputPath, outputBtn),
				widget.NewSeparator(),
				widget.NewLabel("Compression Settings:"),
				qualityPreset,
				widget.NewSeparator(),
				compressBtn,
				progress,
				statusLabel,
			),
		),
	)
}

func createFileSection(label string, entry *widget.Entry, btn *widget.Button) fyne.CanvasObject {
	return container.NewBorder(nil, nil, widget.NewLabel(label), btn, entry)
}
