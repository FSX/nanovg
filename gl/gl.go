package gl

/*
#cgo linux pkg-config: gl glew
#cgo linux LDFLAGS: -lm

#include <GL/glew.h>
*/
import "C"
import "errors"

type bitField int

const (
	COLOR_BUFFER_BIT   bitField = C.GL_COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT            = C.GL_DEPTH_BUFFER_BIT
	ACCUM_BUFFER_BIT            = C.GL_ACCUM_BUFFER_BIT
	STENCIL_BUFFER_BIT          = C.GL_STENCIL_BUFFER_BIT
)

func InitGlew() error {
	C.glewExperimental = C.GL_TRUE
	err := C.glewInit()

	if err != C.GLEW_OK {
		// fmt.Println(C.GoString(C.glewGetErrorString(glew_err)))
		return errors.New("glew error")
	}

	// Why did I need to do this?
	C.glGetError()

	return nil
}

func Viewport(x, y, with, height int) {
	C.glViewport(
		(C.GLint)(x), (C.GLint)(y),
		(C.GLsizei)(with), (C.GLsizei)(height))
}

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(
		(C.GLclampf)(red), (C.GLclampf)(green),
		(C.GLclampf)(blue), (C.GLclampf)(alpha))
}

func Clear(mask bitField) {
	C.glClear((C.GLbitfield)(mask))
}
