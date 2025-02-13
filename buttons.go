package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func createInputButton(w fyne.Window, i int) *widget.Button{
	button:= widget.NewButton("Browse...", func() {
		fileOpen := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, err error) {
				if err != nil || r == nil {
					return
				}
				inputEntries[i].SetText(r.URI().Path())
			}, w,)

		// Limit to only .pdf files
		fileOpen.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fileOpen.Resize(fyne.NewSize(1000, 800))
		fileOpen.Show()
	})
	return button
}
func createOutputButton(w fyne.Window, outputPath *widget.Entry) *widget.Button{
	return widget.NewButton("Browse...", func() {
		
		// Validate an input has been selected
		if inputEntries[0].Text == "" {
			dialog.ShowInformation("No Input Selected","Please choose an input file first.", w)
			return
		}

        folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
            if err != nil || uri == nil {
                return
            }
            // Store the selected folder path
            outputPath.SetText(uri.Path())
        }, w)
        folderDialog.Show()
    })
}
func createCompressButton(
	outputPath *widget.Entry, 
	w fyne.Window, 
	statusLabel *widget.Label, 
	progress *widget.ProgressBar)*widget.Button{
	return widget.NewButton("Compress PDF", func() {
		
		outPath := outputPath.Text
		allEmpty := true

		// If no output has been set throw up the dialog
		if outPath == "" {dialog.ShowInformation("Missing Path", "Please specify an output path.",w)}
		
		// Check for empty
		for _, button := range inputButtons {
			if button.Text != "" {
				allEmpty = false
				break
			}
		}

		// Return a helpful message if they forgot to add an input to compress.
		if allEmpty {dialog.ShowInformation("Missing Paths", "Please specify both input and output files before compressing.",w)}
		
		// Iterate over each file and call compressIt
		for _, entry := range inputEntries {
			inPath := entry.Text
			// They may not go in order so skip if empty.	
			if inPath == ""{ 
				continue 
			}

			// Basic validation for PDF extension
			if strings.ToLower(filepath.Ext(inPath)) != ".pdf" {
				dialog.ShowError(fmt.Errorf("input file must be a PDF: %s", inPath), w)
				return
			}		

			inputExt := filepath.Ext(entry.Text)
			baseName := strings.TrimSuffix(filepath.Base(entry.Text), inputExt)
			suggestedName := baseName + "_compressed" + inputExt
			finalPath := fmt.Sprintf("%s%s", outPath, suggestedName)

			// Start a goroutine so the UI won't freeze
			go func() {compressIt(w, statusLabel, progress, inPath, finalPath)}() 
		}
	})
}
