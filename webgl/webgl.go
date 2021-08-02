package webgl

import (
	"errors"
	"syscall/js"
)

const (
	COLOR_BUFFER_BIT = 0x00004000
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
