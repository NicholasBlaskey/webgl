package webgl

import (
	"syscall/js"
)

type Texture struct {
	JsTexture js.Value
}

func (gl *Gl) CreateTexture() *Texture {
	t := gl.JsGl.Call("createTexture")
	if t.IsUndefined() {
		return nil
	}
	return &Texture{t}
}

func (gl *Gl) DeleteTexture(texture *Texture) {
	gl.JsGl.Call("deleteTexture", texture.JsTexture)
}

func (gl *Gl) BindTexture(target int, texture *Texture) {
	gl.JsGl.Call("bindTexture", target, texture.JsTexture)
}

func (gl *Gl) ActiveTexture(texUnit int) {
	gl.JsGl.Call("activeTexture", texUnit)
}

func (gl *Gl) TexImage2DArray(target, level, internalFormat,
	width, height, border, format, texType int, pixels interface{}) {

	gl.JsGl.Call("texImage2D", target, level, internalFormat,
		width, height, border, format, texType, toTypedArray(pixels))
}

func (gl *Gl) TexImage2D(target, level, internalFormat, format, texType int, pixels interface{}) {
	gl.JsGl.Call("texImage2D", target, level, internalFormat, format, texType, pixels)
}

func (gl *Gl) TexParameteri(target, pname, param int) {
	gl.JsGl.Call("texParameteri", target, pname, param)
}
