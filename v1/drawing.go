package v1

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"strings"
)

const (
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	cellShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"

	foodShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(0, 1, 0, 1);
    }
` + "\x00"
)

type Drawable struct {
	prog uint32
	x    int
	y    int
}

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func DrawCanvas(window *glfw.Window) {
	vaos := make([]uint32, 0)
	vbos := make([]uint32, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == nil {
				continue
			}
			vao, vbo := grid[x][y].Draw()
			vaos = append(vaos, vao)
			vbos = append(vbos, vbo)
		}
	}

	Render(window)

	for _, v := range vaos {
		gl.DeleteVertexArrays(1, &v)
	}
	for _, v := range vbos {
		gl.DeleteBuffers(1, &v)
	}
	log.Println("deleted ", len(vaos), " ", len(vbos))
}

func (d *Drawable) Draw() (uint32, uint32) {
	gl.UseProgram(d.prog)
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(columns)
			position = float32(d.x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(d.y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	vao, vbo := makeVao(points)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
	return vao, vbo
}

func makeVao(points []float32) (uint32, uint32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao, vbo
}

func Render(window *glfw.Window) {
	glfw.PollEvents()
	window.SwapBuffers()
}

func initWindow() *glfw.Window {

	//init glfw and create a window
	if err := glfw.Init(); err != nil {
		log.Fatal("can't init glfw: ", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Random stuff", nil, nil)
	if err != nil {
		log.Fatal("can't create window: ", err)
	}
	window.MakeContextCurrent()

	//init OpenGL
	if err := gl.Init(); err != nil {
		log.Fatal("can't init opengl", err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	return window
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		lg := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(lg))

		return 0, fmt.Errorf("failed to compile %v: %v", source, lg)
	}

	return shader, nil
}
