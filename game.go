package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	lastClickAt  time.Time
	lastCPU      time.Time
	turn         string
	playerPoints int64
	cpuPoints    int64
}

const (
	cpuTurn    = "cpu"
	playerTurn = "player"
	endTurn    = "end"
)

const debouncer = 150 * time.Millisecond

func (g *Game) Restart() {
	g.lastClickAt = time.Now()
	g.playerPoints = 0
	g.cpuPoints = 0
	g.turn = playerTurn
}
func (g *Game) Update() error {

	switch g.turn {
	case playerTurn:
		if time.Since(g.lastClickAt) < debouncer {
			return nil
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			player.selected = player.nextSelection()
			g.lastClickAt = time.Now()
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			player.selected = player.nextSelection()
			g.lastClickAt = time.Now()
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			log.Printf("selecting card with number %d", player.selected.number)

			g.turn = cpuTurn
			g.lastCPU = time.Now()
			g.lastClickAt = time.Now()
		}

	case cpuTurn:
		if time.Since(g.lastCPU) > 1*time.Second {
			selected := rival.cpuSelect()
			selected.selected = true
			selected.covered = false
			rival.selected = selected

			// Check who wins this match
			if rival.selected.number > player.selected.number {
				g.cpuPoints++
			} else {
				g.playerPoints++
			}
			// use selected cards
			log.Println(player.selected.number)
			player.selected.used = true
			player.selected = nil

			rival.selected.used = true
			rival.selected = nil

			if len(player.playableCards()) == 0 {
				g.turn = endTurn
			} else {
				g.turn = playerTurn
			}
		}
	case endTurn:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			initGame()
			g.Restart()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, v := range player.cards {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(640/2-witdh/2+(witdh+10)*(i-1)), 300)
		if v.selected {
			op.ColorM.Scale(2, 2, 0, 1)
		}
		if v.used {
			op.ColorM.Scale(0, 1, 1, 1)
		}
		screen.DrawImage(v.getImage(), op)

	}
	for i, v := range rival.cards {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(640/2-witdh/2+(witdh+10)*(i-1)), 40)
		if v.selected {
			op.ColorM.Scale(2, 2, 0, 1)
		}
		if v.used {
			op.ColorM.Scale(0, 1, 1, 1)
		}

		screen.DrawImage(v.getImage(), op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", len(d.cards)))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player: %d vs CPU: %d - turn %s", g.playerPoints, g.cpuPoints, g.turn))
}
