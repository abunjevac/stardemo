package play

import (
	"stardemo/gk"
)

type Crate struct {
	Size      int32
	Guns      []*Gun
	id        int
	effect    *gk.ImageCarousel
	pos       gk.Pos
	speed     int32
	collected bool
	destroyed bool
}

func NewCrate(id int, effect *gk.ImageCarousel, pos gk.Pos, speed int32, guns []*Gun) *Crate {
	return &Crate{
		Size:   30,
		Guns:   guns,
		id:     id,
		effect: effect,
		pos:    pos,
		speed:  speed,
	}
}

func (c *Crate) Advance(delta uint64) {
	c.effect.Update(delta)

	c.pos = c.pos.AdjustX(-int32(delta) * c.speed / 100)

	if c.pos.X < 0 {
		c.destroyed = true
	}
}

func (c *Crate) Render(r *gk.Renderer, bounds gk.Rect) {
	rect := c.pos.
		AdjustPos(bounds.Pos()).
		AdjustY(-c.Size/2).
		Rect(c.Size, c.Size)

	r.Copy(c.effect.Current(), rect)
}

func (c *Crate) IsDestroyed() bool {
	return c.destroyed
}

func (c *Crate) Destroy() {
	c.destroyed = true
}
