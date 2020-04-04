package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Cell struct {
	revealed, bomb bool
}

func (c *Cell) Draw(r *sdl.Renderer, f *ttf.Font, vp *sdl.Rect) {
	r.DrawRect(vp)
	if c.revealed && c.bomb {
		c.drawBomb(r, f, vp)
	}
}

func (c *Cell) drawBomb(r *sdl.Renderer, f *ttf.Font, vp *sdl.Rect) {
	surface, err := f.RenderUTF8Solid("B", sdl.Color{255, 0, 0, 255})
	if err != nil {
		panic(err)
	}

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	r.Copy(texture, nil, vp)
}
