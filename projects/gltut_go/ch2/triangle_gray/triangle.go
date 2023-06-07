package main

import (
	"fmt"
	"os"
	"runtime"
	// "unsafe"

	"example.com/gltut_go/gl"
	"example.com/gltut_go/sdl2"
)

var vertices = []float32{
	0.0, 0.5, 1.0, // Vertex 1 (X, Y) : Red
	0.5, -0.5, 0.5, // Vertex 2 (X, Y) : Green
	-0.5, -0.5, 0, // Vertex 3 (X, Y) : Blue
}

func main() {

	runtime.LockOSThread()

	sdl2.SDL_Init(sdl2.SDL_INIT_VIDEO)

	sdl2.SDL_GL_SetAttribute(sdl2.SDL_GL_CONTEXT_PROFILE_MASK, int(sdl2.SDL_GL_CONTEXT_PROFILE_CORE))
	sdl2.SDL_GL_SetAttribute(sdl2.SDL_GL_CONTEXT_MAJOR_VERSION, 3)
	sdl2.SDL_GL_SetAttribute(sdl2.SDL_GL_CONTEXT_MINOR_VERSION, 2)
	sdl2.SDL_GL_SetAttribute(sdl2.SDL_GL_STENCIL_SIZE, 8)

	var window *sdl2.SDL_Window = sdl2.SDL_CreateWindow(
		"OpenGL",
		sdl2.SDL_WINDOWPOS_UNDEFINED, sdl2.SDL_WINDOWPOS_UNDEFINED,
		800, 600,
		(uint32)(sdl2.SDL_WINDOW_OPENGL),
	)

	var context sdl2.SDL_GLContext = sdl2.SDL_GL_CreateContext(window)

	gl.GlewInit(true)

	var vao gl.GLuint
	gl.GlGenVertexArrays1(&vao)
	checkGlError()
	gl.GlBindVertexArray(vao)
	checkGlError()

	var vbo gl.GLuint
	gl.GlGenBuffers1(&vbo)
	gl.GlBindBuffer(gl.GL_ARRAY_BUFFER, vbo)
	gl.GlBufferData_f(gl.GL_ARRAY_BUFFER, len(vertices), vertices, gl.GL_STATIC_DRAW)

	checkGlError()

	var vert_shader = createShader(gl.GL_VERTEX_SHADER, "shader.vert")
	var frag_shader = createShader(gl.GL_FRAGMENT_SHADER, "shader.frag")

	fmt.Printf("Vertex shader: %v\n", vert_shader)
	fmt.Printf("Fragment shader: %v\n", frag_shader)

	var prog = gl.GlCreateProgram()
	checkGlError()
	gl.GlAttachShader(prog, vert_shader)
	checkGlError()
	gl.GlAttachShader(prog, frag_shader)
	checkGlError()

	gl.GlBindFragDataLocation(prog, 0, "outColor")
	checkGlError()

	gl.GlLinkProgram(prog)
	checkGlError()
	gl.GlUseProgram(prog)
	checkGlError()

	var pos_attr = gl.GlGetAttribLocation(prog, "position")
	checkGlError()
	gl.GlVertexAttribPointer(gl.GLuint(pos_attr), 2, gl.GL_FLOAT, false, 3, 0)
	checkGlError()
	gl.GlEnableVertexAttribArray(gl.GLuint(pos_attr))
	checkGlError()

	var color_attr = gl.GlGetAttribLocation(prog, "color")
	checkGlError()
	gl.GlVertexAttribPointer(gl.GLuint(color_attr), 3, gl.GL_FLOAT, false, 3, 2)
	checkGlError()
	gl.GlEnableVertexAttribArray(gl.GLuint(color_attr))
	checkGlError()

main_loop:
	for {

		// Event process
		var e sdl2.SDL_Event
		if 1 == sdl2.SDL_PollEvent(&e) {

			var event_type = e.Type()

			if sdl2.SDL_QUIT == event_type {
				break main_loop
			}

			if sdl2.SDL_KEYUP == event_type {
				var e1 = e.Key(nil)
				if e1.Keysym.Sym == sdl2.SDL_Keycode(sdl2.SDLK_ESCAPE) {
					break main_loop
				}
			}
		} // if

		gl.GlDrawArrays(gl.GL_TRIANGLES, 0, 3)
		checkGlError()

		sdl2.SDL_GL_SwapWindow(window)

	} // for

	sdl2.SDL_GL_DeleteContext(context)
	sdl2.SDL_DestroyWindow(window)
	sdl2.SDL_Quit()
}

func createShader(shader_type gl.GLenum, file string) gl.GLuint {

	var src string
	if a, err := os.ReadFile(file); nil == err {
		src = string(a)
	} else {
		fmt.Printf("Cannot open file '%v': %v\n", file, err)
	}

	var shader = gl.GlCreateShader(shader_type)
	gl.GlShaderSource1(shader, &src, nil)
	checkGlError()
	gl.GlCompileShader(shader)
	checkGlError()

	var status gl.GLint
	gl.GlGetShaderiv(shader, gl.GL_COMPILE_STATUS, &status)
	checkGlError()

	if gl.GL_TRUE != gl.GLboolean(status) {
		fmt.Println("Compiling shader failed")
	}

	// Retrieve log
	var log string
	gl.GlGetShaderInfoLog(shader, 512, nil, &log)
	if len(log) > 0 {
		fmt.Println(log)
	}

	return shader
}

func checkGlError() {
	if err := gl.GlGetError(); 0 != err {
		panic(fmt.Sprintf("GL error: %v", err))
	}
}
