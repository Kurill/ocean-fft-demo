package internal

import "github.com/go-gl/gl/v3.3-core/gl"

// Mesh represents a basic vertex/index buffer combination.
type Mesh struct {
	vao        uint32
	vbo        uint32
	ebo        uint32
	indexCount int32
}

// NewGrid generates a simple NxN grid on the XZ plane in the range [-1,1].
func NewGrid(n int) *Mesh {
	step := float32(2.0 / float32(n))

	vertices := make([]float32, 0, (n+1)*(n+1)*3)
	for i := 0; i <= n; i++ {
		z := -1.0 + float32(i)*step
		for j := 0; j <= n; j++ {
			x := -1.0 + float32(j)*step
			vertices = append(vertices, x, 0.0, z)
		}
	}

	indices := make([]uint32, 0, n*n*6)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			start := uint32(i*(n+1) + j)
			indices = append(indices,
				start, start+1, start+uint32(n+1),
				start+1, start+uint32(n+2), start+uint32(n+1))
		}
	}

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	gl.BindVertexArray(0)

	return &Mesh{vao: vao, vbo: vbo, ebo: ebo, indexCount: int32(len(indices))}
}

// Draw renders the mesh using the currently bound shader program.
func (m *Mesh) Draw() {
	gl.BindVertexArray(m.vao)
	gl.DrawElements(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}

// Delete frees the OpenGL resources associated with the mesh.
func (m *Mesh) Delete() {
	gl.DeleteVertexArrays(1, &m.vao)
	gl.DeleteBuffers(1, &m.vbo)
	gl.DeleteBuffers(1, &m.ebo)
}
