package gk

import (
	"image/png"
	"math"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Flip sdl.RendererFlip

const (
	FlipNone       = Flip(sdl.FLIP_NONE)
	FlipHorizontal = Flip(sdl.FLIP_HORIZONTAL)
)

// Make linter happy.
var _ = FlipHorizontal

type Renderer struct {
	peer *sdl.Renderer
}

func newRenderer(peer *sdl.Renderer) *Renderer {
	return &Renderer{peer}
}

func (r *Renderer) Destroy() {
	panicErr(r.peer.Destroy())
}

func (r *Renderer) Clear() {
	panicErr(r.peer.Clear())
}

func (r *Renderer) Present() {
	r.peer.Present()
}

func (r *Renderer) Copy(texture *Texture, dst Rect) {
	panicErr(r.peer.Copy(texture.peer, nil, dst.peer()))
}

func (r *Renderer) CopyDirect(texture *Texture, pos Pos) {
	panicErr(r.peer.Copy(texture.peer, nil, &sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: texture.W,
		H: texture.H,
	}))
}

func (r *Renderer) CopyDirectF(texture *Texture, pos FPos) {
	panicErr(r.peer.CopyF(texture.peer, nil, &sdl.FRect{
		X: pos.X,
		Y: pos.Y,
		W: float32(texture.W),
		H: float32(texture.H),
	}))
}

func (r *Renderer) CopyScaled(texture *Texture, src, dst Rect) {
	panicErr(r.peer.Copy(texture.peer, src.peer(), dst.peer()))
}

func (r *Renderer) CopyRotate(texture *Texture, dst Rect, angle int32, flip Flip) {
	panicErr(r.peer.CopyEx(texture.peer, nil, dst.peer(), float64(angle), nil, sdl.RendererFlip(flip)))
}

func (r *Renderer) CopyRotateF(texture *Texture, dst FRect, angle float64, flip Flip) {
	panicErr(r.peer.CopyExF(texture.peer, nil, dst.peer(), angle, nil, sdl.RendererFlip(flip)))
}

func (r *Renderer) CreateTextureFromSurface(surface *Surface) *Texture {
	texture, err := r.peer.CreateTextureFromSurface(surface.peer)

	panicErr(err)

	return newTexture(texture)
}

func (r *Renderer) CreateSinglePixelTexture(color sdl.Color) *Texture {
	texture, err := r.peer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)

	panicErr(err)

	pixels := make([]byte, 4)

	pixels[0] = color.R
	pixels[1] = color.G
	pixels[2] = color.B
	pixels[3] = color.A

	panicErr(texture.SetBlendMode(sdl.BLENDMODE_BLEND))
	panicErr(texture.Update(nil, unsafe.Pointer(&pixels[0]), 4))

	return newTexture(texture)
}

func (r *Renderer) CreateHorizontalLineTexture(color sdl.Color, size Size) *Texture {
	surface := NewSurface(size.W, size.H)

	for y := int32(0); y < size.H; y++ {
		surface.HorLine(0, size.W, y, color)
	}

	return r.CreateTextureFromSurface(surface)
}

func (r *Renderer) CreateVerticalLineTexture(color sdl.Color, size Size) *Texture {
	surface := NewSurface(size.W, size.H)

	for x := int32(0); x < size.W; x++ {
		surface.VerLine(0, size.H, x, color)
	}

	return r.CreateTextureFromSurface(surface)
}

func (r *Renderer) CreateRectTexture(color sdl.Color, size Size) *Texture {
	surface := NewSurface(size.W, size.H)

	surface.Fill(color)

	return r.CreateTextureFromSurface(surface)
}

