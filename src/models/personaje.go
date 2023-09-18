package models

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Personaje struct {
	Circulo *canvas.Circle
	X       int // Posición X del personaje
	Y       int // Posición Y del personaje
}

func NewPersonaje(x, y int) *Personaje {
	circulo := canvas.NewCircle(color.White)
	circulo.StrokeColor = color.Gray{0x99}
	circulo.StrokeWidth = 10
	return &Personaje{
		Circulo: circulo,
		X:       x,
		Y:       y,
	}
}

func (p *Personaje) MoverIzquierda() {
	p.X -= 10 // Mueve el personaje 10 unidades a la izquierda
	p.Circulo.Move(fyne.NewPos(float32(p.X), float32(p.Y)))
}

func (p *Personaje) MoverDerecha() {
	p.X += 10 // Mueve el personaje 10 unidades a la derecha
	p.Circulo.Move(fyne.NewPos(float32(p.X), float32(p.Y)))
}

func (p *Personaje) Dibujar() *fyne.Container {
	x := float32(p.X)
	y := float32(p.Y)

	p.Circulo.Move(fyne.NewPos(0, 0))

	container := container.NewGridWrap(
		fyne.NewSize(50, 50),
		p.Circulo,
	)
	container.Move(fyne.NewPos(0, 0))

	go func() {
		time.Sleep(1 * time.Second)
		fyne.CurrentApp().Driver().CanvasForObject(p.Circulo).Refresh(p.Circulo)
		p.Circulo.Move(fyne.NewPos(x, y))
		container.Move(fyne.NewPos(x, y))
	}()

	return container

}

func (p *Personaje) MostrarPosicion() {
	fmt.Printf("Posición del personaje: X=%d, Y=%d\n", p.X, p.Y)
}
