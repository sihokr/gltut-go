package main

import (
	"fmt"
	"runtime"

	"example.com/gltut_go/gl"
	"example.com/gltut_go/sdl2"
)

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

	// This has to be called AFTER OpenGL context creation!!
	gl.GlewInit()

	var buffers = make([]gl.GLuint, 1)
	gl.GlGenBuffers(1, buffers)

	fmt.Println(buffers[0])

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

		sdl2.SDL_GL_SwapWindow(window)

	} // for

	sdl2.SDL_GL_DeleteContext(context)
	sdl2.SDL_DestroyWindow(window)
	sdl2.SDL_Quit()
}
