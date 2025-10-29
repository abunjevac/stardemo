package play

import (
	"math"

	"stardemo/gk"
)

func advanceSteady(speed int32) func(delta uint64, pos gk.Pos) gk.Pos {
	return func(delta uint64, pos gk.Pos) gk.Pos {
		return pos.AdjustX(-int32(delta) * speed / 40)
	}
}

func sineWave(speed int32, axis, amplitude int32) func(delta uint64, pos gk.Pos) gk.Pos {
	return waveFn(speed, axis, amplitude, math.Sin)
}

func cosineWave(speed int32, axis, amplitude int32) func(delta uint64, pos gk.Pos) gk.Pos {
	return waveFn(speed, axis, amplitude, math.Cos)
}

func waveFn(speed int32, axis, amplitude int32, f func(float64) float64) func(delta uint64, pos gk.Pos) gk.Pos {
	fAxis := float64(axis)
	fAmplitude := float64(amplitude)
	phase := float64(0)

	return func(delta uint64, pos gk.Pos) gk.Pos {
		phase += float64(delta)

		x := pos.X - (int32(delta) * speed / 40)
		y := fAxis + fAmplitude*f(phase/500)

		return gk.NewPos(x, int32(y))
	}
}

func NewLevel1(b *EntityBuilder) *Level {
	return &Level{
		SpawnEnemies: func() []*Enemy {
			return []*Enemy{
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(0, b.Height()/2),
					advanceSteady(5),
					50),
			}
		},
		SpawnCrates: func() []*Crate {
			return []*Crate{
				b.NewCrate(
					b.Position(0, b.Height()/2-100),
					15,
					[]GunType{GunTypeMissileLauncher, GunTypeMissileLauncher},
				),
			}
		},
	}
}

func NewLevel2(b *EntityBuilder) *Level {
	return &Level{
		SpawnEnemies: func() []*Enemy {
			return []*Enemy{
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(0, b.Height()/2),
					advanceSteady(5),
					50),
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(100, b.Height()/2-150),
					advanceSteady(5),
					50),
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(100, b.Height()/2+150),
					advanceSteady(5),
					50),
			}
		},
		SpawnCrates: func() []*Crate {
			return nil
		},
	}
}

func NewLevel3(b *EntityBuilder) *Level {
	return &Level{
		SpawnEnemies: func() []*Enemy {
			return []*Enemy{
				b.NewEnemy(
					ShipClipper,
					[]GunType{GunTypeMissileBLauncher},
					b.Position(0, b.Height()/2),
					sineWave(4, b.Height()/2, 200),
					150),
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(100, b.Height()/2-200),
					advanceSteady(5),
					50),
				b.NewEnemy(
					ShipAerie,
					[]GunType{GunTypeMissileLauncher},
					b.Position(100, b.Height()/2+200),
					advanceSteady(5),
					50),
			}
		},
		SpawnCrates: func() []*Crate {
			return []*Crate{
				b.NewCrate(
					b.Position(200, b.Height()/2),
					15,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
				),
			}
		},
	}
}

func NewLevel4(b *EntityBuilder) *Level {
	return &Level{
		SpawnEnemies: func() []*Enemy {
			return []*Enemy{
				b.NewEnemy(
					ShipFury,
					[]GunType{GunTypeAcuitLauncher},
					b.Position(0, 0),
					sineWave(5, b.Height()/2, 200),
					250),
				b.NewEnemy(
					ShipClipper,
					[]GunType{GunTypeAcuitLauncher},
					b.Position(100, 0),
					sineWave(4, b.Height()/2-200, 200),
					150),
				b.NewEnemy(
					ShipClipper,
					[]GunType{GunTypeAcuitLauncher},
					b.Position(100, 0),
					cosineWave(4, b.Height()/2+200, 200),
					150),
				b.NewEnemy(
					ShipFury,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
					b.Position(300, 0),
					sineWave(4, b.Height()/2-50, 200),
					300),
			}
		},
		SpawnCrates: func() []*Crate {
			return nil
		},
	}
}

func NewLevel5(b *EntityBuilder) *Level {
	return &Level{
		SpawnEnemies: func() []*Enemy {
			return []*Enemy{
				b.NewEnemy(
					ShipFury,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
					b.Position(0, 0),
					cosineWave(4, b.Height()/2+200, 200),
					250),
				b.NewEnemy(
					ShipClipper,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
					b.Position(100, 0),
					sineWave(4, b.Height()/2-200, 200),
					150),
				b.NewEnemy(
					ShipClipper,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
					b.Position(100, 0),
					cosineWave(4, b.Height()/2+200, 200),
					150),
				b.NewEnemy(
					ShipFury,
					[]GunType{GunTypeAcuitLauncher, GunTypeAcuitLauncher},
					b.Position(300, 0),
					cosineWave(4, b.Height()/2+200, 200),
					300),
			}
		},
		SpawnCrates: func() []*Crate {
			return nil
		},
	}
}
