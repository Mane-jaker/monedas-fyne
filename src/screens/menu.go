package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Mane-jaker/archero-2/src/screens"
)

func MainMenu() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("Menú Principal")
	myWindow.Resize(fyne.NewSize(450, 850))

	// Fondo azul oscuro
	bgColor := color.RGBA{R: 0, G: 0, B: 128, A: 255}
	background := canvas.NewRectangle(bgColor)
	background.Resize(fyne.NewSize(450, 850))
	bgContainer := container.NewWithoutLayout(background)

	// Botones
	boton1Jugador := widget.NewButton("1 Jugador", func() {
		gameWindow := screens.GameScreen()

		gameWindow.ShowAndRun

	})

	boton2Jugadores := widget.NewButton("2 Jugadores", func() {
		//boton2
	})

	botonSalir := widget.NewButton("Salir del Juego", func() {
		myApp.Quit()
	})

	// Estilo de botones
	boton1Jugador.Importance = widget.HighImportance
	boton2Jugadores.Importance = widget.HighImportance
	botonSalir.Importance = widget.HighImportance

	// Contenedor para los botones
	botones := container.NewVBox(
		boton1Jugador,
		boton2Jugadores,
		botonSalir,
	)
	botones.Layout = layout.NewVBoxLayout()

	contenido := container.NewVBox(
		canvas.NewRectangle(color.Transparent), // Agrega un espacio en blanco como margen arriba
		container.New(layout.NewCenterLayout(), botones),
	)

	// Agregar el fondo y el contenido a la ventana
	myWindow.SetContent(container.NewVBox(
		bgContainer,
		contenido,
	))

	return myWindow
}
