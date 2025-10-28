package main

import "github.com/ebitengine/purego"

var OpenGLPath string = "/System/Library/Frameworks/OpenGL.framework/OpenGL"

func LoadOpenGL() {
	_, err := purego.Dlopen(OpenGLPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	println("Done")
}
