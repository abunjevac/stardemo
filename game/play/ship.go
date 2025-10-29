package play

import (
	"stardemo/gk"
)

const (
	ShipAerie   = "aerie"
	ShipClipper = "clipper"
	ShipFury    = "mfury"
)

type Ship struct {
	Size          int32
	image         *gk.Texture
	trail         *gk.Texture
	pos           gk.Pos
	faction       Faction
	trailOffset   int32
	trailAnim     int32
	hit           bool
	destroyed     bool
	destroyEffect *gk.ImageCarousel
}

func NewShip(image *gk.Texture, trail *gk.Texture, dir Direction, faction Faction, pos gk.Pos) *Ship {
	return &Ship{
		Size:        64,
		image:       image,
		trail:       trail,
		faction:     faction,
		pos:         pos,
		trailOffset: int32(60 * dir),
	}
}

func (s *Ship) Advance(delta uint64) {
	if s.hit {
		s.destroyEffect.Update(delta)

		if s.destroyEffect.Done() {
			s.destroyed = true
		}
	}

	s.trailAnim++

	if s.trailAnim > 10 {
		s.trailAnim = 0
	}
}

func (s *Ship) RenderShip(r *gk.Renderer, pos gk.Pos) {
	var rx int32

	if s.hit {
		rx = gk.RandN2(-3, 3)
	}

	r.Copy(s.trail, gk.Rect{
		X: pos.X + rx - s.trailOffset,
		Y: pos.Y + rx,
		W: s.Size,
		H: s.Size + s.trailAnim/4,
	})

	r.Copy(s.image, gk.Rect{
		X: pos.X + rx,
		Y: pos.Y + rx,
		W: s.Size,
		H: s.Size,
	})

	if s.hit {
		rect := gk.Rect{
			X: pos.X - s.Size/2,
			Y: pos.Y - s.Size/2,
			W: s.Size * 2,
			H: s.Size * 2,
		}

		r.Copy(s.destroyEffect.Previous(), rect)
		r.Copy(s.destroyEffect.Current(), rect)
	}
}

func (s *Ship) Wrecked(hitImages []*gk.Texture) {
	s.hit = true
	s.destroyEffect = gk.NewImageCarousel(hitImages, 2000, false)
}

func (s *Ship) IsHit() bool {
	return s.hit
}

func (s *Ship) IsDestroyed() bool {
	return s.destroyed
}

func (s *Ship) CollidesWith(other *Ship) bool {
	x1, y1 := s.pos.X+s.Size/2, s.pos.Y+s.Size/2
	x2, y2 := other.pos.X+s.Size/2, other.pos.Y+s.Size/2

	r1, r2 := s.Size/2, other.Size/2

	dist := gk.Dist(x1, y1, x2, y2)

	return dist < float64(r1+r2)*0.8
}

func (s *Ship) TakesDamageFrom(bullet *Bullet) bool {
	// friendly fire?
	if bullet.Faction == s.faction {
		return false
	}

	x1, y1 := s.pos.X+s.Size/2, s.pos.Y+s.Size/2
	x2, y2 := bullet.pos.X+bullet.Width/2, bullet.pos.Y+bullet.Height/2

	r1, r2 := s.Size/2, bullet.Width/2

	dist := gk.Dist(x1, y1, x2, y2)

	return dist < float64(r1+r2)*0.6
}

func (s *Ship) CanCollect(crate *Crate) bool {
	x1, y1 := s.pos.X+s.Size/2, s.pos.Y+s.Size/2
	x2, y2 := crate.pos.X+crate.Size/2, crate.pos.Y+crate.Size/2

	r1, r2 := s.Size/2, crate.Size/2

	dist := gk.Dist(x1, y1, x2, y2)

	return dist < float64(r1+r2)*0.8
}
