package play

type Level struct {
	SpawnEnemies func() []*Enemy
	SpawnCrates  func() []*Crate
}
