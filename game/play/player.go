package play

import (
	"iter"
	"slices"

	"stardemo/gk"
)

type Player struct {
	*Ship

	renderer *gk.Renderer
	stats    *Stats

	guns []*Gun

	verticalThrust   *Thrust
	horizontalThrust *Thrust

	upPressed    bool
	downPressed  bool
	leftPressed  bool
	rightPressed bool
}

func NewPlayer(renderer *gk.Renderer, bounds gk.Rect, images *gk.ImageCache, stats *Stats) *Player {
	const size = 64

	return &Player{
		Ship: NewShip(
			images.GetRotated("starship", 90),
			images.GetRotated("trail", 90),
			DirectionRightWay,
			FactionUs,
			gk.NewPos(-60, bounds.H/2-size/2),
		),

		renderer: renderer,
		stats:    stats,

		verticalThrust:   NewThrust(1300, 800, 400),
		horizontalThrust: NewThrust(1300, 800, 400),
	}
}

func (p *Player) Update(delta uint64, bounds gk.Rect) {
	// slow scene entry
	if p.pos.X < 60 {
		p.pos.X += 2
	}

	// advance ship
	p.Advance(delta)

	newY := p.pos.Y + p.verticalThrust.Advance(delta, !p.hit && p.upPressed, !p.hit && p.downPressed)
	newX := p.pos.X + p.horizontalThrust.Advance(delta, !p.hit && p.leftPressed, !p.hit && p.rightPressed)

	// keep player within screen bounds
	if newY < 0 {
		newY = 0

		p.verticalThrust.FullStop()
	} else if newY > bounds.H-p.Size {
		newY = bounds.H - p.Size

		p.verticalThrust.FullStop()
	}

	if newX < 0 {
		newX = 0

		p.horizontalThrust.FullStop()
	} else if newX > bounds.W-p.Size {
		newX = bounds.W - p.Size

		p.horizontalThrust.FullStop()
	}

	p.pos.Y = newY
	p.pos.X = newX

	p.stats.VerticalThrust = p.verticalThrust.Percentage()
	p.stats.HorizontalThrust = p.horizontalThrust.Percentage()
}

func (p *Player) Render(bounds gk.Rect) {
	p.RenderShip(p.renderer, p.pos.AdjustPos(bounds.Pos()))
}

func (p *Player) Destroy() {
}

func (p *Player) steerUpPressed() {
	p.upPressed = true
}

func (p *Player) steerUpReleased() {
	p.upPressed = false
}

func (p *Player) steerDownPressed() {
	p.downPressed = true
}

func (p *Player) steerDownReleased() {
	p.downPressed = false
}

func (p *Player) steerLeftPressed() {
	p.leftPressed = true
}

func (p *Player) steerLeftReleased() {
	p.leftPressed = false
}

func (p *Player) steerRightPressed() {
	p.rightPressed = true
}

func (p *Player) steerRightReleased() {
	p.rightPressed = false
}

func (p *Player) Equip(guns []*Gun) {
	p.guns = guns
}

func (p *Player) Fire() iter.Seq[*Bullet] {
	if p.hit {
		return slices.Values([]*Bullet{})
	}

	var fired []*Bullet

	for i, gun := range p.guns {
		bulletPos := p.pos.AdjustY(p.Size / 2)

		if len(p.guns) > 1 {
			if i == 0 {
				bulletPos = bulletPos.AdjustY(-p.Size / 3)
			} else {
				bulletPos = bulletPos.AdjustY(p.Size / 3)
			}
		}

		if bullet := gun.Fire(bulletPos, DirectionRightWay); bullet != nil {
			fired = append(fired, bullet)
		}
	}

	return slices.Values(fired)
}
