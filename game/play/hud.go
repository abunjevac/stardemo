package play

import (
	"fmt"

	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type HUD struct {
	renderer                  *gk.Renderer
	bounds                    gk.Rect
	stats                     *Stats
	font                      *gk.Font
	textFactory               *gk.TextFactory
	lifeTexture               *gk.Texture
	verticalThrustIndicator   *ThrustIndicator
	horizontalThrustIndicator *ThrustIndicator
	presentedScore            int
	scoreTimer                uint32
	lifeTimer                 float64
	lifeVisible               bool
}

func NewHUD(
	renderer *gk.Renderer,
	bounds gk.Rect,
	fontFactory *gk.FontCache,
	textFactory *gk.TextFactory,
	imageCache *gk.ImageCache,
	stats *Stats,
) *HUD {
	verticalThrustIndBounds := gk.NewRect(bounds.X+10, bounds.H-170, 40, 120)
	horizontalThrustIndBounds := gk.NewRect(bounds.X+95, bounds.H-44, 120, 40)

	verticalThrustValue := func() int {
		return stats.VerticalThrust
	}

	horizontalThrustValue := func() int {
		return stats.HorizontalThrust
	}

	return &HUD{
		renderer:    renderer,
		bounds:      bounds,
		stats:       stats,
		font:        fontFactory.GetFont("zorque", 32).Italic(),
		textFactory: textFactory,

		verticalThrustIndicator: NewThrustIndicator(
			renderer,
			verticalThrustIndBounds,
			verticalThrustValue,
			OrientationVertical,
		),
		horizontalThrustIndicator: NewThrustIndicator(
			renderer,
			horizontalThrustIndBounds,
			horizontalThrustValue,
			OrientationHorizontal,
		),
		lifeTexture: imageCache.Get("starship"),
		lifeVisible: true,
	}
}

func (h *HUD) Update(delta uint64) {
	const lifeFlashRate = 0.5

	if h.presentedScore < h.stats.Score {
		h.scoreTimer += uint32(delta)

		if h.scoreTimer >= 3 {
			h.presentedScore += min(10, h.stats.Score-h.presentedScore)
			h.scoreTimer = 0
		}
	}

	h.lifeTimer += float64(delta) / 1000.0

	if h.lifeTimer >= lifeFlashRate {
		h.lifeVisible = !h.lifeVisible
		h.lifeTimer = 0
	}

	h.verticalThrustIndicator.Update(delta)
	h.horizontalThrustIndicator.Update(delta)
}

func (h *HUD) Render() {
	h.renderScore()
	h.renderLives()
	h.renderTrustIndicators()
}

func (h *HUD) Destroy() {
	h.verticalThrustIndicator.Destroy()
}

func (h *HUD) renderScore() {
	score := h.textFactory.NewShadedText(fmt.Sprintf("%06d", h.presentedScore), h.font, sdl.Color{R: 0, G: 255, B: 255, A: 255})
	defer score.Destroy()

	h.renderer.Copy(score, gk.Rect{
		X: h.bounds.W - 200,
		Y: 10,
		W: score.W,
		H: score.H,
	})
}

func (h *HUD) renderLives() {
	const (
		size    = 40
		spacing = 10
	)

	for i := 0; i < h.stats.Lives; i++ {
		if i+1 == h.stats.Lives && !h.lifeVisible {
			continue
		}

		h.renderer.Copy(h.lifeTexture, gk.Rect{
			X: h.bounds.X + 20 + int32(i)*(size+spacing),
			Y: 10,
			W: size,
			H: size,
		})
	}
}

func (h *HUD) renderTrustIndicators() {
	h.verticalThrustIndicator.Render()
	h.horizontalThrustIndicator.Render()
}
