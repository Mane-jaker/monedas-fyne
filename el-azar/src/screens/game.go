package screens

import (
	"azar/src/models"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

func NewGame(window fyne.Window, app fyne.App) *Game {
	return &Game{window: window, app: app, container: container.NewWithoutLayout()}
}

func (g *Game) Play(bet string) {
	results := make(chan string)
	label := canvas.NewText("3", color.White)

	label.Resize(fyne.NewSize(100, 100))

	label.TextSize = 30

	label.Move(fyne.NewPos(735, 100))

	g.container.Add(
		label,
	)

	timer := time.NewTimer(time.Second)

	count := 3

	updateLabel := func() {
		count--

		label.Text = fmt.Sprint(count)
	}

	go func() {
		for range timer.C {
			updateLabel()
			timer.Reset(time.Second)
			if count == 0 {
				timer.Stop()
			}
		}
	}()

	for i := 0; i < 3; i++ {
		m := models.NewCoin()
		position := fyne.NewPos(float32(300+350*i), 300.0)
		m.Draw(g.container, position)
		g.window.SetContent(g.container)
		go m.Throw(results, time.Now().Second())
	}

	star, tree := 0, 0
	for i := 0; i < 3; i++ {
		result := <-results
		if result == "star" {
			star++
		} else {
			tree++
		}
	}

	winner := "star"
	if tree > star {
		winner = "tree"
	}

	message := "¡Has perdido!"
	if winner == bet {
		message = "¡Has ganado!"
	}

	text := canvas.NewText(message, color.White)

	text.TextSize = 30

	text.Move(fyne.NewPos(650, 600))

	menu := NewMainMenu(g.window, g.app)

	restart := widget.NewButton("Reintentar", menu.PrincipalMenu)

	restart.Resize(fyne.NewSize(200, 50))

	restart.Move(fyne.NewPos(645, 700))

	g.container.Add(
		text,
	)

	g.container.Add(restart)

	g.window.SetContent(g.container)
}
