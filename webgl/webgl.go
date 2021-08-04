package webgl

import (
	"errors"
	"syscall/js"
)

const (
	COLOR_BUFFER_BIT = 0x00004000

	FRAGMENT_SHADER = 0x8B30
	VERTEX_SHADER   = 0x8B31

	LINK_STATUS    = 0x8B82
	COMPILE_STATUS = 0x8B81

	POINTS         = 0x0000
	LINES          = 0x0001
	LINE_LOOP      = 0x0002
	LINE_STRIP     = 0x0003
	TRIANGLES      = 0x0004
	TRIANGLE_STRIP = 0x0005
	TRIANGLE_FAN   = 0x0006

	ARRAY_BUFFER         = 0x8892
	ELEMENT_ARRAY_BUFFER = 0x8893
	STATIC_DRAW          = 0x88E4

	FLOAT = 0x1406
)

type Gl struct {
	JsGl js.Value
}

func FromCanvas(canvas js.Value) (*Gl, error) {
	gl := canvas.Call("getContext", "webgl")
	if gl.IsUndefined() {
		return nil, errors.New("Unable to create webgl rendering context")
	}

	return &Gl{gl}, nil
}

func (gl *Gl) ClearColor(r, g, b, a float32) {
	gl.JsGl.Call("clearColor", r, g, b, a)
}

func (gl *Gl) Clear(mask int) {
	gl.JsGl.Call("clear", mask)
}

func (gl *Gl) DrawArrays(mode, first, count int) {
	gl.JsGl.Call("drawArrays", mode, first, count)
}
