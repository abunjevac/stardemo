package play

import (
	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type Announcement struct {
	renderer    *gk.Renderer
	bounds      gk.Rect
	font        *gk.Font
	textFactory *gk.TextFactory
	text        string
	timer       float64
	duration    float64
	fadeIn      float64
	fadeOut     float64
	alpha       uint8
}

func NewAnnouncement(
	renderer *gk.Renderer,
	bounds gk.Rect,
	fontFactory *gk.FontCache,
	textFactory *gk.TextFactory,
) *Announcement {
	return &Announcement{
		renderer:    renderer,
		bounds:      bounds,
		font:        fontFactory.GetFont("zorque", 32).Italic(),
		textFactory: textFactory,
		duration:    2.0, // display for 2 seconds
		fadeIn:      0.5, // fade in over 0.5 seconds
		fadeOut:     0.5, // fade out over 0.5 seconds
		alpha:       0,   // start fully transparent
	}
}

func (a *Announcement) AnimateText(text string) {
	a.text = text
	a.timer = 0
	a.alpha = 0 // start with 0 alpha for fade-in
}

func (a *Announcement) Update(delta uint64) {
	if a.text == "" {
		return
	}

	deltaSeconds := float64(delta) / 1000.0

	a.timer += deltaSeconds

	// fade in
	if a.timer < a.fadeIn {
		progress := a.timer / a.fadeIn

		a.alpha = uint8(255 * progress)

		return
	}

	// fade out
	if a.timer >= a.duration {
		fadeProgress := (a.timer - a.duration) / a.fadeOut

		if fadeProgress >= 1.0 {
			a.text = ""
			a.alpha = 0
		} else {
			a.alpha = uint8(255 * (1.0 - fadeProgress))
		}
	} else {
		// fully visible during the main duration
		a.alpha = 255
	}
}

func (a *Announcement) Render() {
	if a.text == "" {
		return
	}

	a.RenderText(a.text, a.alpha)
}

func (a *Announcement) RenderText(text string, alpha uint8) {
	orangeColor := sdl.Color{R: 255, G: 165, B: 0, A: alpha}

	tex := a.textFactory.NewShadedText(text, a.font, orangeColor)
	defer tex.Destroy()

	w := tex.W
	h := tex.H
	x := (a.bounds.W - w) / 2
	y := (a.bounds.H - h) / 2

	a.renderer.Copy(tex, gk.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	})
}

func (a *Announcement) Destroy() {
}
