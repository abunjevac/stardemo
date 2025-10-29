package gk

type ImageCarousel struct {
	images     []*Texture
	continuous bool
	index      int
	timer      uint64
	step       uint64
	done       bool
}

func NewImageCarousel(images []*Texture, duration uint64, continuous bool) *ImageCarousel {
	return &ImageCarousel{
		images:     images,
		continuous: continuous,
		index:      0,
		timer:      0,
		step:       duration / uint64(len(images)),
	}
}

func (c *ImageCarousel) Update(delta uint64) {
	if c.done {
		return
	}

	c.timer += delta

	if c.timer >= c.step {
		c.timer = 0
		c.index++

		if c.index+1 >= len(c.images) {
			if c.continuous {
				c.index = 0
			} else {
				c.done = true
			}
		}
	}
}

func (c *ImageCarousel) Done() bool {
	return c.done
}

func (c *ImageCarousel) Previous() *Texture {
	return c.images[max(0, c.index-1)]
}

func (c *ImageCarousel) Current() *Texture {
	return c.images[c.index]
}
