package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type dropZone struct {
	widget.BaseWidget
	background  *canvas.Rectangle
	label       *widget.Label
	inputEntry  *widget.Entry // Reference to input path entry
}


func newDropZone(entry *widget.Entry, w fyne.Widget) *dropZone {
	d := &dropZone{
		inputEntry: entry,
	}
	d.ExtendBaseWidget()
	d.background = canvas.NewRectangle(color.NRGBA{R: 229, G: 229, B: 229, A: 255})
	d.label = widget.NewLabel("Drop PDF file here")
	return d
}

type dropZoneRenderer struct {
    d *dropZone
}

func (d *dropZone) CreateRenderer() *dropZoneRenderer {
    return &dropZoneRenderer{
    }
}

func (r *dropZoneRenderer) Layout(size fyne.Size) {
    r.d.background.Resize(size)
    r.d.label.Resize(r.d.label.MinSize())
    r.d.label.Move(fyne.NewPos(
        (size.Width-r.d.label.MinSize().Width)/2,
        (size.Height-r.d.label.MinSize().Height)/2,
    ))
}

func (r *dropZoneRenderer) MinSize() fyne.Size {
    return fyne.NewSize(200.0, 100)
}
func (r *dropZoneRenderer) Destroy() {
    return 
} 

func (r *dropZoneRenderer) Refresh() {
    r.d.background.Refresh()
    r.d.label.Refresh()
}
func (d *dropZone) MinSize() fyne.Size{
    return fyne.NewSize(200, 100)
}




func (d *dropZone) DragEnter(event *fyne.DragEvent) {
    d.background.FillColor = color.NRGBA{R: 205, G: 205, B: 205, A: 255}
    d.background.Refresh()
}
func (d *dropZone) DragLeave() {
    d.background.FillColor = color.NRGBA{R: 229, G: 229, B: 229, A: 255}
    d.background.Refresh()
}

func (d *dropZone) DragMove(event *fyne.DragEvent) {}

func (d *dropZone) Drop(event *fyne.DragEvent) {
    d.background.FillColor = color.NRGBA{R: 229, G: 229, B: 229, A: 255}
    d.background.Refresh()
    

    // if uris, ok := event.Data.([]fyne.URI); ok && len(uris) > 0 {
    //     if ext := filepath.Ext(uris[0].Path()); ext == ".pdf" {
    //         inputPath.SetText(uris[0].Path())
    //     }
    // }
}