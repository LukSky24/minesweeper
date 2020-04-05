package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Grid struct {
	cols, rows, bombs int
	cells             []*Cell
}

func CreateGrid(cols, rows, bombs int) *Grid {
	g := Grid{}
	g.cols = cols
	g.rows = rows
	g.bombs = bombs

	for i := 0; i < g.cols*g.rows; i++ {
		g.cells = append(g.cells, &Cell{false, false, false, 0})
	}
	g.plantBombs()

	return &g
}

func (g *Grid) getCell(x, y int) *Cell {
	return g.cells[coordsToIndex(x, y, g.cols, g.rows)]
}

func (g *Grid) getCellNeighbours(x, y int) (neighbours map[int]*Cell) {
	neighbours = make(map[int]*Cell)

	for c := -1; c <= 1; c++ {
		for r := -1; r <= 1; r++ {
			if c == 0 && r == 0 {
				continue
			}

			if x+c < 0 || x+c >= g.cols || y+r < 0 || y+r >= g.rows {
				continue
			}

			neighbours[coordsToIndex(x+c, y+r, g.cols, g.rows)] = g.getCell(x+c, y+r)
		}
	}

	return neighbours
}

func (g *Grid) RevealOn(x, y int) {
	c := g.getCell(x, y)
	c.marked = false

	if c.bomb {
		g.revealAll()
		return
	}

	bombCount := 0
	n := g.getCellNeighbours(x, y)
	for _, c := range n {
		if c.bomb {
			bombCount++
		}
	}
	c.revealed = true
	if bombCount == 0 {
		for i, nc := range n {
			if !nc.revealed {
				g.RevealOn(indexToCoords(i, g.cols, g.rows))
			}
		}
	} else {
		c.count = bombCount
	}
}

func (g *Grid) ToggleMarkOn(x, y int) {
	c := g.getCell(x, y)
	if c.revealed {
		return
	}

	c.marked = !c.marked
}

func (g *Grid) revealAll() {
	for _, c := range g.cells {
		c.revealed = true
		c.marked = false
	}
}

func (g *Grid) Reset() {
	for i := 0; i < g.cols*g.rows; i++ {
		g.cells[i].bomb = false
		g.cells[i].marked = false
		g.cells[i].revealed = false
		g.cells[i].count = 0
	}

	g.plantBombs()
}

func (g *Grid) plantBombs() {
	rand.Seed(time.Now().Unix())
	for i, ci := range rand.Perm(g.cols * g.rows) {
		if i >= g.bombs {
			break
		}
		g.cells[ci].bomb = true
	}
}

func (g *Grid) Draw(r *sdl.Renderer, f *ttf.Font, vp sdl.Rect) {
	r.SetViewport(&vp)
	cellW := vp.W / int32(g.cols)
	cellH := vp.H / int32(g.rows)

	r.SetDrawColor(0, 0, 0, 255)

	for i, c := range g.cells {
		x, y := indexToCoords(i, g.cols, g.rows)
		cellVp := sdl.Rect{int32(x) * cellW, int32(y) * cellH, cellW, cellH}

		c.Draw(r, f, &cellVp)
	}

	if g.victory() {
		fmt.Println("victory!")
	}
}

func (g *Grid) victory() bool {
	revealed := 0
	for _, c := range g.cells {
		if c.revealed {
			revealed++
		}
	}
	return g.cols*g.rows-revealed == g.bombs
}

func coordsToIndex(x, y, c, r int) int {
	return y*c + x
}

func indexToCoords(i, c, r int) (x, y int) {
	y = i / c
	x = i - (y * c)

	return x, y
}
