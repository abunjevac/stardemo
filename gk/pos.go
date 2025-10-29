package gk

import "fmt"

var NullPos = Pos{}

type Pos struct {
	X int32
	Y int32
}

func NewPos(x, y int32) Pos {
	return Pos{X: x, Y: y}
}

func (p Pos) Adjust(dx, dy int32) Pos {
	return Pos{X: p.X + dx, Y: p.Y + dy}
}

func (p Pos) AdjustX(dx int32) Pos {
	return Pos{X: p.X + dx, Y: p.Y}
}

func (p Pos) AdjustY(dy int32) Pos {
	return Pos{X: p.X, Y: p.Y + dy}
}

func (p Pos) AdjustPos(p2 Pos) Pos {
	return Pos{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Pos) Rect(w, h int32) Rect {
	return Rect{X: p.X, Y: p.Y, W: w, H: h}
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

type FPos struct {
	X float32
	Y float32
}

func NewFPos(x, y float32) FPos {
	return FPos{X: x, Y: y}
}
