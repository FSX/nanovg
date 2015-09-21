package nanovg

/*
#cgo linux pkg-config: gl glew
#cgo linux LDFLAGS: -lm

#include <GL/glew.h>

#include "nanovg.h"
#define NANOVG_GL2_IMPLEMENTATION
#include "nanovg_gl.h"
#include "nanovg_gl_utils.h"
*/
import "C"
import "errors"

type createFlags int

const (
	ANTIALIAS       createFlags = C.NVG_ANTIALIAS
	STENCIL_STROKES             = C.NVG_STENCIL_STROKES
	DEBUG                       = C.NVG_DEBUG
)

type Context struct {
	ctx *C.NVGcontext
}

func CreateCtx(flags createFlags) (*Context, error) {
	ctx := C.nvgCreateGL2((C.int)(flags))
	if ctx == nil {
		return nil, errors.New("cannot create nanovg context")
	}

	return &Context{ctx}, nil
}

func (c *Context) BeginFrame(windowWidth, windowHeight int, devicePixelRatio float32) {
	C.nvgBeginFrame(c.ctx,
		(C.int)(windowWidth), (C.int)(windowHeight),
		(C.float)(devicePixelRatio))
}

func (c *Context) CancelFrame() {
	C.nvgCancelFrame(c.ctx)
}

func (c *Context) EndFrame() {
	C.nvgEndFrame(c.ctx)
}

func (c *Context) Close() {
	C.nvgDeleteGL2(c.ctx)
	c.ctx = nil
}
