package main

import (
	"github.com/hajimehoshi/ebiten"
	v2 "github.om/Simulations/v2"
	"log"
)

func main() {
	//ebiten.SetFullscreen(true)
	//ebiten.SetCursorVisible(false)
	//x, y := ebiten.ScreenSizeInFullscreen()
	//s := ebiten.DeviceScaleFactor()

	world := v2.NewWorld()

	err := ebiten.Run(world.Update, 500, 600, 1, "Sim")
	//err := ebiten.Run(world.Update, int(float64(x)*s), int(float64(y)*s), 1/s, "Sim")
	if err != nil {
		log.Println(err)
	}
}
