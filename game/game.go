package game

import (
	"stardemo/game/activity"
	"stardemo/game/menu"
	"stardemo/game/play"
	"stardemo/gk"
	"stardemo/ui"
)

type Game struct {
	renderer    *gk.Renderer
	keyboard    *gk.Keyboard
	textFactory *gk.TextFactory
	defaultFont *gk.Font
	running     bool
	intent      activity.Intent
	stage       Stage
	buildMenu   func() *menu.Menu
	buildPlay   func() *play.Play
}

func New(ctx *ui.Context) *Game {
	return &Game{
		renderer:    ctx.Renderer,
		keyboard:    ctx.Keyboard,
		textFactory: ctx.TextFactory,
		defaultFont: ctx.DefaultFont,
		running:     true,
		buildMenu: func() *menu.Menu {
			return menu.New(ctx)
		},
		buildPlay: func() *play.Play {
			return play.New(ctx)
		},
	}
}

func (g *Game) Run() {
	g.replaceStage(g.buildMenu())
	g.runLoop()
	g.replaceStage(nil)
}

func (g *Game) runLoop() {
	const fps = 60

	ticker := gk.NewTicker(fps)

	for g.running {
		ticker.Mark()

		g.processInput()
		g.processIntent()
		g.update(ticker.FrameStep)
		g.render()

		ticker.Yield()
	}
}

func (g *Game) update(delta uint64) {
	if s := g.stage; s != nil {
		s.Update(delta)
	}
}

func (g *Game) render() {
	r := g.renderer

	r.Clear()

	if s := g.stage; s != nil {
		s.Render()
	}

	r.Present()
}

func (g *Game) processIntent() {
	switch g.intent {
	case activity.IntentQuit:
		g.running = false
	case activity.IntentMenu:
		g.replaceStage(g.buildMenu())
	case activity.IntentPlay:
		g.replaceStage(g.buildPlay())
	default:
	}

	g.intent = activity.IntentNone
}

func (g *Game) processInput() {
	switch action := pollAction(g.keyboard); action {
	case activity.ActionQuit:
		g.intent = activity.IntentQuit

		return
	case activity.ActionNone:
		g.intent = activity.IntentNone

		return
	default:
		if s := g.stage; s != nil {
			g.intent = s.Act(action)

			return
		}

		g.intent = activity.IntentNone
	}
}

func (g *Game) replaceStage(newStage Stage) {
	if s := g.stage; s != nil {
		s.Destroy()
	}

	g.stage = newStage
}
