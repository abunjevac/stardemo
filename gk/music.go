package gk

import (
	"path/filepath"

	"github.com/veandco/go-sdl2/mix"
)

type Music struct {
	peer *mix.Music
}

func NewMusic(path string) *Music {
	mus, err := mix.LoadMUS(filepath.Join(musicDir, path))

	panicErr(err)

	return &Music{mus}
}

func (m *Music) Play(channel int) {
	panicErr(m.peer.Play(channel))
}

func (m *Music) Destroy() {
	m.peer.Free()
}
