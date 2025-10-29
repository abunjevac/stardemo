package menu

import (
	"math"
	"strings"

	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

var text = `
.#####..#######....#....######...#####..#.....#.###.######
#.....#....#......#.#...#.....#.#.....#.#.....#..#..#.....#
#..........#.....#...#..#.....#.#.......#.....#..#..#.....#
.#####.....#....#.....#.######...#####..#######..#..######
......#....#....#######.#...#.........#.#.....#..#..#
#.....#....#....#.....#.#....#..#.....#.#.....#..#..#
.#####.....#....#.....#.#.....#..#####..#.....#.###.#`

type Banner struct {
	renderer    *gk.Renderer
	bounds      gk.Rect
	pulseTimer  float64
	pulseSpeed  float64
	baseColor   sdl.Color
	shadowColor sdl.Color
	cellSize    int32
	cellGap     int32
	offsetX     int32
	offsetY     int32
	lines       []string
}

func NewBanner(renderer *gk.Renderer, bounds gk.Rect) *Banner {
	const (
		cellSize = 14
		cellGap  = 4
	)

	lines := strings.Split(strings.TrimSpace(text), "\n")

	return &Banner{
		renderer:    renderer,
		bounds:      bounds,
		pulseTimer:  0,
		pulseSpeed:  15.0,
		baseColor:   sdl.Color{R: 220, G: 180, B: 50, A: 255},
		shadowColor: sdl.Color{R: 120, G: 60, B: 20, A: 255},
		cellSize:    cellSize,
		cellGap:     cellGap,
		offsetX:     (bounds.W - int32(len(lines[0]))*(cellSize+cellGap)) / 2,
		offsetY:     150,
		lines:       lines,
	}
}

func (b *Banner) Update(delta uint64) {
	deltaSeconds := float64(delta) / 1000.0

	b.pulseTimer += deltaSeconds * b.pulseSpeed
}

func (b *Banner) Render() {
	pulseValue := (math.Sin(b.pulseTimer) + 1) / 2

	red := gk.MinU8(255, float64(b.baseColor.R)+pulseValue*35)
	green := gk.MaxU8(100, float64(b.baseColor.G)-pulseValue*40)
	blue := gk.MaxU8(20, float64(b.baseColor.B)-pulseValue*30)

	color := sdl.Color{R: red, G: green, B: blue, A: 255}
	size := gk.NewSize(b.cellSize, b.cellSize)

	cellTex := b.renderer.CreateRectTexture(color, size)
	defer cellTex.Destroy()

	shadowTex := b.renderer.CreateRectTexture(b.shadowColor, size)
	defer shadowTex.Destroy()

	for y, line := range b.lines {
		for x, char := range line {
			if char == '#' {
				posX := b.offsetX + int32(x)*(b.cellSize+b.cellGap)
				posY := b.offsetY + int32(y)*(b.cellSize+b.cellGap)

				b.renderer.Copy(shadowTex, gk.NewRect(
					posX+2,
					posY+2,
					b.cellSize,
					b.cellSize,
				))

				b.renderer.Copy(cellTex, gk.NewRect(
					posX,
					posY,
					b.cellSize,
					b.cellSize,
				))
			}
		}
	}
}

func (b *Banner) Destroy() {
}
