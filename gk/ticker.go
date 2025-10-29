package gk

import "github.com/veandco/go-sdl2/sdl"

type Ticker struct {
	FrameStep uint64
	markTime  uint64
}

func NewTicker(targetFPS uint64) *Ticker {
	return &Ticker{
		FrameStep: 1000 / targetFPS,
	}
}

func (t *Ticker) Mark() {
	t.markTime = sdl.GetTicks64()
}

func (t *Ticker) Yield() {
	timeFromMark := sdl.GetTicks64() - t.markTime

	if timeFromMark < t.FrameStep {
		sdl.Delay(uint32(t.FrameStep - timeFromMark))
	}
}
