package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

const (
	title   = "3 cards"
	command = `Press ENTER to start the game`
)

var (
	mplusNormalFont font.Face
	mplusTitleFont  font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusTitleFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type menu struct {
	changeScreen func(s screen)
}

func (m *menu) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		initGame()
		g := &Game{turn: playerTurn}
		g.Restart()
		m.changeScreen(g)
	}
	return nil
}

func (m *menu) Draw(screen *ebiten.Image) {
	text.Draw(screen, command, mplusNormalFont, 40, 150, color.White)
	text.Draw(screen, title, mplusTitleFont, 100, 80, color.Gray16{0x4444})
}
