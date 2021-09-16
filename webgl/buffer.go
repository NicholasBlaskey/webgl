package webgl

import (
	"syscall/js"

	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type Buffer struct {
	JsBuffer js.Value
}

func (gl *Gl) CreateBuffer() *Buffer {
	buff := gl.JsGl.Call("createBuffer")
	if buff.IsUndefined() {
		return nil
	}
	return &Buffer{buff}
}

func (gl *Gl) DeleteBuffer(buff *Buffer) {
	gl.JsGl.Call("deleteBuffer", buff.JsBuffer)
}

func (gl *Gl) BindBuffer(buffType int, buff *Buffer) {
	gl.JsGl.Call("bindBuffer", buffType, buff.JsBuffer)
}

func (gl *Gl) BufferData(buffType int, data interface{}, drawHint int) {
	gl.JsGl.Call("bufferData", buffType, toTypedArray(data), drawHint)
}

func (gl *Gl) VertexAttribPointer(index, size, dataType int, normalized bool, stride, offset int) {
	gl.JsGl.Call("vertexAttribPointer", index, size, dataType, normalized, stride, offset)
}

func (gl *Gl) EnableVertexAttribArray(attribLoc int) {
	gl.JsGl.Call("enableVertexAttribArray", attribLoc)
}

func (gl *Gl) GetAttribLocation(program *Program, attribute string) int {
	return gl.JsGl.Call("getAttribLocation", program.JsProgram, attribute).Int()
}

// This shouldn't be this hard.
// https://github.com/golang/go/issues/32402
func sliceToByteSlice(s interface{}) []byte {
	switch s := s.(type) {
	case []int8:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		return *(*[]byte)(unsafe.Pointer(h))
	case []int16:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 2
		h.Cap *= 2
		return *(*[]byte)(unsafe.Pointer(h))
	case []int32:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 4
		h.Cap *= 4
		return *(*[]byte)(unsafe.Pointer(h))
	case []int64:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 8
		h.Cap *= 8
		return *(*[]byte)(unsafe.Pointer(h))
	case []uint8:
		return s
	case []uint16:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 2
		h.Cap *= 2
		return *(*[]byte)(unsafe.Pointer(h))
	case []uint32:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 4
		h.Cap *= 4
		return *(*[]byte)(unsafe.Pointer(h))
	case []uint64:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 8
		h.Cap *= 8
		return *(*[]byte)(unsafe.Pointer(h))
	case []float32:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 4
		h.Cap *= 4
		return *(*[]byte)(unsafe.Pointer(h))
	case []float64:
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		h.Len *= 8
		h.Cap *= 8
		return *(*[]byte)(unsafe.Pointer(h))
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at sliceToBytesSlice: %T", s))
	}
}

func toTypedArray(s interface{}) js.Value {
	switch s := s.(type) {
	case []int8:
		a := js.Global().Get("Uint8Array").New(len(s))
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Int8Array").New(buf, a.Get("byteOffset"), a.Get("byteLength"))
	case []int16:
		a := js.Global().Get("Uint8Array").New(len(s) * 2)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Int16Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/2)
	case []int32:
		a := js.Global().Get("Uint8Array").New(len(s) * 4)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Int32Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/4)
	case []uint8:
		a := js.Global().Get("Uint8Array").New(len(s))
		js.CopyBytesToJS(a, s)
		runtime.KeepAlive(s)
		return a
	case []uint16:
		a := js.Global().Get("Uint8Array").New(len(s) * 2)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Uint16Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/2)
	case []uint32:
		a := js.Global().Get("Uint8Array").New(len(s) * 4)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Uint32Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/4)
	case []float32:
		a := js.Global().Get("Uint8Array").New(len(s) * 4)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Float32Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/4)
	case []float64:
		a := js.Global().Get("Uint8Array").New(len(s) * 8)
		js.CopyBytesToJS(a, sliceToByteSlice(s))
		runtime.KeepAlive(s)
		buf := a.Get("buffer")
		return js.Global().Get("Float64Array").New(buf, a.Get("byteOffset"), a.Get("byteLength").Int()/8)
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}
}
