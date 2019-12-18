package v2

import "github.com/hajimehoshi/ebiten"

type Agent struct {
	Facing Vector
	pos    Vector
	Genom  Genes

	HungerDecreaseRate float64
	Hunger             float64
	image              *ebiten.Image
}

func (a *Agent) Tick() bool {
	a.Hunger = a.Hunger - a.HungerDecreaseRate
	return a.Hunger > 0
}

func (a *Agent) Image() *ebiten.Image {
	return a.image
}

func (a *Agent) Pos() Vector {
	return a.pos
}