package main

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func createSlider(sliderLabel *widget.Label) widget.Slider {
	return widget.Slider{
		Value: 100.0,
		Min: 100,
		Max: 300,
		Step: 20,

		OnChanged: func(val float64) {
			sliderLabel.SetText(fmt.Sprintf("Compression Value: %.1f", val))
		},
		OnChangeEnded: func(val float64) {
			fmt.Println("Final Value Selected:", val)
		},
	}
}