package gl

// #cgo CFLAGS: -DGLEW_STATIC
// #cgo LDFLAGS: -lglew32 -lopengl32
//
// #include <GL/glew.h>
// #include <GL/gl.h>
//
import "C"

func init() {
	// C.glewExperimental = C.GL_TRUE
	// C.glewInit()
}
