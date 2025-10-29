package play

import (
	"math"

	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type Orientation int

const (
	OrientationHorizontal Orientation = iota
	OrientationVertical
)

type ThrustIndicator struct {
	renderer *gk.Renderer
	bounds   gk.Rect
	value    func() int
	orient   Orientation

	animTimer  float64
	barBg      *gk.Texture
	barForward *gk.Texture
	barReverse *gk.Texture
	barNeutral *gk.Texture
}

func NewThrustIndicator(renderer *gk.Renderer, bounds gk.Rect, value func() int, orient Orientation) *ThrustIndicator {
	createThrustBar := func(color sdl.Color) *gk.Texture {
		return renderer.CreateRectTexture(
			color,
			gk.Size{W: 1, H: 1},
		)
	}

	return &ThrustIndicator{
		renderer: renderer,
		bounds:   bounds,
		value:    value,
		orient:   orient,

		animTimer:  0,
		barBg:      createThrustBar(sdl.Color{R: 60, G: 60, B: 60, A: 150}),
		barForward: createThrustBar(sdl.Color{R: 0, G: 150, B: 255, A: 220}),
		barReverse: createThrustBar(sdl.Color{R: 255, G: 80, B: 0, A: 220}),
		barNeutral: createThrustBar(sdl.Color{R: 100, G: 100, B: 100, A: 180}),
	}
}

func (t *ThrustIndicator) Update(delta uint64) {
	t.animTimer += float64(delta) / 1000.0
}

func (t *ThrustIndicator) Render() {
	const (
		thrustBarWidth   = int32(200)
		thrustBarHeight  = int32(20)
		thrustSegments   = int32(20)
		thrustSegmentGap = int32(2)
	)

	thrustPct := t.value()

	halfSegments := thrustSegments / 2
	activeSeg := int32(math.Abs(float64(thrustPct)) / 100.0 * float64(halfSegments))
	pulse := 0.7 + 0.3*math.Sin(t.animTimer*8)

	if t.orient == OrientationHorizontal {
		segmentWidth := (thrustBarWidth - (thrustSegments-1)*thrustSegmentGap) / thrustSegments
		x0 := t.bounds.X + (t.bounds.W-thrustBarWidth)/2

		for i := int32(0); i < thrustSegments; i++ {
			x := x0 + i*(segmentWidth+thrustSegmentGap)

			var texture *gk.Texture

			if i < halfSegments {
				segmentIndex := halfSegments - 1 - i

				if thrustPct < 0 && segmentIndex < activeSeg {
					texture = t.barReverse
				} else {
					texture = t.barBg
				}
			} else {
				segmentIndex := i - halfSegments

				if thrustPct > 0 && segmentIndex < activeSeg {
					texture = t.barForward
				} else {
					texture = t.barBg
				}
			}

			height := thrustBarHeight + int32(pulse*16)

			t.renderer.Copy(texture, gk.Rect{
				X: x,
				Y: t.bounds.Y,
				W: segmentWidth,
				H: height,
			})
		}
	} else { // OrientationVertical
		segmentHeight := (thrustBarWidth - (thrustSegments-1)*thrustSegmentGap) / thrustSegments
		y0 := t.bounds.Y + (t.bounds.H-thrustBarWidth)/2
		width := thrustBarHeight + int32(pulse*16)

		for i := int32(0); i < thrustSegments; i++ {
			y := y0 + i*(segmentHeight+thrustSegmentGap)

			var texture *gk.Texture

			if i < halfSegments {
				segmentIndex := halfSegments - 1 - i

				if thrustPct < 0 && segmentIndex < activeSeg {
					texture = t.barReverse
				} else {
					texture = t.barBg
				}
			} else {
				segmentIndex := i - halfSegments

				if thrustPct > 0 && segmentIndex < activeSeg {
					texture = t.barForward
				} else {
					texture = t.barBg
				}
			}

			t.renderer.Copy(texture, gk.Rect{
				X: t.bounds.X,
				Y: y,
				W: width,
				H: segmentHeight,
			})
		}
	}
}

func (t *ThrustIndicator) Destroy() {
	t.barBg.Destroy()
	t.barForward.Destroy()
	t.barReverse.Destroy()
	t.barNeutral.Destroy()
}
