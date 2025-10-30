package main

import "github.com/ebitengine/purego"

var GLFW_SAMPLES int = 0x0002100D
var GLFW_CONTEXT_VERSION_MAJOR int = 0x00022002
var GLFW_CONTEXT_VERSION_MINOR int = 0x00022003
var GLFW_OPENGL_FORWARD_COMPAT int = 0x00022006
var GLFW_OPENGL_PROFILE int = 0x00022008
var GLFW_OPENGL_CORE_PROFILE int = 0x00032001
var GLFW_STICKY_KEYS int = 0x00033002
var GLFW_KEY_ESCAPE int = 256
var GLFW_PRESS int = 1


var GlfwInit func() int
var GlfwWindowHint func(hint int, value int)
var GlfwCreateWindow func(width int, height int, title string, monitor uintptr, share uintptr) uintptr
var GlfwTerminate func()
var GlfwMakeContextCurrent func(window uintptr)
var GlfwSetInputMode func(window uintptr, mode int, value int)
var GlfwSwapBuffers func (window uintptr)
var GlfwPollEvents func()
var GlfwGetKey func(window uintptr, key int) int
var GlfwWindowShouldClose func(window uintptr) int

var GLFWPath string = "/usr/lib/libglfw.so"

func LoadGLFW() {
	lib, err := purego.Dlopen(GLFWPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&GlfwInit, lib, "glfwInit")
	purego.RegisterLibFunc(&GlfwWindowHint, lib, "glfwWindowHint")
	purego.RegisterLibFunc(&GlfwCreateWindow, lib, "glfwCreateWindow")
	purego.RegisterLibFunc(&GlfwTerminate, lib, "glfwTerminate")
	purego.RegisterLibFunc(&GlfwMakeContextCurrent, lib, "glfwMakeContextCurrent")
	purego.RegisterLibFunc(&GlfwSetInputMode, lib, "glfwSetInputMode")
	purego.RegisterLibFunc(&GlfwSwapBuffers, lib, "glfwSwapBuffers")
	purego.RegisterLibFunc(&GlfwPollEvents, lib, "glfwPollEvents")
	purego.RegisterLibFunc(&GlfwGetKey, lib, "glfwGetKey")
	purego.RegisterLibFunc(&GlfwWindowShouldClose, lib, "glfwWindowShouldClose")
}
