package play

import "stardemo/gk"

type BulletType int

const (
	BulletTypeMissile BulletType = iota
	BulletTypeMissileB
	BulletTypeAcuit
)

type BulletSpec struct {
	image map[Direction]*gk.Texture
	speed int32
}

type BulletFactory struct {
	specs  map[BulletType]BulletSpec
	nextId int
}

func NewBulletFactory(imageCache *gk.ImageCache) *BulletFactory {
	return &BulletFactory{
		specs: map[BulletType]BulletSpec{
			BulletTypeMissile: {
				image: map[Direction]*gk.Texture{
					DirectionRightWay: imageCache.GetRotated("missile", 90),
					DirectionWrongWay: imageCache.GetRotated("missile", -90),
				},
				speed: 7,
			},
			BulletTypeMissileB: {
				image: map[Direction]*gk.Texture{
					DirectionRightWay: imageCache.GetRotated("missile-b", 90),
					DirectionWrongWay: imageCache.GetRotated("missile-b", -90),
				},
				speed: 10,
			},
			BulletTypeAcuit: {
				image: map[Direction]*gk.Texture{
					DirectionRightWay: imageCache.GetRotated("acuit", 90),
					DirectionWrongWay: imageCache.GetRotated("acuit", -90),
				},
				speed: 15,
			},
		},
	}
}

func (f *BulletFactory) NewBullet(
	bulletType BulletType,
	pos gk.Pos,
	direction Direction,
	faction Faction,
) *Bullet {
	f.nextId++

	bulletSpec := f.specs[bulletType]

	return NewBullet(f.nextId, bulletSpec.image[direction], faction, pos, direction, bulletSpec.speed)
}
