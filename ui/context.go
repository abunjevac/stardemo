package ui

import (
	"stardemo/gk"
)

type Context struct {
	Bounds      gk.Rect
	Renderer    *gk.Renderer
	Keyboard    *gk.Keyboard
	FontCache   *gk.FontCache
	ImageCache  *gk.ImageCache
	TextFactory *gk.TextFactory
	DefaultFont *gk.Font
	Effects     *gk.Effects
}
