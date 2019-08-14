package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"math/rand"
	"runtime"
	"time"
)

const (
	width  = 1000
	height = 1000

	rows    = 100
	columns = 100

	fps = 10
)

var (
	grid    [][]*Cell
	shaders map[string]uint32
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := initWindow()
	defer glfw.Terminate()

	grid = make([][]*Cell, rows, rows)
	for i := 0; i < columns; i++ {
		grid[i] = make([]*Cell, columns, columns)
	}

	shaders = make(map[string]uint32, 0)
	shaders["vertex"], _ = compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	shaders["cell"], _ = compileShader(cellShaderSource, gl.FRAGMENT_SHADER)
	shaders["food"], _ = compileShader(foodShaderSource, gl.FRAGMENT_SHADER)

	rand.Seed(time.Now().UnixNano())

	cells := make([]*Cell, 10, 10)
	for i := range cells {
		cells[i] = NewLiveCell()
		grid[cells[i].x][cells[i].y] = cells[i]
	}

	for !window.ShouldClose() {

		t := time.Now()

		for _, cell := range cells {
			cell.Tick()
		}

		spawnFood()
		if rand.Float64() < 0.5 {
			spawnFood()
		}

		DrawCanvas(window)

		log.Println("tick in ", time.Since(t).Truncate(time.Millisecond))

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))

	}

}

func spawnFood() {
	f := NewFoodCell()
	grid[f.x][f.y] = f
}
