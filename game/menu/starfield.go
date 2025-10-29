package menu

import (
	"math"

	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

// RotationDirection represents the direction of starfield rotation
type RotationDirection int

const (
	RotateNone RotationDirection = iota
	RotateLeft
	RotateRight
)

type Star struct {
	x       float64
	y       float64
	z       float64
	speed   float64
	size    int32
	color   sdl.Color
	angle   float64
	active  bool
	centerX float64
	centerY float64
}

type StarKey struct {
	color sdl.Color
	size  gk.Size
}

type Starfield struct {
	renderer         *gk.Renderer
	bounds           gk.Rect
	starNum          int
	stars            []*Star
	maxZ             float64
	minZ             float64
	maxSize          int32
	minSize          int32
	centerX          float64
	centerY          float64
	rotationDir      RotationDirection
	rotationTimer    float64
	rotationInterval float64
	rotationSpeed    float64
	textureCache     map[StarKey]*gk.Texture
}

func NewStarfield(renderer *gk.Renderer, bounds gk.Rect) *Starfield {
	const starNum = 600

	centerX := float64(bounds.W) / 2
	centerY := float64(bounds.H) / 2

	s := &Starfield{
		renderer:         renderer,
		bounds:           bounds,
		starNum:          starNum,
		stars:            make([]*Star, starNum),
		maxZ:             1000.0,
		minZ:             1.0,
		maxSize:          4,
		minSize:          1,
		centerX:          centerX,
		centerY:          centerY,
		rotationDir:      RotateNone,
		rotationTimer:    0,
		rotationInterval: 2.0, // change direction every 2 seconds
		rotationSpeed:    0.3, // rotation speed in radians per second
		textureCache:     make(map[StarKey]*gk.Texture),
	}

	// initialize stars with scattered positions across the screen
	for i := 0; i < starNum; i++ {
		// random angle for direction from center
		angle := gk.RandAngle()

		// Calculate a random distance from center that scatters stars across the screen
		// but leaves a small space around the center
		maxDist := math.Max(float64(bounds.W), float64(bounds.H)) / 2
		minDist := maxDist * 0.1 // leave 10% space around center
		distFromCenter := gk.Rand(minDist, maxDist)

		s.stars[i] = &Star{
			x:       centerX + math.Cos(angle)*distFromCenter,
			y:       centerY + math.Sin(angle)*distFromCenter,
			z:       gk.Rand(s.minZ, s.maxZ),
			speed:   s.randSpeed(),
			angle:   angle,
			active:  true,
			centerX: centerX,
			centerY: centerY,
		}

		// set initial size and color based on z position
		s.updateStarProperties(s.stars[i])
	}

	return s
}

func (s *Starfield) Update(delta uint64) {
	deltaSeconds := float64(delta) / 1000.0

	s.rotationTimer += deltaSeconds

	if s.rotationTimer >= s.rotationInterval {
		s.rotationTimer = 0
		s.rotationDir = RotationDirection(gk.RandN(3))
	}

	var rotationAngle float64

	switch s.rotationDir {
	case RotateLeft:
		rotationAngle = s.rotationSpeed * deltaSeconds
	case RotateRight:
		rotationAngle = -s.rotationSpeed * deltaSeconds
	case RotateNone:
		rotationAngle = 0
	}

	for _, star := range s.stars {
		if !star.active {
			continue
		}

		if rotationAngle != 0 {
			relX := star.x - s.centerX
			relY := star.y - s.centerY

			// rotation matrix
			cosTheta := math.Cos(rotationAngle)
			sinTheta := math.Sin(rotationAngle)
			newRelX := relX*cosTheta - relY*sinTheta
			newRelY := relX*sinTheta + relY*cosTheta

			star.x = s.centerX + newRelX
			star.y = s.centerY + newRelY

			star.angle += rotationAngle
		}

		// calculate movement speed based on distance from center
		// stars move faster as they get further from center
		distFromCenter := gk.Dist(star.x, star.y, s.centerX, s.centerY)
		speedFactor := 1.0 + distFromCenter/100.0

		// move star outward from center along its angle
		moveAmount := star.speed * speedFactor * deltaSeconds * 150.0

		star.x += math.Cos(star.angle) * moveAmount
		star.y += math.Sin(star.angle) * moveAmount

		if s.isOutOfBounds(star.x, star.y) {
			star.active = false
		}

		// update size and color based on distance from center
		// further from center = closer to viewer + larger and brighter
		s.updateStarProperties(star)
	}

	for _, star := range s.stars {
		if !star.active {
			s.resetStar(star)
		}
	}
}

func (s *Starfield) Render() {
	for _, star := range s.stars {
		if !star.active {
			continue
		}

		screenX := int32(star.x)
		screenY := int32(star.y)

		starRect := gk.NewRect(screenX, screenY, star.size, star.size)

		key := StarKey{
			color: star.color,
			size:  starRect.Size(),
		}

		starTexture, found := s.textureCache[key]
		if !found {
			starTexture = s.renderer.CreateRectTexture(star.color, starRect.Size())

			s.textureCache[key] = starTexture
		}

		s.renderer.Copy(starTexture, starRect)
	}
}

func (s *Starfield) Destroy() {
	for _, tex := range s.textureCache {
		tex.Destroy()
	}

	clear(s.textureCache)
}

func (s *Starfield) updateStarProperties(star *Star) {
	// calculate size based on z position (stars further away are smaller)
	zFactor := 1.0 - (star.z-s.minZ)/(s.maxZ-s.minZ)

	star.size = s.minSize + int32(float64(s.maxSize-s.minSize)*zFactor)

	// calculate brightness based on z position (stars further away are dimmer)
	brightness := uint8(100.0*zFactor + 100.0)

	star.color = sdl.Color{R: brightness, G: brightness, B: brightness, A: 255}
}

func (s *Starfield) resetStar(star *Star) {
	star.angle = gk.RandAngle()

	// start near the center, but not exactly at it
	minDist := 10.0  // minimum distance from center
	maxDist := 300.0 // maximum distance from center
	distFromCenter := gk.Rand(minDist, maxDist)

	star.x = s.centerX + math.Cos(star.angle)*distFromCenter
	star.y = s.centerY + math.Sin(star.angle)*distFromCenter

	star.z = gk.Rand(s.minZ, s.maxZ)

	star.speed = s.randSpeed()

	star.active = true

	s.updateStarProperties(star)
}

func (s *Starfield) isOutOfBounds(x, y float64) bool {
	return x < 0 || x > float64(s.bounds.W) || y < 0 || y > float64(s.bounds.H)
}

func (s *Starfield) randSpeed() float64 {
	return gk.Rand(0.1, 0.4)
}
