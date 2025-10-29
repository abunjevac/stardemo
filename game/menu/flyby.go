package menu

import (
	"math"

	"stardemo/gk"
)

type FlyBy struct {
	renderer *gk.Renderer
	bounds   gk.Rect
	hero     *gk.Texture
	time     float64
}

func NewFlyBy(renderer *gk.Renderer, bounds gk.Rect, images *gk.ImageCache) *FlyBy {
	return &FlyBy{
		renderer: renderer,
		bounds:   bounds,
		hero:     images.GetScaled("starship", 0.3),
	}
}

func (f *FlyBy) Update(delta uint64) {
	// increment time for animation, speed controls how fast the figure-eight is traced
	f.time += float64(delta) / 1000
}

func (f *FlyBy) Render() {
	// parametric equations: x = a*sin(t), y = b*sin(t)*cos(t)
	t := f.time

	// scale factors for the figure-eight size
	a := float32(f.bounds.W) * 0.35
	b := float32(f.bounds.H) * 0.35

	// center position
	cx := float32(f.bounds.X) + float32(f.bounds.W)/2
	cy := float32(f.bounds.Y) + float32(f.bounds.H)/2

	// calculate position on figure-eight
	x := cx + a*float32(math.Sin(t))
	y := cy + b*float32(math.Sin(t)*math.Cos(t))

	// calculate velocity vector for rotation (derivatives of parametric equations)
	dx := a * float32(math.Cos(t))
	dy := b * float32(math.Cos(2*t))

	// calculate rotation angle based on direction of movement
	// add 90 degrees (π/2) because the hero image is facing up (not right)
	angle := math.Atan2(float64(dy), float64(dx))*180/math.Pi + 90

	w := float32(f.hero.W)
	h := float32(f.hero.H)

	f.renderer.CopyRotateF(f.hero, gk.FRect{
		X: x - w/2,
		Y: y - h/2,
		W: w,
		H: h,
	}, angle, gk.FlipNone)
}

func (f *FlyBy) Destroy() {}
