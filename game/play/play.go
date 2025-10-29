package play

import (
	"fmt"

	"stardemo/game/activity"
	"stardemo/gk"
	"stardemo/ui"
)

type Play struct {
	starfield        *Starfield
	scenery          *Scenery
	hud              *HUD
	announcement     *Announcement
	stats            *Stats
	bullets          *Bullets
	enemies          *Enemies
	crates           *Crates
	levelProvider    *LevelProvider
	activeArea       gk.Rect
	activateLevel    bool
	currentLevel     int
	player           *Player
	buildPlayer      func() *Player
	buildHitEffect   func() []*gk.Texture
	buildCrateEffect func() []*gk.Texture
}

func New(ctx *ui.Context) *Play {
	const (
		reservedForHUD   = 70
		bottomSlackSpace = 10
	)

	stats := &Stats{Score: 0, Lives: 3}

	activeArea := gk.NewRect(
		ctx.Bounds.X,
		ctx.Bounds.Y+reservedForHUD,
		ctx.Bounds.W,
		ctx.Bounds.H-reservedForHUD-bottomSlackSpace,
	)

	bulletFactory := NewBulletFactory(ctx.ImageCache)
	gunFactory := NewGunFactory(bulletFactory)
	crateFactory := NewCrateFactory(ctx.Effects, gunFactory)
	entityBuilder := NewLevelBuilder(ctx.ImageCache, gunFactory, crateFactory, activeArea)

	return &Play{
		starfield:     NewStarfield(ctx.Renderer, ctx.Bounds),
		scenery:       NewScenery(ctx.Renderer, ctx.Bounds),
		hud:           NewHUD(ctx.Renderer, ctx.Bounds, ctx.FontCache, ctx.TextFactory, ctx.ImageCache, stats),
		announcement:  NewAnnouncement(ctx.Renderer, ctx.Bounds, ctx.FontCache, ctx.TextFactory),
		bullets:       NewBullets(ctx.Renderer),
		enemies:       NewEnemies(ctx.Renderer),
		crates:        NewCrates(ctx.Renderer),
		levelProvider: NewLevelProvider(entityBuilder),
		stats:         stats,
		activeArea:    activeArea,
		activateLevel: true,
		currentLevel:  0,
		buildPlayer: func() *Player {
			player := NewPlayer(ctx.Renderer, activeArea, ctx.ImageCache, stats)

			player.Equip([]*Gun{
				gunFactory.NewGun(GunTypeMissileLauncher, FactionUs),
			})

			return player
		},
		buildHitEffect: func() []*gk.Texture {
			return ctx.Effects.Get("impact")
		},
		buildCrateEffect: func() []*gk.Texture {
			return ctx.Effects.Get("box")
		},
	}
}

func (p *Play) Update(delta uint64) {
	if p.player != nil && p.player.destroyed {
		p.player = nil

		p.stats.Lives--

		// leave enemies/crates flying around if there are no lives left
		if p.stats.Lives > 0 {
			p.enemies.Destroy()
			p.crates.Destroy()
		}
	}

	if p.player == nil && p.stats.Lives > 0 {
		p.player = p.buildPlayer()
		p.activateLevel = true
	}

	if p.activateLevel {
		p.activateLevel = false

		p.announcement.AnimateText(fmt.Sprintf("Level %d", p.currentLevel+1))

		p.startLevel()
	}

	if p.enemies.WipedOut() && p.crates.AllCollected() && p.player != nil {
		p.currentLevel++

		p.stats.Score += 100 * p.currentLevel

		if p.currentLevel >= p.levelProvider.Count() {
			p.currentLevel = 0
		}

		p.activateLevel = true
	}

	p.starfield.Update(delta)
	p.hud.Update(delta)
	p.announcement.Update(delta)

	p.bullets.Update(delta, p.activeArea)
	p.enemies.Update(delta)
	p.crates.Update(delta)

	if p.player == nil {
		return
	}

	p.player.Update(delta, p.activeArea)

	p.checkDamagesFromEnemies()

	if p.player.IsHit() {
		return
	}

	p.processEnemyFiring()
	p.checkShipCollisions()
	p.checkDamagesToEnemies()
	p.collectCrates()
}

