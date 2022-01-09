package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"log"
)

//go:embed pixel_style1.png
var b []byte

var coins *ebiten.Image
var buttons *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	buttons, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the sub-image. Only the red part should be rendered.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 100)
	screen.DrawImage(buttons.SubImage(image.Rect(16*5, 0, 16*6, 16)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
