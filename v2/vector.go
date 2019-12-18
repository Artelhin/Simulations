package v2

import "math"

type Vector struct {
	X, Y       float64
	Normalized bool
}

func (v *Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	if !v.Normalized {
		v.X = v.X / v.Len()
		v.Y = v.Y / v.Len()
		v.Normalized = true
	}
}

func NewVector(x, y float64) Vector {
	return Vector{x, y, false}
}

func NewNormalizedVector(x, y float64) Vector {
	v := Vector{x, y, false}
	v.Normalize()
	return v
}
