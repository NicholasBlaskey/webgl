package webgl

import (
	"errors"
	"syscall/js"
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

func (gl *Gl) DrawElements(mode, count, dataType, offset int) {
	gl.JsGl.Call("drawElements", mode, count, dataType, offset)
}

func (gl *Gl) Viewport(x, y, width, height int) {
	gl.JsGl.Call("viewport", x, y, width, height)
}

func (gl *Gl) Enable(mask int) {
	gl.JsGl.Call("enable", mask)
}

func (gl *Gl) Disable(mask int) {
	gl.JsGl.Call("disable", mask)
}

func (gl *Gl) Scissor(x, y, width, height int) {
	gl.JsGl.Call("scissor", x, y, width, height)
}
