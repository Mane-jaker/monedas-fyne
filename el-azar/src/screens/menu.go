package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createLabelContainers(labels []*widget.Label) []*fyne.Container {
	var labelContainers []*fyne.Container

	for _, label := range labels {
		// Crea un contenedor horizontal para cada etiqueta con espaciadores
		container := container.NewHBox(layout.NewSpacer(), label, layout.NewSpacer())
		labelContainers = append(labelContainers, container)
	}

	return labelContainers
}

type MainMenu struct {
	window fyne.Window
	app    fyne.App
}

func NewMainMenu(window fyne.Window, app fyne.App) *MainMenu {
	return &MainMenu{window: window, app: app}
}

func (m *MainMenu) PrincipalMenu() {

	titleLabel := widget.NewLabel("EL - AZAR")

	askGame := widget.NewLabel("¿A qué cara le apuestas?")

	labels := []*widget.Label{
		titleLabel,
		askGame,
	}

	labelContainers := createLabelContainers(labels)

	star := widget.NewButton("Estrella", func() { m.startGame("star") })

	tree := widget.NewButton("Trebol", func() { m.startGame("tree") })

	exit := widget.NewButton("Salir del juego", m.exitGame)

	contentContainer := container.NewVBox(labelContainers[0], labelContainers[1], star, tree, exit)

	centeredContent := container.NewCenter(contentContainer)

	m.window.SetContent(centeredContent)

}

func (m *MainMenu) startGame(apuesta string) {
	game := NewGame(m.window, m.app)
	game.Play(apuesta)
}

func (m *MainMenu) exitGame() {
	m.window.Close()
}
