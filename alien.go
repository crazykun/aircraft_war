package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Alien represents an alien.
// Alien结构体表示外星人
type Alien struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./images/alien.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		image: img,
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      0,
			y:      0,
		},
		speedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) outOfScreen(cfg *Config) bool {
	return alien.y > float64(cfg.ScreenHeight)
}
