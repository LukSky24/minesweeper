package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Grid struct {
	cols, rows int
	cells      []Cell
}

func CreateGrid(cols, rows, bombs int) *Grid {
	g := Grid{}
	g.cols = cols
	g.rows = rows

	for i := 0; i < cols*rows; i++ {
		g.cells = append(g.cells, Cell{false, false})
	}

	bx := rand.Perm(bombs)
	by := rand.Perm(bombs)
	for i := 0; i < bombs; i++ {
		g.getCell(bx[i], by[i]).bomb = true
	}

	return &g
}

func (g *Grid) getCell(x, y int) *Cell {
	return &g.cells[coordsToIndex(x, y, g.cols, g.rows)]
}

func (g *Grid) RevealOn(x, y int) {
	g.getCell(x, y).revealed = true
}

func (g *Grid) Draw(r *sdl.Renderer, f *ttf.Font, vp sdl.Rect) {
	r.SetViewport(&vp)
	cellW := r.GetViewport().W / int32(g.cols)
	cellH := r.GetViewport().H / int32(g.rows)

	r.SetDrawColor(0, 0, 0, 255)

	for i, c := range g.cells {
		x, y := indexToCoords(i, g.cols, g.rows)
		cellVp := sdl.Rect{int32(x) * cellW, int32(y) * cellH, cellW, cellH}

		c.Draw(r, f, &cellVp)
	}
}

func coordsToIndex(x, y, c, r int) int {
	return y*c + x
}

func indexToCoords(i, c, r int) (x, y int) {
	y = i / r
	x = i - (y * r)

	return x, y
}
