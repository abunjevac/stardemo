package gk

import "strconv"

type Effects struct {
	imageCache *ImageCache
	cache      map[string][]*Texture
}

func NewEffects(imageCache *ImageCache) *Effects {
	return &Effects{
		imageCache: imageCache,
		cache:      make(map[string][]*Texture),
	}
}

func (e *Effects) Load(name string, from, to int) {
	for i := from; i <= to; i++ {
		e.cache[name] = append(e.cache[name], e.imageCache.Get(name+strconv.Itoa(i)))
	}
}

func (e *Effects) Get(name string) []*Texture {
	return e.cache[name]
}
