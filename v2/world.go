package v2

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type World struct {
	Objects []Object
	Agents  []*Agent
}

const Hunger = 10000

func NewWorld() *World {
	agent := Agent{
		Facing: NewNormalizedVector(1, 0),
		pos:    Vector{250, 300, false},
		Genom: Genes{
			VisionRange: 200,
			Speed:       2,
			Size:        10,
		},
		Hunger: Hunger,
	}
	agent.HungerDecreaseRate = agent.Genom.CalculateEnergyUsage()
	agent.image = agent.Genom.GenImage()

	world := World{
		Objects: append(make([]Object, 0), &agent),
		Agents:  append(make([]*Agent, 0), &agent),
	}
	return &world
}

func (w *World) Update(screen *ebiten.Image) error {

	for _, a := range w.Agents {
		//must be normalized
		log.Printf("%f + %f", a.pos.X, a.Genom.Speed)
		a.pos.X += a.Facing.X * a.Genom.Speed
		a.pos.Y += a.Facing.Y * a.Genom.Speed
	}

	for _, o := range w.Objects {
		//log.Printf("draw at %f %f", o.Pos().X, o.Pos().Y)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(o.Pos().X, o.Pos().Y)
		screen.DrawImage(o.Image(), opts)
	}

	ebitenutil.DebugPrint(screen, "working\nkinda")

	return nil
}
