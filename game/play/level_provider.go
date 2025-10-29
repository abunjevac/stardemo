package play

type LevelProvider struct {
	entityBuilder *EntityBuilder
	builders      []func(*EntityBuilder) *Level
}

func NewLevelProvider(entityBuilder *EntityBuilder) *LevelProvider {
	return &LevelProvider{
		entityBuilder: entityBuilder,
		builders: []func(*EntityBuilder) *Level{
			NewLevel1,
			NewLevel2,
			NewLevel3,
			NewLevel4,
			NewLevel5,
		},
	}
}

func (p *LevelProvider) Get(level int) *Level {
	return p.builders[level](p.entityBuilder)
}

func (p *LevelProvider) Count() int {
	return len(p.builders)
}
