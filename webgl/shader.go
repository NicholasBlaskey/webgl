package webgl

import (
	"syscall/js"
)

type Program struct {
	JsProgram js.Value
}

type Shader struct {
	JsShader js.Value
}

func (gl *Gl) CreateProgram() *Program {
	prog := gl.JsGl.Call("createProgram")
	if prog.IsUndefined() {
		return nil
	}
	return &Program{prog}
}

func (gl *Gl) CreateShader(shaderType int) *Shader {
	shader := gl.JsGl.Call("createShader", shaderType)
	if shader.IsUndefined() {
		return nil
	}
	return &Shader{shader}
}

func (gl *Gl) AttachShader(program *Program, shader *Shader) {
	gl.JsGl.Call("attachShader", program.JsProgram, shader.JsShader)
}

func (gl *Gl) LinkProgram(program *Program) {
	gl.JsGl.Call("linkProgram", program.JsProgram)
}

func (gl *Gl) GetProgramParameter(program *Program, paramType int) int {
	param := gl.JsGl.Call("getProgramParameter", program.JsProgram, paramType)
	return jsIntOrBoolToInt(param)
}

func (gl *Gl) GetProgramInfoLog(program *Program) string {
	return gl.JsGl.Call("getProgramInfoLog").String()
}

func (gl *Gl) DeleteProgram(program *Program) {
	gl.JsGl.Call("deleteProgram", program.JsProgram)
}

func (gl *Gl) DeleteShader(shader *Shader) {
	gl.JsGl.Call("deleteShader", shader.JsShader)
}

func (gl *Gl) ShaderSource(shader *Shader, source string) {
	gl.JsGl.Call("shaderSource", shader.JsShader, source)
}

func (gl *Gl) CompileShader(shader *Shader) {
	gl.JsGl.Call("compileShader", shader.JsShader)
}

func (gl *Gl) GetShaderParameter(shader *Shader, paramType int) int {
	param := gl.JsGl.Call("getShaderParameter", shader.JsShader, paramType)
	return jsIntOrBoolToInt(param)
}

func (gl *Gl) GetShaderInfoLog(shader *Shader) string {
	return gl.JsGl.Call("getShaderInfoLog", shader.JsShader).String()
}

func (gl *Gl) UseProgram(program *Program) {
	gl.JsGl.Call("useProgram", program.JsProgram)
}

func jsIntOrBoolToInt(x js.Value) int {
	if x.Type() == js.TypeBoolean {
		if x.Bool() {
			return 1
		}
		return 0
	}
	return x.Int()
}
