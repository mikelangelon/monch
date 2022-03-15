package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed deck.png
var deckImage []byte

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	d      *deck
	player *hand
	rival  *hand
)

func init() {
	initGame()
}

func initGame() {
	img, _, err := image.Decode(bytes.NewReader(deckImage))
	if err != nil {
		log.Fatal(err)
	}

	d = newDeck(ebiten.NewImageFromImage(img))
	d.Shuffle()

	player = d.hand(false)
	rival = d.hand(true)
}

type app struct {
	currentScreen screen
}

func (a *app) Update() error {
	return a.currentScreen.Update()
}

func (a *app) Draw(screen *ebiten.Image) {
	a.currentScreen.Draw(screen)
}

func (a *app) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (a *app) changeScreen(s screen) {
	a.currentScreen = s
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World!")

	app := &app{}
	app.currentScreen = &menu{changeScreen: app.changeScreen}
	if err := ebiten.RunGame(app); err != nil {
		panic(err)
	}
}
