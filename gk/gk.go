package gk

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// standard directories
const (
	imagesDir = "ui/assets/images/"
	fontsDir  = "ui/assets/fonts/"
	soundsDir = "ui/assets/sounds/"
	musicDir  = "ui/assets/music/"
)

func Initialize() {
	v := sdl.Version{}

	sdl.GetVersion(&v)

	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_INFO)
	sdl.Log("SDL Version: %d.%d.%d", v.Major, v.Minor, v.Patch)

	panicErr(sdl.Init(sdl.INIT_EVERYTHING))
	panicErr(ttf.Init())

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
}

func Destroy() {
	sdl.Quit()
}
