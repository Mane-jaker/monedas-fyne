package models

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Nivel struct {
	Elementos []*fyne.Container // Elementos del nivel, como obstáculos, enemigos, etc.
}

func NewNivel() *Nivel {
	nivel := &Nivel{
		Elementos: make([]*fyne.Container, 0),
	}

	// Agrega elementos al nivel, por ejemplo, bordes y un cuadrado negro
	nivel.agregarBordes()
	nivel.agregarCuadradoNegro()

	return nivel
}

func (n *Nivel) agregarBordes() {
	// Agregar bordes a los lados de la pantalla
	leftBorder := canvas.NewRectangle(color.Gray{0x99})
	leftBorder.Resize(fyne.Size{Width: 10, Height: 600}) // Utiliza fyne.Size en lugar de fyne.NewSize
	leftContainer := container.NewWithoutLayout(leftBorder)
	leftContainer.Move(fyne.NewPos(0, 0))
	n.Elementos = append(n.Elementos, leftContainer)

	rightBorder := canvas.NewRectangle(color.Gray{0x99})
	rightBorder.Resize(fyne.Size{Width: 10, Height: 600}) // Utiliza fyne.Size en lugar de fyne.NewSize
	rightContainer := container.NewWithoutLayout(rightBorder)
	rightContainer.Move(fyne.NewPos(790, 0))
	n.Elementos = append(n.Elementos, rightContainer)

	topBorder := canvas.NewRectangle(color.Gray{0x99})
	topBorder.Resize(fyne.Size{Width: 800, Height: 10}) // Utiliza fyne.Size en lugar de fyne.NewSize
	topContainer := container.NewWithoutLayout(topBorder)
	topContainer.Move(fyne.NewPos(0, 0))
	n.Elementos = append(n.Elementos, topContainer)

	bottomBorder := canvas.NewRectangle(color.Gray{0x99})
	bottomBorder.Resize(fyne.Size{Width: 800, Height: 10}) // Utiliza fyne.Size en lugar de fyne.NewSize
	bottomContainer := container.NewWithoutLayout(bottomBorder)
	bottomContainer.Move(fyne.NewPos(0, 590))
	n.Elementos = append(n.Elementos, bottomContainer)
}

func (n *Nivel) agregarCuadradoNegro() {
	// Agregar un cuadrado negro en el centro del nivel
	negro := canvas.NewRectangle(color.Black)
	negro.Resize(fyne.NewSize(100, 100))
	cuadradoNegro := container.NewWithoutLayout(negro)
	cuadradoNegro.Move(fyne.NewPos(350, 250))
	n.Elementos = append(n.Elementos, cuadradoNegro)
}
