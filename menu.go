package main

import (
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	selected    = 0
	lastClickAt time.Time
)

type menu struct {
	changeScreen func(s screen)
}

func (m *menu) Update() error {
	if time.Since(lastClickAt) < debouncer {
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		lastClickAt = time.Now()
		selected--
		selected %= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		lastClickAt = time.Now()
		selected++
		selected %= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if selected == 0 {
			initGame()
			g := &Game{state: playerTurn}
			g.Restart()
			m.changeScreen(g)
		}

	}
	return nil
}

func (m *menu) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0), float64(160))
	screen.DrawImage(createMenuImage(), op)
}

func createMenuImage() *ebiten.Image {
	dc := gg.NewContext(screenWidth, screenHeight)
	dc.SetRGB(112, 98, 200)
	dc.DrawRectangle(100, 0, screenWidth-200, 50)
	dc.DrawRectangle(100, 50+30, screenWidth-200, 50)
	dc.DrawRectangle(100, 50*2+30*2, screenWidth-200, 50)
	dc.Fill()
	dc.SetRGB(229, 227, 0)
	dc.DrawRectangle(120, float64(50*selected)+float64(30*selected)+5, screenWidth-240, 40)
	dc.Fill()
	dc.SetRGB(112, 223, 200)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	dc.DrawStringAnchored("Start", screenWidth/2, 20, 0.5, 0.5)
	dc.DrawStringAnchored("Challenges", screenWidth/2, 20+50+30, 0.5, 0.5)
	dc.DrawStringAnchored("Exit", screenWidth/2, 20+50*2+30*2, 0.5, 0.5)
	return ebiten.NewImageFromImage(dc.Image())
}
