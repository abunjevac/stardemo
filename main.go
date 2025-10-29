package main

import (
	"stardemo/game"
	"stardemo/gk"
	"stardemo/ui"
)

func main() {
	gk.Initialize()
	defer gk.Destroy()

	window := gk.NewWindow(gk.WindowConfig{
		Title: "Starship",
		// Pos:   gk.NewPos(2810, 50),
		// Size:  gk.NewSize(1920, 1080),
	})
	defer window.Destroy()

	window.Scale(100)

	bounds := window.Rect()
	renderer := window.CreateRenderer()
	keyboard := gk.NewKeyboard()
	fontCache := gk.NewFontCache()
	imageCache := gk.NewImageCache(renderer)
	textFactory := gk.NewTextFactory(renderer)
	defaultFont := fontCache.GetFont("zorque", 28)
	effects := gk.NewEffects(imageCache)

	effects.Load("impact", 0, 7)
	effects.Load("box", 0, 7)

	g := game.New(&ui.Context{
		Bounds:      bounds,
		Renderer:    renderer,
		Keyboard:    keyboard,
		FontCache:   fontCache,
		ImageCache:  imageCache,
		TextFactory: textFactory,
		DefaultFont: defaultFont,
		Effects:     effects,
	})

	g.Run()

	imageCache.Destroy()
	fontCache.Destroy()
	renderer.Destroy()
}
