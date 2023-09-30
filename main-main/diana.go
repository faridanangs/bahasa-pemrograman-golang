package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	angle float64
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.White)

	angle += 0.05

	x := float64(screenWidth)/2 + 100*math.Cos(angle)
	y := float64(screenHeight)/2 + 100*math.Sin(angle)

	ebitenutil.DrawRect(screen, x, y, 50, 50, color.RGBA{255, 0, 0, 255})

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Go Animation"); err != nil {
		log.Fatal(err)
	}
}