func (p *Play) Render() {
	p.starfield.Render()
	p.scenery.Render()
	p.hud.Render()
	p.bullets.Render(p.activeArea)
	p.enemies.Render(p.activeArea)
	p.crates.Render(p.activeArea)

	if p.player == nil {
		p.announcement.RenderText("Oh no! We are doomed!", 255)

		return
	}

	p.announcement.Render()
	p.player.Render(p.activeArea)
}

func (p *Play) Destroy() {
	p.starfield.Destroy()
	p.hud.Destroy()
	p.announcement.Destroy()
	p.bullets.Destroy()
	p.enemies.Destroy()
	p.crates.Destroy()

	if p.player != nil {
		p.player.Destroy()
	}
}

func (p *Play) Act(action activity.Action) activity.Intent {
	if p.stats.Lives == 0 {
		return activity.IntentMenu
	}

	switch action {
	case activity.ActionUpPressed:
		p.player.steerUpPressed()
	case activity.ActionDownPressed:
		p.player.steerDownPressed()
	case activity.ActionUpReleased:
		p.player.steerUpReleased()
	case activity.ActionDownReleased:
		p.player.steerDownReleased()
	case activity.ActionLeftPressed:
		p.player.steerLeftPressed()
	case activity.ActionRightPressed:
		p.player.steerRightPressed()
	case activity.ActionLeftReleased:
		p.player.steerLeftReleased()
	case activity.ActionRightReleased:
		p.player.steerRightReleased()
	case activity.ActionFire:
		p.processPlayerFiring()
	case activity.ActionEscape:
		return activity.IntentMenu
	default:
	}

	return activity.IntentNone
}

func (p *Play) processPlayerFiring() {
	if p.player == nil {
		return
	}

	for bullet := range p.player.Fire() {
		p.bullets.Add(bullet)
	}
}

func (p *Play) processEnemyFiring() {
	for e := range p.enemies.All() {
		if !e.IsVisible(p.activeArea) {
			continue
		}

		for bullet := range e.Fire() {
			p.bullets.Add(bullet)
		}
	}
}

func (p *Play) checkShipCollisions() {
	player := p.player

	for enemy := range p.enemies.All() {
		if player.CollidesWith(enemy.Ship) {
			player.Wrecked(p.buildHitEffect())
			enemy.Wrecked(p.buildHitEffect())

			return
		}
	}
}

func (p *Play) checkDamagesFromEnemies() {
	player := p.player

	for bullet := range p.bullets.All() {
		if player.TakesDamageFrom(bullet) {
			player.Wrecked(p.buildHitEffect())
			bullet.Destroy()

			return
		}
	}
}

func (p *Play) checkDamagesToEnemies() {
	for bullet := range p.bullets.All() {
		for enemy := range p.enemies.All() {
			if !enemy.TakesDamageFrom(bullet) {
				continue
			}

			enemy.Wrecked(p.buildHitEffect())
			bullet.Destroy()

			score := enemy.Score
			dist := gk.DistPos(p.player.pos, enemy.pos)
			reward := gk.Clamp(dist, 0, 2000) * float64(score) / 2000

			p.stats.Score += score + int(reward)

			return
		}
	}
}

func (p *Play) collectCrates() {
	player := p.player

	for crate := range p.crates.All() {
		if player.CanCollect(crate) {
			player.Equip(crate.Guns)
			crate.Destroy()

			return
		}
	}
}

func (p *Play) startLevel() {
	level := p.levelProvider.Get(p.currentLevel)

	if level == nil {
		return
	}

	for _, enemy := range level.SpawnEnemies() {
		p.enemies.Add(enemy)
	}

	for _, crate := range level.SpawnCrates() {
		p.crates.Add(crate)
	}
}
