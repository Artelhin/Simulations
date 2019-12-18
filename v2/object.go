package v2

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Image() *ebiten.Image
	Pos() Vector
}