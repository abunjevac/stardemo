package play

import (
	"sync"

	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

type Star struct {
	x     float64
	y     float64
	speed float64
	size  int32
	color sdl.Color
}

type Starfield struct {
	renderer     *gk.Renderer
	bounds       gk.Rect
	stars        []*Star
	textureCache map[sdl.Color]*gk.Texture
	starPool     *sync.Pool
}

func NewStarfield(renderer *gk.Renderer, bounds gk.Rect) *Starfield {
	const starCount = 400

	s := &Starfield{
		renderer:     renderer,
		bounds:       bounds,
		stars:        make([]*Star, starCount),
		textureCache: make(map[sdl.Color]*gk.Texture),
		starPool: &sync.Pool{
			New: func() interface{} {
				return &Star{}
			},
		},
	}

	for i := 0; i < starCount; i++ {
		star := s.createStar()

		// scatter stars across the screen
		star.x = gk.Rand(0, float64(s.bounds.W))

		s.stars[i] = star
	}

	return s
}

func (s *Starfield) Update(delta uint64) {
	deltaSeconds := float64(delta) / 1000.0

	for i, star := range s.stars {
		star.x -= star.speed * deltaSeconds

		if star.x < -float64(star.size) {
			s.resetStar(i)
		}
	}
}

func (s *Starfield) Render() {
	for _, star := range s.stars {
		if star.x < -float64(star.size) || star.x > float64(s.bounds.W) {
			continue
		}

		screenX := int32(star.x)
		screenY := int32(star.y)

		starRect := gk.NewRect(screenX, screenY, star.size, star.size)

		starTexture, found := s.textureCache[star.color]
		if !found {
			starTexture = s.renderer.CreateRectTexture(star.color, starRect.Size())

			s.textureCache[star.color] = starTexture
		}

		s.renderer.Copy(starTexture, starRect)
	}
}

func (s *Starfield) Destroy() {
	for _, tex := range s.textureCache {
		tex.Destroy()
	}

	clear(s.textureCache)

	for _, star := range s.stars {
		s.starPool.Put(star)
	}
}

func (s *Starfield) resetStar(index int) {
	s.starPool.Put(s.stars[index])

	s.stars[index] = s.createStar()
}

func (s *Starfield) createStar() *Star {
	x := s.bounds.W + 4
	y := gk.Rand(0, float64(s.bounds.H))

	speed := gk.Rand(50, 200)
	brightness := uint8(40 + speed/2)

	star := s.starPool.Get().(*Star)

	star.x = float64(x)
	star.y = y
	star.speed = speed * 2
	star.size = 1 + gk.RandN(4)
	star.color = sdl.Color{R: brightness, G: brightness, B: brightness, A: 255}

	return star
}
