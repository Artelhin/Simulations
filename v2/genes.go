package v2

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"math"
)

type Genes struct {
	VisionRange float64
	Speed       float64
	Size        float64
}

func (g *Genes) CalculateEnergyUsage() float64 {
	return g.VisionRange + 10*math.Pow(g.Speed, 2) + 10*math.Pow(g.Size, 3)
}

func (g *Genes) GenImage() *ebiten.Image {
	img, _ := ebiten.NewImage(int(g.Size), int(g.Size), ebiten.FilterNearest)
	img.Fill(color.NRGBA{uint8(g.VisionRange), uint8(g.Speed), uint8(g.Size), 1})
	log.Println("rgb: ", uint8(g.VisionRange), uint8(g.Speed), uint8(g.Size))
	return img
}
