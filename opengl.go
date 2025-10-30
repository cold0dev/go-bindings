package main

import "github.com/ebitengine/purego"

var OpenGLPathMac string = "/System/Library/Frameworks/OpenGL.framework/OpenGL"
var OpenGLPathLinux string = "/usr/lib/libOpenGL.so"

type GLbitfield uint

var GL_COLOR_BUFFER_BIT GLbitfield = 0x00004000
var GL_DEPTH_BUFFER_BIT GLbitfield = 0x00000100
var GL_TRUE int = 1

var GlClear func(mask GLbitfield)

func LoadOpenGL() {
	lib, err := purego.Dlopen(OpenGLPathLinux, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&GlClear, lib, "glClear")
}
