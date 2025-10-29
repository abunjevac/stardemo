package gk

import "fmt"

type ImageCache struct {
	renderer *Renderer
	images   map[string]*Texture
}

func NewImageCache(r *Renderer) *ImageCache {
	return &ImageCache{
		renderer: r,
		images:   make(map[string]*Texture),
	}
}

func (c *ImageCache) Get(name string) *Texture {
	if image, found := c.images[name]; found {
		return image
	}

	image := c.renderer.LoadImageTexture(name)

	c.images[name] = image

	return image
}

func (c *ImageCache) GetScaled(name string, scale float32) *Texture {
	key := fmt.Sprintf("%s@s%.3f", name, scale)

	if image, found := c.images[key]; found {
		return image
	}

	image := c.renderer.LoadImageTextureScaled(name, scale)

	c.images[key] = image

	return image
}

func (c *ImageCache) GetRotated(name string, angle float64) *Texture {
	key := fmt.Sprintf("%s@r%.3f", name, angle)

	if image, found := c.images[key]; found {
		return image
	}

	image := c.renderer.RotateTexture(c.Get(name), angle)

	c.images[key] = image

	return image
}

func (c *ImageCache) Destroy() {
	for _, image := range c.images {
		image.Destroy()
	}

	clear(c.images)
}
