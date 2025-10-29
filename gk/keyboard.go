package gk

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	Ctrl  = 1
	Shift = 2
)

type Keyboard struct {
	prevState    map[sdl.Keycode]uint8
	currentState map[sdl.Keycode]uint8
	currentMod   int8
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		prevState:    make(map[sdl.Keycode]uint8),
		currentState: make(map[sdl.Keycode]uint8),
	}
}

func (k *Keyboard) Update(e *sdl.KeyboardEvent) {
	k.prevState = make(map[sdl.Keycode]uint8)

	for key, state := range k.currentState {
		k.prevState[key] = state
	}

	k.currentState[e.Keysym.Sym] = e.State

	k.currentMod = 0

	if e.Keysym.Mod&sdl.KMOD_CTRL != 0 {
		k.currentMod |= Ctrl
	}

	if e.Keysym.Mod&sdl.KMOD_SHIFT != 0 {
		k.currentMod |= Shift
	}
}

func (k *Keyboard) IsPressed(key sdl.Keycode) bool {
	return k.isKeyDown(key) && k.currentMod == 0
}

func (k *Keyboard) IsPressedMod(key sdl.Keycode, mod int8) bool {
	return k.isKeyDown(key) && k.currentMod == mod
}

func (k *Keyboard) IsReleased(key sdl.Keycode) bool {
	return k.isKeyReleased(key) && k.currentMod == 0
}

func (k *Keyboard) IsReleasedMod(key sdl.Keycode, mod int8) bool {
	return k.isKeyReleased(key) && k.currentMod == mod
}

func (k *Keyboard) isKeyDown(key sdl.Keycode) bool {
	if _, ok := k.currentState[key]; !ok {
		return false
	}

	return k.prevState[key] == sdl.RELEASED && k.currentState[key] == sdl.PRESSED
}

// PressedIndex returns index of pressed key.
// Modifiers are not checked.
func (k *Keyboard) PressedIndex(keys ...sdl.Keycode) (int, bool) {
	for i, key := range keys {
		if k.isKeyDown(key) {
			return i, true
		}
	}

	return -1, false
}

func (k *Keyboard) isKeyReleased(key sdl.Keycode) bool {
	if _, ok := k.currentState[key]; !ok {
		return false
	}

	return k.prevState[key] == sdl.PRESSED && k.currentState[key] == sdl.RELEASED
}