func (r *Renderer) LoadImageTexture(name string) *Texture {
	f, err := os.Open(imagesDir + filepath.Clean(name) + ".png")
	if err != nil {
		panic(err)
	}

	defer func() {
		panicErr(f.Close())
	}()

	img, err := png.Decode(f)

	panicErr(err)

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	pixels := make([]byte, w*h*4)
	bIndex := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			pixels[bIndex] = byte(r / 256)
			bIndex++

			pixels[bIndex] = byte(g / 256)
			bIndex++

			pixels[bIndex] = byte(b / 256)
			bIndex++

			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}

	tex := r.pixelsToTexture(pixels, int32(w), int32(h))

	panicErr(tex.SetBlendMode(sdl.BLENDMODE_BLEND))

	return newTexture(tex)
}

// RotateTexture rotates the given source texture by angle degrees and returns a new texture.
func (r *Renderer) RotateTexture(src *Texture, angle float64) *Texture {
	// compute bounding box of rotated rectangle
	w := float64(src.W)
	h := float64(src.H)
	rad := angle * math.Pi / 180.0

	newW := int32(math.Max(1, math.Ceil(math.Abs(w*math.Cos(rad))+math.Abs(h*math.Sin(rad)))))
	newH := int32(math.Max(1, math.Ceil(math.Abs(w*math.Sin(rad))+math.Abs(h*math.Cos(rad)))))

	target, err := r.peer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_TARGET, newW, newH)
	panicErr(err)

	// ensure alpha blending on the resulting texture
	panicErr(target.SetBlendMode(sdl.BLENDMODE_BLEND))

	// set render target to the new texture
	panicErr(r.peer.SetRenderTarget(target))

	// clear the target to fully transparent
	cr, cg, cb, ca, _ := r.peer.GetDrawColor()

	panicErr(r.peer.SetDrawColor(0, 0, 0, 0))
	panicErr(r.peer.Clear())

	// center the src texture within the target and rotate around the center
	dst := sdl.Rect{
		X: (newW - src.W) / 2,
		Y: (newH - src.H) / 2,
		W: src.W,
		H: src.H,
	}

	panicErr(r.peer.CopyEx(src.peer, nil, &dst, angle, nil, sdl.FLIP_NONE))

	// restore previous renderer state
	panicErr(r.peer.SetDrawColor(cr, cg, cb, ca))
	panicErr(r.peer.SetRenderTarget(nil))

	return newTexture(target)
}

func (r *Renderer) LoadImageTextureScaled(name string, scale float32) *Texture {
	f, err := os.Open(imagesDir + filepath.Clean(name) + ".png")
	if err != nil {
		panic(err)
	}

	defer func() {
		panicErr(f.Close())
	}()

	img, err := png.Decode(f)

	panicErr(err)

	srcW := img.Bounds().Max.X
	srcH := img.Bounds().Max.Y

	dstW := int(math.Max(1, math.Round(float64(float32(srcW)*scale))))
	dstH := int(math.Max(1, math.Round(float64(float32(srcH)*scale))))

	pixels := make([]byte, dstW*dstH*4)
	bIndex := 0

	for y := 0; y < dstH; y++ {
		srcY := int(float32(y) / scale)

		if srcY >= srcH {
			srcY = srcH - 1
		}

		for x := 0; x < dstW; x++ {
			srcX := int(float32(x) / scale)
			if srcX >= srcW {
				srcX = srcW - 1
			}

			r8, g8, b8, a8 := img.At(srcX, srcY).RGBA()

			pixels[bIndex] = byte(r8 / 256)
			bIndex++
			pixels[bIndex] = byte(g8 / 256)
			bIndex++
			pixels[bIndex] = byte(b8 / 256)
			bIndex++
			pixels[bIndex] = byte(a8 / 256)
			bIndex++
		}
	}

	tex := r.pixelsToTexture(pixels, int32(dstW), int32(dstH))

	panicErr(tex.SetBlendMode(sdl.BLENDMODE_BLEND))

	return newTexture(tex)
}

func (r *Renderer) pixelsToTexture(pixels []byte, w, h int32) *sdl.Texture {
	tex, err := r.peer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, w, h)
	if err != nil {
		panic(err)
	}

	if err = tex.Update(nil, unsafe.Pointer(&pixels[0]), int(w)*4); err != nil {
		panic(err)
	}

	return tex
}
