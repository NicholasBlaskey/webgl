# webgl

### Code gen

The constants are generated by downloading [this page](https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Constants) and parsing the html.

This can be done by running the command

```
go run codeGen/codeGen.go > webgl/constants.go
```
