package gk

import (
	"path/filepath"

	"github.com/veandco/go-sdl2/mix"
)

type Sound struct {
	peer *mix.Chunk
}

func NewSound(path string) *Sound {
	mus, err := mix.LoadWAV(filepath.Join(soundsDir, path))

	panicErr(err)

	return &Sound{mus}
}

func (m *Sound) Play(channel int) {
	_, err := m.peer.Play(channel, 0)

	panicErr(err)
}

func (m *Sound) Destroy() {
	m.peer.Free()
}
