package gk

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

type WindowConfig struct {
	Title string
	Pos   Pos
	Size  Size
}

type Window struct {
	peer *sdl.Window
}

func NewWindow(cfg WindowConfig) *Window {
	peer, err := sdl.CreateWindow(cfg.Title, cfg.Pos.X, cfg.Pos.Y, cfg.Size.W, cfg.Size.H, sdl.WINDOW_SHOWN)

	panicErr(err)

	return &Window{peer: peer}
}

func (w *Window) Scale(pct int) {
	winWidth, winHeight := GetScreenSize()

	if pct == 100 {
		// WINDOW_FULLSCREEN_DESKTOP is not working on MacOS
		if runtime.GOOS == "darwin" {
			panicErr(w.peer.SetFullscreen(sdl.WINDOW_FULLSCREEN))
			w.peer.SetSize(winWidth, winHeight)

			return
		}

		// WINDOW_FULLSCREEN would give us a 640x480 resolution
		panicErr(w.peer.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP))

		return
	}

	panicErr(w.peer.SetFullscreen(0))

	scale := float32(pct) / 100
	winWidth = Scale(winWidth, scale)
	winHeight = Scale(winHeight, scale)

	w.peer.SetSize(winWidth, winHeight)
}

func (w *Window) Size() Size {
	return NewSize(w.peer.GetSize())
}

func (w *Window) Rect() Rect {
	return w.Size().Rect()
}

func (w *Window) Destroy() {
	panicErr(w.peer.Destroy())
}

func (w *Window) CreateRenderer() *Renderer {
	r, err := sdl.CreateRenderer(w.peer, -1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)

	panicErr(err)

	return newRenderer(r)
}
