package play

import (
	"stardemo/gk"
)

type Bullet struct {
	Width     int32
	Height    int32
	Faction   Faction
	id        int
	image     *gk.Texture
	pos       gk.Pos
	direction Direction
	speed     int32
	destroyed bool
}

func NewBullet(id int, image *gk.Texture, faction Faction, pos gk.Pos, direction Direction, speed int32) *Bullet {
	return &Bullet{
		Width:     40,
		Height:    12,
		id:        id,
		image:     image,
		Faction:   faction,
		pos:       pos,
		direction: direction,
		speed:     speed,
	}
}

func (b *Bullet) Advance(delta uint64, bounds gk.Rect) {
	b.pos = b.pos.AdjustX(int32(delta) * int32(b.direction) * b.speed / 10)

	if b.pos.X < 0 || b.pos.X > bounds.W {
		b.destroyed = true
	}
}

func (b *Bullet) Render(r *gk.Renderer, bounds gk.Rect) {
	rect := b.pos.
		AdjustPos(bounds.Pos()).
		AdjustY(-b.Height/2).
		Rect(b.Width, b.Height)

	r.Copy(b.image, rect)
}

func (b *Bullet) IsDestroyed() bool {
	return b.destroyed
}

func (b *Bullet) Destroy() {
	b.destroyed = true
}
