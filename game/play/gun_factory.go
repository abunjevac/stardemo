package play

type GunType int

const (
	GunTypeMissileLauncher = iota
	GunTypeMissileBLauncher
	GunTypeAcuitLauncher
)

type GunSpec struct {
	bulletType   BulletType
	fireDelay    int
	maxPerSecond int
}

type GunFactory struct {
	bulletFactory *BulletFactory
	specs         map[GunType]GunSpec
}

func NewGunFactory(bulletFactory *BulletFactory) *GunFactory {
	return &GunFactory{
		bulletFactory: bulletFactory,
		specs: map[GunType]GunSpec{
			GunTypeMissileLauncher:  {bulletType: BulletTypeMissile, fireDelay: 300, maxPerSecond: 3},
			GunTypeMissileBLauncher: {bulletType: BulletTypeMissileB, fireDelay: 200, maxPerSecond: 3},
			GunTypeAcuitLauncher:    {bulletType: BulletTypeAcuit, fireDelay: 100, maxPerSecond: 4},
		},
	}
}

func (f *GunFactory) NewGun(gunType GunType, faction Faction) *Gun {
	spec := f.specs[gunType]

	return NewGun(f.bulletFactory, spec.bulletType, spec.fireDelay, spec.maxPerSecond, faction)
}
