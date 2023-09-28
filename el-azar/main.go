package main

import (
	"azar/src/screens"

	"fyne.io/fyne/v2"
)

func main() {
	window := screens.NewMainWindow("Juego Go", fyne.Size{Width: 450, Height: 650})
	window.Start()
}
