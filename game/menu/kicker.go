package menu

import (
	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type Kicker struct {
	renderer    *gk.Renderer
	bounds      gk.Rect
	textTexture *gk.Texture
	visible     bool
	timer       float64
	flashRate   float64
}

func NewKicker(renderer *gk.Renderer, bounds gk.Rect, textFactory *gk.TextFactory, font *gk.Font) *Kicker {
	const message = "Press SPACE to save the Universe"

	textColor := sdl.Color{R: 0, G: 255, B: 255, A: 255}
	textTexture := textFactory.NewText(message, font, textColor)

	return &Kicker{
		renderer:    renderer,
		bounds:      bounds,
		textTexture: textTexture,
		visible:     true,
		timer:       0,
		flashRate:   0.5,
	}
}

func (k *Kicker) Update(delta uint64) {
	k.timer += float64(delta) / 1000.0

	if k.timer >= k.flashRate {
		k.visible = !k.visible
		k.timer = 0
	}
}

func (k *Kicker) Render() {
	if !k.visible {
		return
	}

	centerX := k.bounds.X + k.bounds.W/2
	centerY := k.bounds.Y + k.bounds.H/2

	k.renderer.Copy(k.textTexture, gk.Rect{
		X: centerX - k.textTexture.W/2,
		Y: centerY - k.textTexture.H/2,
		W: k.textTexture.W,
		H: k.textTexture.H,
	})
}

func (k *Kicker) Destroy() {
	k.textTexture.Destroy()
}
