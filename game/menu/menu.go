package menu

import (
	"stardemo/game/activity"
	"stardemo/gk"
	"stardemo/ui"
)

type Menu struct {
	elements []gk.Element
}

func New(ctx *ui.Context) *Menu {
	return &Menu{
		elements: []gk.Element{
			NewStarfield(ctx.Renderer, ctx.Bounds),
			NewBanner(ctx.Renderer, ctx.Bounds),
			NewCredits(ctx.Renderer, ctx.Bounds, ctx.FontCache, ctx.TextFactory),
			NewKicker(ctx.Renderer, ctx.Bounds, ctx.TextFactory, ctx.DefaultFont),
			NewFlyBy(ctx.Renderer, ctx.Bounds, ctx.ImageCache),
		},
	}
}

func (m *Menu) Update(delta uint64) {
	gk.UpdateAll(m.elements, delta)
}

func (m *Menu) Render() {
	gk.RenderAll(m.elements)
}

func (m *Menu) Destroy() {
	gk.DestroyAll(m.elements)
}

func (m *Menu) Act(action activity.Action) activity.Intent {
	if action == activity.ActionFire {
		return activity.IntentPlay
	}

	return activity.IntentNone
}
