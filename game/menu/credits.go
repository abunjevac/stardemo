package menu

import (
	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type Credits struct {
	bounds      gk.Rect
	fontCache   *gk.FontCache
	textFactory *gk.TextFactory
	scrollers   []*Scroller
}

func NewCredits(renderer *gk.Renderer, bounds gk.Rect, fontCache *gk.FontCache, textFactory *gk.TextFactory) *Credits {
	const scrollHeight = 32

	red := sdl.Color{R: 200, G: 70, B: 70, A: 255}
	green := sdl.Color{R: 70, G: 200, B: 70, A: 255}
	blue := sdl.Color{R: 70, G: 70, B: 200, A: 255}

	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	yellow := sdl.Color{R: 255, G: 255, B: 0, A: 255}

	font := fontCache.GetFont("zorque", 20)

	scrollerRect := func(index int) gk.Rect {
		return gk.Rect{
			X: bounds.X,
			Y: bounds.H - int32(index+1)*scrollHeight,
			W: bounds.W,
			H: scrollHeight,
		}
	}

	backBuilder := func(color sdl.Color) *gk.Texture {
		return renderer.CreateRectTexture(color, bounds.Size())
	}

	textBuilder := func(text string, foreColor sdl.Color) *gk.Texture {
		const sep = "              -=-             "

		return textFactory.NewShadedText(text+sep, font, foreColor)
	}

	texts := []string{
		"Welcome to Starship Demo",
		"Demo written by Bungee 2025",
		"Thanks for playing!",
	}

	scroller1 := NewScroller(renderer, scrollerRect(0), backBuilder(blue), textBuilder(texts[2], yellow), 10, 330)
	scroller2 := NewScroller(renderer, scrollerRect(1), backBuilder(green), textBuilder(texts[1], white), 14, 210)
	scroller3 := NewScroller(renderer, scrollerRect(2), backBuilder(red), textBuilder(texts[0], yellow), 18, 90)

	return &Credits{
		bounds:      bounds,
		fontCache:   fontCache,
		textFactory: textFactory,
		scrollers:   []*Scroller{scroller1, scroller2, scroller3},
	}
}

func (c *Credits) Update(delta uint64) {
	for _, s := range c.scrollers {
		s.Update(delta)
	}
}

func (c *Credits) Render() {
	for _, s := range c.scrollers {
		s.Render()
	}
}

func (c *Credits) Destroy() {
	for _, s := range c.scrollers {
		s.Destroy()
	}
}
