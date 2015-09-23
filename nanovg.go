package nanovg

/*
#cgo linux pkg-config: gl glew
#cgo linux LDFLAGS: -lm

#include <GL/glew.h>

#include "nanovg.h"
#define NANOVG_GL3_IMPLEMENTATION
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

type Color struct {
	v C.NVGcolor
}

type Context struct {
	ctx *C.NVGcontext
}

func CreateCtx(flags createFlags) (*Context, error) {
	ctx := C.nvgCreateGL3((C.int)(flags))
	if ctx == nil {
		return nil, errors.New("cannot create nanovg context")
	}

	return &Context{ctx}, nil
}

func (c *Context) Close() {
	C.nvgDeleteGL3(c.ctx)
	c.ctx = nil
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

// Color utils //

// TODO: NVGcolor nvgRGB(unsigned char r, unsigned char g, unsigned char b);
// TODO: NVGcolor nvgRGBf(float r, float g, float b);

func RGBA(r, g, b, a uint8) Color {
	return Color{C.nvgRGBA(
		(C.uchar)(r), (C.uchar)(g),
		(C.uchar)(b), (C.uchar)(a))}
}

// TODO: NVGcolor nvgRGBAf(float r, float g, float b, float a);
// TODO: NVGcolor nvgLerpRGBA(NVGcolor c0, NVGcolor c1, float u);
// TODO: NVGcolor nvgTransRGBA(NVGcolor c0, unsigned char a);
// TODO: NVGcolor nvgTransRGBAf(NVGcolor c0, float a);
// TODO: NVGcolor nvgHSL(float h, float s, float l);
// TODO: NVGcolor nvgHSLA(float h, float s, float l, unsigned char a);

// Render styles //

func (c *Context) StrokeColor(color Color) {
	C.nvgStrokeColor(c.ctx, color.v)
}

// TODO: void nvgStrokePaint(NVGcontext* ctx, NVGpaint paint);

func (c *Context) FillColor(color Color) {
	C.nvgFillColor(c.ctx, color.v)
}

// TODO: void nvgFillPaint(NVGcontext* ctx, NVGpaint paint);
// TODO: void nvgMiterLimit(NVGcontext* ctx, float limit);

func (c *Context) StrokeWidth(size float32) {
	C.nvgStrokeWidth(c.ctx, (C.float)(size))
}

// TODO: void nvgLineCap(NVGcontext* ctx, int cap);
// TODO: void nvgLineJoin(NVGcontext* ctx, int join);
// TODO: void nvgGlobalAlpha(NVGcontext* ctx, float alpha);

// Paths //

func (c *Context) BeginPath() {
	C.nvgBeginPath(c.ctx)
}

func (c *Context) MoveTo(x, y float32) {
	C.nvgMoveTo(c.ctx, (C.float)(x), (C.float)(y))
}

func (c *Context) LineTo(x, y float32) {
	C.nvgLineTo(c.ctx, (C.float)(x), (C.float)(y))
}

// TODO: void nvgBezierTo(NVGcontext* ctx, float c1x, float c1y, float c2x, float c2y, float x, float y);
// TODO: void nvgQuadTo(NVGcontext* ctx, float cx, float cy, float x, float y);
// TODO: void nvgArcTo(NVGcontext* ctx, float x1, float y1, float x2, float y2, float radius);

func (c *Context) ClosePath() {
	C.nvgClosePath(c.ctx)
}

// TODO: void nvgPathWinding(NVGcontext* ctx, int dir);
// TODO: void nvgArc(NVGcontext* ctx, float cx, float cy, float r, float a0, float a1, int dir);

func (c *Context) Rect(x, y, w, h float32) {
	C.nvgRect(c.ctx,
		(C.float)(x), (C.float)(y),
		(C.float)(w), (C.float)(h))
}

// TODO: void nvgRoundedRect(NVGcontext* ctx, float x, float y, float w, float h, float r);
// TODO: void nvgEllipse(NVGcontext* ctx, float cx, float cy, float rx, float ry);
// TODO: void nvgCircle(NVGcontext* ctx, float cx, float cy, float r);

func (c *Context) Fill() {
	C.nvgFill(c.ctx)
}

func (c *Context) Stroke() {
	C.nvgStroke(c.ctx)
}
