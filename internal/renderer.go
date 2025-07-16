package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	oceanMesh     *Mesh
	shaderProgram uint32
)

// InitRenderer prepares OpenGL state, compiles shaders and generates the grid mesh.
func InitRenderer() {
	gl.Enable(gl.DEPTH_TEST)

	vs, err := loadShaderSource("shaders/ocean.vert")
	if err != nil {
		log.Fatalf("failed to load vertex shader: %v", err)
	}
	fs, err := loadShaderSource("shaders/ocean.frag")
	if err != nil {
		log.Fatalf("failed to load fragment shader: %v", err)
	}

	shaderProgram, err = newProgram(vs, fs)
	if err != nil {
		log.Fatalf("failed to compile shaders: %v", err)
	}

	oceanMesh = NewGrid(128)
}

// Update is a placeholder for future per-frame updates.
func Update() {}

// Render draws the ocean mesh using a very basic shader.
func Render() {
	gl.ClearColor(0.0, 0.5, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(shaderProgram)
	oceanMesh.Draw()
}

// Cleanup releases OpenGL resources.
func Cleanup() {
	if oceanMesh != nil {
		oceanMesh.Delete()
	}
	gl.DeleteProgram(shaderProgram)
}

func loadShaderSource(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("compile shader: %v", log)
	}
	return shader, nil
}

func newProgram(vertexSrc, fragmentSrc string) (uint32, error) {
	v, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	f, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		gl.DeleteShader(v)
		return 0, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, v)
	gl.AttachShader(program, f)
	gl.LinkProgram(program)

	gl.DeleteShader(v)
	gl.DeleteShader(f)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("link program: %v", log)
	}

	return program, nil
}
