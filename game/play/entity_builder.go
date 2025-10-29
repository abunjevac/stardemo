package play

import (
	"stardemo/gk"
)

type EntityBuilder struct {
	imageCache   *gk.ImageCache
	gunFactory   *GunFactory
	crateFactory *CrateFactory
	bounds       gk.Rect

	nextId int
}

func NewLevelBuilder(
	imageCache *gk.ImageCache,
	gunFactory *GunFactory,
	crateFactory *CrateFactory,
	height gk.Rect,
) *EntityBuilder {
	return &EntityBuilder{
		imageCache:   imageCache,
		gunFactory:   gunFactory,
		crateFactory: crateFactory,
		bounds:       height,
	}
}

func (l *EntityBuilder) NewEnemy(
	shipName string,
	gunTypes []GunType,
	pos gk.Pos,
	updater PositionUpdater,
	score int,
) *Enemy {
	l.nextId++

	ship := l.newShip(shipName, pos)
	guns := l.newGuns(gunTypes...)

	return NewEnemy(l.nextId, ship, guns, updater, score)
}

func (l *EntityBuilder) newShip(image string, pos gk.Pos) *Ship {
	return NewShip(
		l.imageCache.GetRotated(image, -90),
		l.imageCache.GetRotated("trail", -90),
		DirectionWrongWay,
		FactionThem,
		pos,
	)
}

func (l *EntityBuilder) newGuns(gunType ...GunType) []*Gun {
	guns := make([]*Gun, 0, 2)

	for _, gt := range gunType {
		guns = append(guns, l.gunFactory.NewGun(gt, FactionThem))
	}

	return guns
}

func (l *EntityBuilder) NewCrate(pos gk.Pos, speed int32, gunTypes []GunType) *Crate {
	return l.crateFactory.NewCrate(pos, speed, gunTypes)
}

func (l *EntityBuilder) Height() int32 {
	return l.bounds.H
}

func (l *EntityBuilder) Position(x, y int32) gk.Pos {
	const sceneEntryDelay = 100

	return gk.NewPos(l.bounds.W+sceneEntryDelay+x, y)
}
