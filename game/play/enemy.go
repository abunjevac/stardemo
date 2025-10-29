package play

import (
	"iter"
	"slices"
	"time"

	"stardemo/gk"
)

type PositionUpdater func(delta uint64, pos gk.Pos) gk.Pos

type Enemy struct {
	*Ship

	id           int
	guns         []*Gun
	updater      PositionUpdater
	Score        int
	cooling      bool
	coolingStart int64
}

func NewEnemy(
	id int,
	ship *Ship,
	guns []*Gun,
	updater PositionUpdater,
	score int,
) *Enemy {
	return &Enemy{
		id:      id,
		Ship:    ship,
		guns:    guns,
		updater: updater,
		Score:   score,
	}
}

func (e *Enemy) Update(delta uint64) {
	const trailSize = 30

	e.Advance(delta)

	e.pos = e.updater(delta, e.pos)

	if e.pos.X+e.Size+trailSize < 0 {
		e.destroyed = true
	}
}

func (e *Enemy) Render(r *gk.Renderer, bounds gk.Rect) {
	e.RenderShip(r, e.pos.AdjustPos(bounds.Pos()))
}

func (e *Enemy) Fire() iter.Seq[*Bullet] {
	if e.hit {
		return slices.Values([]*Bullet{})
	}

	now := time.Now().UnixMilli()

	if e.cooling {
		if now-e.coolingStart >= 1000 {
			e.cooling = false
		} else {
			return slices.Values([]*Bullet{})
		}
	}

	var fired []*Bullet

	for i, gun := range e.guns {
		bulletPos := e.pos.AdjustY(e.Size / 2)

		if len(e.guns) > 1 {
			if i == 0 {
				bulletPos = bulletPos.AdjustY(-e.Size / 3)
			} else {
				bulletPos = bulletPos.AdjustY(e.Size / 3)
			}
		}

		if bullet := gun.Fire(bulletPos, DirectionWrongWay); bullet != nil {
			fired = append(fired, bullet)
		}
	}

	if fired != nil && !e.cooling {
		if gk.RandN(2000) > 25 {
			e.cooling = true
			e.coolingStart = now
		}
	}

	return slices.Values(fired)
}

func (e *Enemy) IsVisible(area gk.Rect) bool {
	// right side inside area?
	return e.pos.X+e.Size < area.W
}
