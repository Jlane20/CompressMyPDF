package main

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func compressIt(
	w fyne.Window, 
	statusLabel *widget.Label, 
	progress *widget.ProgressBar, 
	inPath string, 
	outPath string){
			// Update status to "Compressing..."
			fyne.CurrentApp().Driver().CreateWindow("Compressing")
			jobStarted(statusLabel, progress)
	
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
	
			// Update GUI with Done
			fyne.CurrentApp().Driver().CreateWindow("DONE")
			jobDone(progress, statusLabel, w, out, err)
	
}