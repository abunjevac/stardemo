package play

import (
	"stardemo/gk"
)

type Planet struct {
	Pos gk.Pos
	Tex *gk.Texture
}

type Scenery struct {
	renderer *gk.Renderer
	planets  []Planet
}

func NewScenery(renderer *gk.Renderer, bounds gk.Rect) *Scenery {
	return &Scenery{
		renderer: renderer,
		planets: []Planet{
			{
				Pos: gk.NewPos(bounds.W-200, 120),
				Tex: func() *gk.Texture {
					t := renderer.LoadImageTextureScaled("gas11", 0.6)

					t.SetAlphaMod(200)

					return t
				}(),
			},
			{
				Pos: gk.NewPos(200, bounds.H-200),
				Tex: func() *gk.Texture {
					t := renderer.LoadImageTextureScaled("pontes", 0.5)

					t.SetAlphaMod(170)

					return t
				}(),
			},
			{
				Pos: gk.NewPos(bounds.W/2-50, bounds.H/2+80),
				Tex: func() *gk.Texture {
					t := renderer.LoadImageTextureScaled("yniu", 0.4)

					t.SetAlphaMod(140)

					return t
				}(),
			},
		},
	}
}

func (s *Scenery) Render() {
	for _, planet := range s.planets {
		s.renderer.CopyDirect(planet.Tex, planet.Pos)
	}
}
