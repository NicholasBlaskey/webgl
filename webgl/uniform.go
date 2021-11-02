package webgl

import (
	"reflect"
	"runtime"
	"unsafe"

	"syscall/js"
)

var jsBuffer js.Value
var jsArray js.Value

func init() {
	jsBuffer = js.Global().Get("Uint8Array").New(4 * 4 * 4)

	buf := jsBuffer.Get("buffer")
	jsArray = js.Global().Get("Float32Array").New(buf,
		0,  // byteOffset
		16, // byteLength / 4
	)
}

type UniformLocation struct {
	JsUniformLocation js.Value
}

func (gl *Gl) GetUniformLocation(program *Program, uniformName string) *UniformLocation {
	ul := gl.JsGl.Call("getUniformLocation", program.JsProgram, uniformName)
	if ul.IsUndefined() {
		return nil
	}
	return &UniformLocation{ul}
}

// 1f
func (gl *Gl) Uniform1f(loc *UniformLocation, x float32) {
	gl.JsGl.Call("uniform1f", loc.JsUniformLocation, x)
}

func (gl *Gl) Uniform1fv(loc *UniformLocation, x []float32) {
	gl.JsGl.Call("uniform1fv", loc.JsUniformLocation, toTypedArray(x))
}

// 2f
func (gl *Gl) Uniform2f(loc *UniformLocation, x, y float32) {
	gl.JsGl.Call("uniform2f", loc.JsUniformLocation, x, y)
}

func (gl *Gl) Uniform2fv(loc *UniformLocation, x []float32) {
	gl.JsGl.Call("uniform2fv", loc.JsUniformLocation, toTypedArray(x))
}

// 3f
func (gl *Gl) Uniform3f(loc *UniformLocation, x, y, z float32) {
	gl.JsGl.Call("uniform3f", loc.JsUniformLocation, x, y, z)
}

func (gl *Gl) Uniform3fv(loc *UniformLocation, x []float32) {
	gl.JsGl.Call("uniform3fv", loc.JsUniformLocation, toTypedArray(x))
}

// 4f
func (gl *Gl) Uniform4f(loc *UniformLocation, x, y, z, w float32) {
	gl.JsGl.Call("uniform4f", loc.JsUniformLocation, x, y, z, w)
}

func (gl *Gl) Uniform4fv(loc *UniformLocation, x []float32) {
	gl.JsGl.Call("uniform4fv", loc.JsUniformLocation, toTypedArray(x))
}

// 1i
func (gl *Gl) Uniform1i(loc *UniformLocation, x int) {
	gl.JsGl.Call("uniform1i", loc.JsUniformLocation, x)
}

func (gl *Gl) Uniform1iv(loc *UniformLocation, x []int) {
	gl.JsGl.Call("uniform1iv", loc.JsUniformLocation, toTypedArray(x))
}

// 2i
func (gl *Gl) Uniform2i(loc *UniformLocation, x, y int) {
	gl.JsGl.Call("uniform2i", loc.JsUniformLocation, x, y)
}

func (gl *Gl) Uniform2iv(loc *UniformLocation, x []int) {
	gl.JsGl.Call("uniform2iv", loc.JsUniformLocation, toTypedArray(x))
}

// 3i
func (gl *Gl) Uniform3i(loc *UniformLocation, x, y, z int) {
	gl.JsGl.Call("uniform3i", loc.JsUniformLocation, x, y, z)
}

func (gl *Gl) Uniform3iv(loc *UniformLocation, x []int) {
	gl.JsGl.Call("uniform3iv", loc.JsUniformLocation, toTypedArray(x))
}

// 4i
func (gl *Gl) Uniform4i(loc *UniformLocation, x, y, z, w int) {
	gl.JsGl.Call("uniform4i", loc.JsUniformLocation, x, y, z, w)
}

func (gl *Gl) Uniform4iv(loc *UniformLocation, x []int) {
	gl.JsGl.Call("uniform4iv", loc.JsUniformLocation, toTypedArray(x))
}

// Mat2
func (gl *Gl) UniformMatrix2fv(loc *UniformLocation, transpose bool, x []float32) {
	gl.JsGl.Call("uniformMatrix2fv", loc.JsUniformLocation, transpose, toTypedArray(x))
}

// Mat3
func (gl *Gl) UniformMatrix3fv(loc *UniformLocation, transpose bool, x []float32) {
	gl.JsGl.Call("uniformMatrix3fv", loc.JsUniformLocation, transpose, toTypedArray(x))
}

// Mat4
// TODO do this to all the other matrix types.
// Optimized for speed.
func (gl *Gl) UniformMatrix4fv(loc *UniformLocation, transpose bool, x []float32) {
	runtime.KeepAlive(x) // Do we need this?
	h := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	h.Len *= 4
	h.Cap *= 4
	byteSlice := *(*[]byte)(unsafe.Pointer(h))
	js.CopyBytesToJS(jsBuffer, byteSlice)

	jsArray.Set("buffer", jsBuffer.Get("buffer"))
	gl.JsGl.Call("uniformMatrix4fv", loc.JsUniformLocation, transpose, jsArray)
}
