package gk

import "github.com/veandco/go-sdl2/sdl"

type TextFactory struct {
	renderer *Renderer
}

func NewTextFactory(renderer *Renderer) *TextFactory {
	return &TextFactory{renderer: renderer}
}

func (f *TextFactory) NewShadedText(text string, font *Font, color sdl.Color) *Texture {
	surf := newShadedTextSurface(text, font, color)
	defer surf.Destroy()

	return f.renderer.CreateTextureFromSurface(surf)
}

func (f *TextFactory) NewText(text string, font *Font, color sdl.Color) *Texture {
	surf := font.CreateBlendedText(text, color)
	defer surf.Destroy()

	return f.renderer.CreateTextureFromSurface(surf)
}

func newShadedTextSurface(text string, font *Font, color sdl.Color) *Surface {
	const shadeSize = 2
	const shadeColor = 16

	txt := font.CreateBlendedText(text, color)
	defer txt.Destroy()

	shade := font.CreateBlendedText(text, sdl.Color{R: shadeColor, G: shadeColor, B: shadeColor})
	defer shade.Destroy()

	surf := NewSurface(txt.W+shadeSize, txt.H+shadeSize)

	shade.Blit(surf, Rect{X: shadeSize, Y: shadeSize, W: txt.W, H: txt.H})
	txt.Blit(surf, Rect{W: txt.W, H: txt.H})

	return surf
}
