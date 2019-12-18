package v1

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"time"
)

/*const (
	width  = 1000
	height = 1000

	rows    = 100
	columns = 100

	fps = 10
)*/

type Config struct {
	Width    int     `yaml:"width"`
	Height   int     `yaml:"height"`
	Rows     int     `yaml:"rows"`
	Columns  int     `yaml:"columns"`
	Cells    int     `yaml:"cells"`
	FoodRate float64 `yaml:"fertility"`
	Fps      int     `yaml:"fps"`
}

var (
	grid    [][]*Cell
	shaders map[string]uint32

	//default
	width  = 1000
	height = 1000

	rows    = 100
	columns = 100

	cells = 10

	fps = 10
)

func init() {
	runtime.LockOSThread()
}

func Configure(fileName string) (Config, error) {
	var cnf Config

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return Config{}, err
	}

	return cnf, nil
}

func main() {

	conf, err := Configure("./config.yaml")
	if err != nil {
		log.Println("can't configure app: ", err)
		log.Println("using default settings")
	} else {
		width = conf.Width
		height = conf.Height
		rows = conf.Rows
		columns = conf.Columns
		fps = conf.Fps
	}

	window := initWindow()
	defer glfw.Terminate()

	grid = make([][]*Cell, rows, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]*Cell, columns, columns)
	}

	shaders = make(map[string]uint32, 0)
	shaders["vertex"], _ = compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	shaders["cell"], _ = compileShader(cellShaderSource, gl.FRAGMENT_SHADER)
	shaders["food"], _ = compileShader(foodShaderSource, gl.FRAGMENT_SHADER)

	rand.Seed(time.Now().UnixNano())

	cells := make([]*Cell, conf.Cells, conf.Cells)
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
		if rand.Float64() < conf.FoodRate {
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
