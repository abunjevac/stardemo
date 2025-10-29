package gk

import (
	"fmt"
)

type Size struct {
	W int32
	H int32
}

func NewSize(w int32, h int32) Size {
	return Size{W: w, H: h}
}

func (s Size) Decompose() (w, h int32) {
	return s.W, s.H
}

func (s Size) Extend(v int32) Size {
	return NewSize(s.W+v, s.H+v)
}

func (s Size) Rect() Rect {
	return Rect{W: s.W, H: s.H}
}

func (s Size) String() string {
	return fmt.Sprintf("%d:%d", s.W, s.H)
}
