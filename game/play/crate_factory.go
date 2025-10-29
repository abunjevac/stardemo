package play

import (
	"stardemo/gk"
)

type CrateFactory struct {
	effects    *gk.Effects
	gunFactory *GunFactory
	nextId     int
}

func NewCrateFactory(effects *gk.Effects, gunFactory *GunFactory) *CrateFactory {
	return &CrateFactory{
		effects:    effects,
		gunFactory: gunFactory,
	}
}

func (f *CrateFactory) NewCrate(pos gk.Pos, speed int32, gunTypes []GunType) *Crate {
	f.nextId++

	carousel := gk.NewImageCarousel(f.effects.Get("box"), 1000, true)
	guns := make([]*Gun, len(gunTypes))

	for i, gunType := range gunTypes {
		guns[i] = f.gunFactory.NewGun(gunType, FactionUs)
	}

	return NewCrate(f.nextId, carousel, pos, speed, guns)
}
