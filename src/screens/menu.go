package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
		// Aquí puedes agregar la lógica para el modo de 1 jugador
	})

	boton2Jugadores := widget.NewButton("2 Jugadores", func() {
		// Aquí puedes agregar la lógica para el modo de 2 jugadores
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
