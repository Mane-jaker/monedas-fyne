package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Coin struct {
	side    string
	sprites [9]*canvas.Image
}

func NewCoin() *Coin {
	return &Coin{
		sprites: [9]*canvas.Image{
			canvas.NewImageFromFile("assets/11.png"),
			canvas.NewImageFromFile("assets/12.png"),
			canvas.NewImageFromFile("assets/13.png"),
			canvas.NewImageFromFile("assets/14.png"),
			canvas.NewImageFromFile("assets/15.png"),
			canvas.NewImageFromFile("assets/16.png"),
			canvas.NewImageFromFile("assets/17.png"),
			canvas.NewImageFromFile("assets/18.png"),
			canvas.NewImageFromFile("assets/21.png"),
		},
	}
}

func (m *Coin) Throw(results chan<- string, rang int) {
	m.side = "" // The coin is spinning
	time.Sleep(2 * time.Second)
	random := (rand.Intn(rang)) % 2
	if (random) == 0 {
		m.side = "star"
	} else {
		m.side = "tree"
	}
	print(m.side)
	results <- m.side
}

func (m *Coin) Draw(container *fyne.Container, position fyne.Position) {
	spriteSize := fyne.NewSize(200, 200)
	go func() {
		spriteIndex := 0
		for {
			time.Sleep(100 * time.Millisecond)
			for i := 0; i < len(m.sprites); i++ {
				container.Remove(m.sprites[i])
			}
			if m.side == "" {
				m.sprites[spriteIndex].Resize(spriteSize)
				m.sprites[spriteIndex].Move(position)
				container.Add(m.sprites[spriteIndex])
				spriteIndex = (spriteIndex + 1) % 8
			} else if m.side == "tree" {
				m.sprites[0].Resize(spriteSize)
				m.sprites[0].Move(position)
				container.Add(m.sprites[0])
			} else {
				m.sprites[8].Resize(spriteSize)
				m.sprites[8].Move(position)
				container.Add(m.sprites[8])
			}
			container.Refresh()
		}
	}()
}
