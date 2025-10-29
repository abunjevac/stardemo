package play

import (
	"time"

	"stardemo/gk"
)

type Gun struct {
	bulletFactory *BulletFactory
	bulletType    BulletType
	faction       Faction
	fireDelay     int
	maxPerSecond  int
	lastShot      int64
	windowStart   int64
	firedInWindow int
}

func NewGun(bulletFactory *BulletFactory, bulletType BulletType, fireDelay int, maxPerSecond int, faction Faction) *Gun {
	return &Gun{
		bulletFactory: bulletFactory,
		bulletType:    bulletType,
		fireDelay:     fireDelay,
		maxPerSecond:  maxPerSecond,
		faction:       faction,
		lastShot:      0,
		windowStart:   0,
		firedInWindow: 0,
	}
}

func (g *Gun) Fire(pos gk.Pos, direction Direction) *Bullet {
	now := time.Now().UnixMilli()

	// reset 1-second window
	if g.windowStart == 0 || now-g.windowStart >= 1000 {
		g.windowStart = now
		g.firedInWindow = 0
	}

	// cap by max shots per second
	if g.firedInWindow >= g.maxPerSecond {
		return nil
	}

	// cap by minimal time between shots (fireDelay in ms)
	if now-g.lastShot < int64(g.fireDelay) {
		return nil
	}

	g.lastShot = now

	g.firedInWindow++

	return g.bulletFactory.NewBullet(g.bulletType, pos, direction, g.faction)
}
