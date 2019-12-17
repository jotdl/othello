package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	ebiten.SetMaxTPS(40)
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! %v", ebiten.CurrentFPS()))
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
