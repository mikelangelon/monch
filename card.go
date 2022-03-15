package main

import (
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	witdh  = 80
	height = 120
)

var otherSide *ebiten.Image

type card struct {
	number   int
	image    *ebiten.Image
	selected bool
	covered  bool
	used     bool
}

func (c *card) getImage() *ebiten.Image {
	if c.covered {
		return otherSide
	}
	return c.image
}

func (d *hand) nextSelection() *card {
	cards := d.playableCards()
	index := -1
	for i, v := range cards {
		if v.selected {
			index = i
			v.selected = false
		}
	}
	index++
	index %= len(cards)
	cards[index].selected = true
	return cards[index]
}

func (d *hand) cpuSelect() *card {
	playable := d.playableCards()
	r := rand.Intn(len(playable))
	return playable[r]
}

func (d *hand) playableCards() []*card {
	var cards []*card
	for _, v := range d.cards {
		if !v.used {
			cards = append(cards, v)
		}
	}
	return cards
}

type deck struct {
	cards []*card
}

type hand struct {
	cards    []*card
	selected *card
}

func (d *deck) getCard() *card {
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}

func (d *deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *deck) hand(covered bool) *hand {
	var cards []*card
	for i := 1; i <= 3; i++ {
		card := d.getCard()
		card.covered = covered
		cards = append(cards, card)
	}
	return &hand{
		cards: cards,
	}
}

func newDeck(img *ebiten.Image) *deck {
	var cards []*card
	otherSide = img.SubImage(image.Rect(0, height*4, witdh, height*5)).(*ebiten.Image)
	for i := 1; i < 14; i++ {
		cards = append(cards, &card{
			number: i,
			image:  img.SubImage(image.Rect(witdh*(i-1), 0, witdh*i, height)).(*ebiten.Image),
		})
	}
	return &deck{
		cards: cards,
	}
}
