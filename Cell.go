package main

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Cell struct {
	revealed, bomb bool
	count          int
}

func (c *Cell) Draw(r *sdl.Renderer, f *ttf.Font, vp *sdl.Rect) {
	r.DrawRect(vp)

	if c.revealed {
		r.SetDrawColor(100, 100, 100, 255)
		r.FillRect(&sdl.Rect{vp.X + 1, vp.Y + 1, vp.W - 2, vp.H - 2})
		r.SetDrawColor(0, 0, 0, 255)

		if c.count > 0 && !c.bomb {
			c.drawString(strconv.Itoa(c.count), r, f, vp, sdl.Color{0, 0, 255, 255})
		}
	}

	if c.revealed && c.bomb {
		c.drawBomb(r, f, vp)
	}
}

func (c *Cell) drawBomb(r *sdl.Renderer, f *ttf.Font, vp *sdl.Rect) {
	c.drawString("B", r, f, vp, sdl.Color{255, 0, 0, 255})
}

func (c *Cell) drawString(letter string, r *sdl.Renderer, f *ttf.Font, vp *sdl.Rect, col sdl.Color) {
	surface, err := f.RenderUTF8Solid(letter, col)
	if err != nil {
		panic(err)
	}

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	r.Copy(texture, nil, vp)
}
