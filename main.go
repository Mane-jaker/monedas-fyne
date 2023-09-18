package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Mane-jaker/archero-2/src/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mi Juego")

	// Establecer el tamaño de la ventana en 300px x 300px
	myWindow.Resize(fyne.Size{Width: 400, Height: 300})

	// Crear una instancia de Personaje
	personaje := models.NewPersonaje(0, 50)

	myWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyLeft {
			// Manejar la tecla izquierda
			personaje.MoverIzquierda()
		} else if event.Name == fyne.KeyRight {
			// Manejar la tecla derecha
			personaje.MoverDerecha()
		}
	})

	// Agregar el contenido del personaje a la ventana
	myWindow.SetContent(personaje.Dibujar())
	myWindow.ShowAndRun()
}
