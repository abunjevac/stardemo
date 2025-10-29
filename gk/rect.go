package gk

import "github.com/veandco/go-sdl2/sdl"

var NullRect = Rect{}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func NewRect(x, y, w, h int32) Rect {
	return Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (r Rect) peer() *sdl.Rect {
	return &sdl.Rect{X: r.X, Y: r.Y, W: r.W, H: r.H}
}

func (r Rect) Pos() Pos {
	return Pos{X: r.X, Y: r.Y}
}

func (r Rect) Size() Size {
	return NewSize(r.W, r.H)
}

func (r Rect) Divide(s Size) Size {
	return NewSize(r.W/s.W, r.H/s.H)
}

func (r Rect) ApplyRatio(w, h int32) Rect {
	return Rect{
		X: r.X,
		Y: r.Y,
		W: r.W * w / 100,
		H: r.H * h / 100,
	}
}

func (r Rect) EnsureRatio(s Size) Rect {
	return Rect{
		X: r.X,
		Y: r.Y,
		W: (r.W / s.W) * s.W,
		H: (r.H / s.H) * s.H,
	}
}

func (r Rect) Implode(d int32) Rect {
	return Rect{
		X: r.X + d/2,
		Y: r.Y + d/2,
		W: r.W - d,
		H: r.H - d,
	}
}

func (r Rect) MoveOrigin(dx, dy int32) Rect {
	return Rect{
		X: r.X + dx,
		Y: r.Y + dy,
		W: r.W,
		H: r.H,
	}
}

func (r Rect) Resize(size Size) Rect {
	return Rect{
		X: r.X,
		Y: r.Y,
		W: size.W,
		H: size.H,
	}
}

func (r Rect) ApplyDeltaY(dy int32) Rect {
	return Rect{
		X: r.X,
		Y: r.Y + dy,
		W: r.W,
		H: r.H,
	}
}

type FRect struct {
	X float32
	Y float32
	W float32
	H float32
}

func NewFRect(x, y, w, h float32) FRect {
	return FRect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (r FRect) peer() *sdl.FRect {
	return &sdl.FRect{X: r.X, Y: r.Y, W: r.W, H: r.H}
}
