package play

import (
	"iter"
	"maps"

	"stardemo/gk"
)

type Bullets struct {
	renderer *gk.Renderer
	bullets  map[int]*Bullet
}

func NewBullets(renderer *gk.Renderer) *Bullets {
	return &Bullets{
		renderer: renderer,
		bullets:  make(map[int]*Bullet),
	}
}

func (b *Bullets) Add(bullet *Bullet) {
	b.bullets[bullet.id] = bullet
}

func (b *Bullets) Remove(id int) {
	delete(b.bullets, id)
}

func (b *Bullets) Update(delta uint64, bounds gk.Rect) {
	for _, bullet := range b.bullets {
		bullet.Advance(delta, bounds)

		if bullet.IsDestroyed() {
			delete(b.bullets, bullet.id)
		}
	}
}

func (b *Bullets) Render(bounds gk.Rect) {
	for _, bullet := range b.bullets {
		bullet.Render(b.renderer, bounds)
	}
}

func (b *Bullets) All() iter.Seq[*Bullet] {
	return maps.Values(b.bullets)
}

func (b *Bullets) Destroy() {
	clear(b.bullets)
}
