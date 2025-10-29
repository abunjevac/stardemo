package gk

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	peer *sdl.Texture
	Size
}

func newTexture(peer *sdl.Texture) *Texture {
	_, _, width, height, err := peer.Query()

	panicErr(err)

	return &Texture{
		peer: peer,
		Size: NewSize(width, height),
	}
}

func (t *Texture) Destroy() {
	panicErr(t.peer.Destroy())
}

func (t *Texture) SetColorMod() {
	panicErr(t.peer.SetColorMod(128, 0, 0))
}

func (t *Texture) SetColorModDim() {
	panicErr(t.peer.SetColorMod(128, 128, 128))
}

func (t *Texture) UnsetColorMod() {
	panicErr(t.peer.SetColorMod(255, 255, 255))
}

func (t *Texture) SetAlphaMod(alpha uint8) {
	panicErr(t.peer.SetAlphaMod(alpha))
}
