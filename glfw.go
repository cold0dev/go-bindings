package main

var GLFW_SAMPLES int = 0x0002100D
var GLFW_CONTEXT_VERSION_MAJOR int = 0x00022002
var GLFW_CONTEXT_VERSION_MINOR int = 0x00022003
var GLFW_OPENGL_FORWARD_COMPAT int = 0x00022006
var GLFW_OPENGL_PROFILE int = 0x00022008
var GLFW_OPENGL_CORE_PROFILE int = 0x00032001

var GlfwInit func() int
var GlfwWindowHint func(hint int, value int)
var GlfwCreateWindow func(width int, height int, title string, monitor uintptr, share uintptr) uintptr
var GlfwTerminate func()
var GlfwMakeContextCurrent func(window uintptr)
var GlfwSetInputMode func(

