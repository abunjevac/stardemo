package gk

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	peer *sdl.Surface
	W    int32
	H    int32
}

func NewSurface(width, height int32) *Surface {
	peer, err := sdl.CreateRGBSurfaceWithFormat(
		0,
		width,
		height,
		32,
		sdl.PIXELFORMAT_ABGR8888)

	panicErr(err)

	return newSurface(peer)
}

func newSurface(peer *sdl.Surface) *Surface {
	return &Surface{
		peer: peer,
		W:    peer.W,
		H:    peer.H,
	}
}

func (s *Surface) Destroy() {
	s.peer.Free()
}

func (s *Surface) Fill(color sdl.Color) {
	panicErr(s.peer.FillRect(nil, s.mapColor(color)))
}

func (s *Surface) FillRect(rect Rect, color sdl.Color) {
	panicErr(s.peer.FillRect(rect.peer(), s.mapColor(color)))
}

func (s *Surface) Blit(dst *Surface, dstRect Rect) {
	panicErr(s.peer.Blit(nil, dst.peer, dstRect.peer()))
}

func (s *Surface) HorLine(from, to, y int32, c sdl.Color) {
	for i := from; i < to; i++ {
		s.peer.Set(int(i), int(y), color.RGBA{
			R: c.A,
			G: c.B,
			B: c.G,
			A: c.R,
		})
	}
}

func (s *Surface) VerLine(from, to, x int32, c sdl.Color) {
	for i := from; i < to; i++ {
		s.peer.Set(int(x), int(i), color.RGBA{
			R: c.A,
			G: c.B,
			B: c.G,
			A: c.R,
		})
	}
}

func (s *Surface) DrawTransparentBackground() {
	w := int32(4)
	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	gray := sdl.Color{R: 192, G: 192, B: 192, A: 255}

	for y := 0; y < int(s.H/w); y++ {
		for x := 0; x < int(s.W/w/2); x++ {
			s.FillRect(Rect{
				X: int32(x)*w*2 + w*int32(y%2),
				Y: int32(y) * w,
				W: w,
				H: w,
			}, white)

			s.FillRect(Rect{
				X: int32(x)*w*2 + w*int32((y+1)%2),
				Y: int32(y) * w,
				W: w,
				H: w,
			}, gray)
		}
	}
}

func (s *Surface) mapColor(color sdl.Color) uint32 {
	return sdl.MapRGBA(s.peer.Format, color.R, color.G, color.B, color.A)
}
