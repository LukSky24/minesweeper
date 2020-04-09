package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Coords structure stores coordinates of Cell in Grid
type Coords struct {
	X, Y int
}

type Grid struct {
	cols, rows, bombs int
	cells             []*Cell
}

// CreateGrid constructs unrevealed game grid with bombs planted
func CreateGrid(cols, rows, bombs int) (*Grid, error) {
	if bombs >= cols*rows {
		return nil, errors.New("Bombs count must not be greater or equal to cells count")
	}

	g := Grid{}
	g.cols = cols
	g.rows = rows
	g.bombs = bombs

	for i := 0; i < g.cols*g.rows; i++ {
		g.cells = append(g.cells, &Cell{false, false, false, 0})
	}
	g.plantBombs()

	return &g, nil
}

func (g *Grid) getCell(coords Coords) (*Cell, error) {
	if coords.X > g.cols-1 || coords.X < 0 || coords.Y > g.rows-1 || coords.Y < 0 {
		return nil, errors.New("Trying to reach fer cell out of range")
	}

	return g.cells[coordsToIndex(coords.X, coords.Y, g.cols)], nil
}

func (g *Grid) getCellNeighbours(coords Coords) (neighbours map[Coords]*Cell) {
	neighbours = make(map[Coords]*Cell)

	for c := -1; c <= 1; c++ {
		for r := -1; r <= 1; r++ {
			if c == 0 && r == 0 {
				continue
			}

			if coords.X+c < 0 || coords.X+c >= g.cols ||
				coords.Y+r < 0 || coords.Y+r >= g.rows {
				continue
			}

			cell, err := g.getCell(Coords{coords.X + c, coords.Y + r})
			if err == nil {
				neighbours[Coords{coords.X + c, coords.Y + r}] = cell
			}
		}
	}

	return neighbours
}

func (g *Grid) RevealOn(x, y int) {
	c, err := g.getCell(Coords{x, y})
	if err != nil {
		return
	}

	c.marked = false

	if c.bomb {
		g.revealAll()
		return
	}

	bombCount := 0
	n := g.getCellNeighbours(Coords{x, y})
	for _, c := range n {
		if c.bomb {
			bombCount++
		}
	}
	c.revealed = true
	if bombCount == 0 {
		for i, nc := range n {
			if !nc.revealed {
				_ = i
				// g.RevealOn(indexToCoords(i, g.cols))
			}
		}
	} else {
		c.count = bombCount
	}
}

func (g *Grid) ToggleMarkOn(x, y int) {
	// c := g.getCell(x, y)
	// if c.revealed {
	// 	return
	// }

	// c.marked = !c.marked
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
		x, y := indexToCoords(i, g.cols)
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

func coordsToIndex(x, y, cols int) int {
	return y*cols + x
}

func indexToCoords(i, cols int) (x, y int) {
	y = i / cols
	x = i - (y * cols)

	return x, y
}
