package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Mane-jaker/archero-2/src/models"
)

func GameScreen() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mi Juego")
	myWindow.Resize(fyne.Size{Width: 450, Height: 850})

	// Crear una instancia de Personaje
	personaje := models.NewPersonaje(0, 50)

	myWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyLeft {
			personaje.MoverIzquierda()
		} else if event.Name == fyne.KeyRight {
			personaje.MoverDerecha()
		}
	})

	// Agregar el contenido del personaje a la ventana
	myWindow.SetContent(personaje.Dibujar())
	myWindow.ShowAndRun()

	return myWindow
}
