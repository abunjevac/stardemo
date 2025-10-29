package play

import (
	"iter"
	"maps"

	"stardemo/gk"
)

type Enemies struct {
	renderer *gk.Renderer
	enemies  map[int]*Enemy
}

func NewEnemies(renderer *gk.Renderer) *Enemies {
	return &Enemies{
		renderer: renderer,
		enemies:  make(map[int]*Enemy),
	}
}

func (e *Enemies) Add(enemy *Enemy) {
	e.enemies[enemy.id] = enemy
}

func (e *Enemies) Remove(id int) {
	delete(e.enemies, id)
}

func (e *Enemies) Update(delta uint64) {
	for _, enemy := range e.enemies {
		enemy.Update(delta)

		if enemy.IsDestroyed() {
			delete(e.enemies, enemy.id)

			return
		}
	}
}

func (e *Enemies) Render(bounds gk.Rect) {
	for _, enemy := range e.enemies {
		enemy.Render(e.renderer, bounds)
	}
}

func (e *Enemies) Destroy() {
	clear(e.enemies)
}

func (e *Enemies) All() iter.Seq[*Enemy] {
	return maps.Values(e.enemies)
}

func (e *Enemies) WipedOut() bool {
	return len(e.enemies) == 0
}
