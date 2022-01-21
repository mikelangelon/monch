package main

import "github.com/hajimehoshi/ebiten/v2"

type screen interface {
	Update() error
	Draw(screen *ebiten.Image)
}
