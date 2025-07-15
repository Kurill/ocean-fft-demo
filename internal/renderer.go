package internal

import "github.com/go-gl/gl/v3.3-core/gl"

func InitRenderer() {}
func Update()       {}
func Render() {
	gl.ClearColor(0.0, 0.5, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
func Cleanup() {}
