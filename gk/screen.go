package gk

import "github.com/veandco/go-sdl2/sdl"

func GetScreenSize() (width, height int32) {
	mode, err := sdl.GetDesktopDisplayMode(0)

	panicErr(err)

	width = mode.W
	height = mode.H

	return width, height
}
