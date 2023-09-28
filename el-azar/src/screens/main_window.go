package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainWindow struct {
	app    fyne.App
	window fyne.Window
	canvas fyne.Canvas
}

func NewMainWindow(title string, size fyne.Size) *MainWindow {
	mainApp := app.New()
	appWindow := mainApp.NewWindow(title)
	appWindow.Resize(size)
	return &MainWindow{
		app:    mainApp,
		window: appWindow,
		canvas: appWindow.Canvas(),
	}
}

func (m *MainWindow) Start() {
	mainMenu := NewMainMenu(m.window, m.app)
	mainMenu.PrincipalMenu()
	m.window.CenterOnScreen()
	m.window.SetFullScreen(true)
	m.window.ShowAndRun()
}
