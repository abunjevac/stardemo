package play

import (
	"iter"
	"maps"

	"stardemo/gk"
)

type Crates struct {
	renderer *gk.Renderer
	crates   map[int]*Crate
}

func NewCrates(renderer *gk.Renderer) *Crates {
	return &Crates{
		renderer: renderer,
		crates:   make(map[int]*Crate),
	}
}

func (c *Crates) Add(crate *Crate) {
	c.crates[crate.id] = crate
}

func (c *Crates) Remove(id int) {
	delete(c.crates, id)
}

func (c *Crates) Update(delta uint64) {
	for _, crate := range c.crates {
		crate.Advance(delta)

		if crate.IsDestroyed() {
			delete(c.crates, crate.id)
		}
	}
}

func (c *Crates) Render(bounds gk.Rect) {
	for _, crate := range c.crates {
		crate.Render(c.renderer, bounds)
	}
}

func (c *Crates) Destroy() {
	clear(c.crates)
}

func (c *Crates) All() iter.Seq[*Crate] {
	return maps.Values(c.crates)
}

func (c *Crates) AllCollected() bool {
	return len(c.crates) == 0
}
