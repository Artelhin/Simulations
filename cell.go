package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"math"
	"math/rand"
)

type Cell struct {
	Drawable
	isFood bool
}

func NewLiveCell() *Cell {
	c := new(Cell)
	c.prog = gl.CreateProgram()
	gl.AttachShader(c.prog, shaders["cell"])
	gl.AttachShader(c.prog, shaders["vertex"])
	gl.LinkProgram(c.prog)
	c.x = rand.Intn(rows)
	c.y = rand.Intn(columns)
	c.isFood = false
	return c
}

func NewFoodCell() *Cell {
	c := new(Cell)
	c.prog = gl.CreateProgram()
	gl.AttachShader(c.prog, shaders["food"])
	gl.AttachShader(c.prog, shaders["vertex"])
	gl.LinkProgram(c.prog)
	c.x = rand.Intn(rows)
	c.y = rand.Intn(columns)
	c.isFood = true
	return c
}

//returns direction to nearest food and distance to it
//0,0,0 if not food nearby
func (c *Cell) checkFoodNearby() (int, int, int) {
	type Coords struct {
		x, y int
	}
	cx := c.x
	cy := c.y
	toCheck := []Coords{
		{0, 3},
		{-1, 2},
		{0, 2},
		{1, 2},
		{-2, 1},
		{-1, 1},
		{0, 1},
		{1, 1},
		{2, 1},
		{-3, 0},
		{-2, 0},
		{-1, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{-2, -1},
		{-1, -1},
		{0, -1},
		{1, -1},
		{2, -1},
		{-1, -2},
		{0, -2},
		{1, -2},
		{0, -3},
	}
	for _, k := range toCheck {
		absx := (rows + cx + k.x) % rows
		absy := (columns + cy + k.y) % columns
		if checkFood(absx, absy) {
			//log.Printf("found food in %d %d", k.x, k.y)
			return k.x, k.y, int(math.Abs(float64(k.x * k.y)))
		}
	}
	return 0, 0, 0
}

func checkFood(x, y int) bool {
	if grid[x][y] == nil || !grid[x][y].isFood {
		return false
	}
	return true
}

func (c *Cell) Tick() {
	x, y, _ := c.checkFoodNearby()

	move := func(x, y int) {
		grid[c.x][c.y] = nil
		c.x = (rows + c.x + x) % rows
		c.y = (columns + c.y + y) % columns
		grid[c.x][c.y] = c
	}

	if x == 0 && y == 0 {
		move(rand.Intn(3)-1, rand.Intn(3)-1)
		return
	}
	if math.Abs(float64(x)) >= math.Abs(float64(y)) {
		move(x,0)
	} else {
		move(0,y)
	}
}
