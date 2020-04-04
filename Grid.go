package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Grid struct {
	cols, rows int
	cells      []Cell
}

func CreateGrid(cols, rows int) *Grid {
	g := Grid{}
	g.cols = cols
	g.rows = rows

	for i := 0; i < cols*rows; i++ {
		g.cells = append(g.cells, Cell{})
	}

	return &g
}

func (g *Grid) Draw(r *sdl.Renderer, f *ttf.Font, vp sdl.Rect) {
	r.SetViewport(&vp)
	cellW := r.GetViewport().W / int32(g.cols)
	cellH := r.GetViewport().H / int32(g.rows)

	r.SetDrawColor(0, 0, 0, 255)

	for i, c := range g.cells {
		y := i / g.rows
		x := i - (y * g.cols)

		cellVp := sdl.Rect{int32(x) * cellW, int32(y) * cellH, cellW, cellH}
		c.Draw(r, f, &cellVp)
	}
}
