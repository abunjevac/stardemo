package menu

import (
	"math"

	"stardemo/gk"
)

type Scroller struct {
	renderer    *gk.Renderer
	bounds      gk.Rect
	baseY       int32
	backTex     *gk.Texture
	textTex     *gk.Texture
	scrollSpeed float32
	scrollX     float32
	textWidth   float32

	initialRotation float32
	rotation        float32
	rotationSpeed   float32
	amplitude       float32
	isPaused        bool
	pauseTimer      float32
	rotationAmount  float32
}

func NewScroller(
	renderer *gk.Renderer,
	bounds gk.Rect,
	backTex, textTex *gk.Texture,
	speed, rotation float32,
) *Scroller {
	// use initialRotation to offset the phase timing
	// rotation cycle: 2s pause + 2s rotation = 4s total
	phaseOffset := (rotation / 360.0) * 4.0

	var isPaused bool
	var pauseTimer, rotationAmount, rot float32

	if phaseOffset < 2.0 {
		// start in pause phase
		isPaused = true
		pauseTimer = phaseOffset
		rot = 360.0 - rotation
		rotationAmount = 0.0
	} else {
		// start in rotation phase
		isPaused = false
		pauseTimer = 0.0
		rotationAmount = (phaseOffset - 2.0) * 180.0 // 180°/s rotation speed
		rot = 360.0 - rotation + rotationAmount

		if rot >= 360.0 {
			rot -= 360.0
		}
	}

	return &Scroller{
		renderer:        renderer,
		bounds:          bounds,
		baseY:           bounds.Y,
		backTex:         backTex,
		textTex:         textTex,
		scrollSpeed:     speed,
		scrollX:         float32(bounds.W),
		textWidth:       float32(textTex.W),
		initialRotation: rotation,
		rotation:        rot,
		rotationSpeed:   180.0,
		amplitude:       40.0,
		isPaused:        isPaused,
		pauseTimer:      pauseTimer,
		rotationAmount:  rotationAmount,
	}
}

func (s *Scroller) Update(delta uint64) {
	deltaF := float32(delta)

	if s.scrollX+s.textWidth < 0 {
		s.scrollX += s.textWidth
	}

	s.scrollX -= s.scrollSpeed * deltaF / 100.0

	if s.isPaused {
		s.pauseTimer += deltaF / 1000.0

		if s.pauseTimer >= 2.0 {
			s.isPaused = false
			s.pauseTimer = 0.0
			s.rotationAmount = 0.0

			// start rotation so that (rotation + initialRotation) = 0, ensuring smooth start
			s.rotation = 360.0 - s.initialRotation
		}
	} else {
		deltaRotation := s.rotationSpeed * deltaF / 1000.0
		s.rotation += deltaRotation
		s.rotationAmount += deltaRotation

		if s.rotation >= 360.0 {
			s.rotation -= 360.0
		}

		// check if completed full 360° cycle
		if s.rotationAmount >= 360.0 {
			// reset to starting position where (rotation + initialRotation) = 0
			s.rotation = 360.0 - s.initialRotation
			s.isPaused = true
		}
	}
}

func (s *Scroller) Render() {
	var verticalOffset float32

	if s.isPaused {
		verticalOffset = 0
	} else {
		// apply initial rotation as phase offset for different wave patterns
		radians := (s.rotation + s.initialRotation) * math.Pi / 180.0

		verticalOffset = float32(math.Sin(float64(radians))) * s.amplitude
	}

	currentY := float32(s.baseY) + verticalOffset

	backDst := gk.NewFRect(
		float32(s.bounds.X),
		currentY,
		float32(s.bounds.W),
		float32(s.bounds.H),
	)

	s.renderer.CopyRotateF(s.backTex, backDst, 0, gk.FlipNone)

	x := s.scrollX
	y := currentY + float32(s.bounds.H-s.textTex.H)/2

	for int32(x) < s.bounds.W {
		s.renderer.CopyDirectF(s.textTex, gk.FPos{
			X: float32(s.bounds.X) + x,
			Y: y,
		})

		x += s.textWidth
	}
}

func (s *Scroller) Destroy() {
	s.backTex.Destroy()
	s.textTex.Destroy()
}
