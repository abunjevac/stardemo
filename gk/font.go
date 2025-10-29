package gk

import (
	"fmt"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	font    *ttf.Font
	file    string
	factory *FontCache
	Height  int32
}

func (f *Font) Bold() *Font {
	return f.factory.lookupOrCreateFont(f.file, int32(f.font.Height()), f.font.GetStyle()|ttf.STYLE_BOLD)
}

func (f *Font) Italic() *Font {
	return f.factory.lookupOrCreateFont(f.file, int32(f.font.Height()), f.font.GetStyle()|ttf.STYLE_ITALIC)
}

func (f *Font) Size(size int32) *Font {
	return f.factory.lookupOrCreateFont(f.file, size, f.font.GetStyle())
}

func (f *Font) TextSize(text string) Size {
	w, h, err := f.font.SizeUTF8(text)

	panicErr(err)

	return NewSize(int32(w), int32(h))
}

func (f *Font) CreateBlendedText(text string, color sdl.Color) *Surface {
	if text == "" {
		return NewSurface(1, 1)
	}

	surface, err := f.font.RenderUTF8Blended(text, color)

	panicErr(err)

	return newSurface(surface)
}

func (f *Font) CreateWrappedBlendedText(text string, color sdl.Color, wrapLength int32) *Surface {
	surface, err := f.font.RenderUTF8BlendedWrapped(text, color, int(wrapLength))

	panicErr(err)

	return newSurface(surface)
}

type FontCache struct {
	fonts map[string]*Font
}

func NewFontCache() *FontCache {
	return &FontCache{fonts: make(map[string]*Font)}
}

func (c *FontCache) GetFont(file string, size int32) *Font {
	return c.lookupOrCreateFont(file, size, ttf.STYLE_NORMAL)
}

func (c *FontCache) Destroy() {
	for _, f := range c.fonts {
		f.font.Close()
	}

	clear(c.fonts)
}

func (c *FontCache) lookupOrCreateFont(file string, size int32, style int) *Font {
	key := c.fontKey(file, size, style)

	if f, ok := c.fonts[key]; ok {
		return f
	}

	font := c.loadFont(file, size, style)

	_, surfaceHeight, err := font.SizeUTF8("Aq")

	panicErr(err)

	f := &Font{
		font:    font,
		file:    file,
		factory: c,
		Height:  int32(surfaceHeight),
	}

	c.fonts[key] = f

	return f
}

func (c *FontCache) fontKey(file string, size int32, style int) string {
	return fmt.Sprintf("%s-%d-%d", file, size, style)
}

func (c *FontCache) loadFont(file string, size int32, style int) *ttf.Font {
	font, err := ttf.OpenFont(filepath.Join(fontsDir, file+".ttf"), int(size))

	panicErr(err)

	font.SetStyle(style)

	return font
}
