package v2

import "github.com/hajimehoshi/ebiten"

type Food struct {
	Fertility float64
	pos Vector
	image *ebiten.Image
}

func (f *Food) Image() *ebiten.Image {
	return f.image
}

func (f *Food) Pos() Vector {
	return f.pos
}