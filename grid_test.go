package main

import "testing"

func TestCreateGrid(t *testing.T) {
	var revealedCount, cellsCount int

	g, _ := CreateGrid(10, 10, 10)
	for _, c := range g.cells {
		cellsCount++
		if c.revealed {
			revealedCount++
		}
	}

	if cellsCount != 100 {
		t.Errorf("Grid 10x10 should contain 100 cells, got %d", cellsCount)
	}

	if revealedCount > 0 {
		t.Errorf("New grid should not contain revealed cells, want 0 revealed got %d", revealedCount)
	}
}

func TestBombsCount(t *testing.T) {
	g, _ := CreateGrid(10, 10, 10)
	var count int
	for _, c := range g.cells {
		if c.bomb {
			count++
		}
	}

	if count != 10 {
		t.Errorf("Wrong bombs count, want 10 got %d", count)
	}
}

func TestCantCreateGridWithTooManyBombs(t *testing.T) {
	g, err := CreateGrid(5, 5, 25)
	if g != nil {
		t.Errorf("Should not allow to create grid with bombs count greater than or equal to cells count")
	}

	if err == nil {
		t.Errorf("Should return error when trying to create grid with bombs count greater or equal to cells count")
	}
}

func TestCantReachNotExistingCell(t *testing.T) {
	var coords = [6][2]int{
		{10, 0},
		{0, 10},
		{10, 10},
		{-1, 0},
		{0, -1},
		{-1, -1}}
	g, _ := CreateGrid(10, 10, 10)

	for _, coord := range coords {
		c, err := g.getCell(coord[0], coord[1])
		if c != nil {
			t.Errorf("Should not allow to reach for cell out of range")
		}

		if err == nil {
			t.Errorf("Should return error when trying to reach for cell out of range")
		}
	}
}

func TestCoordsToIndex(t *testing.T) {
	var tests = map[int][4]int{
		0:  {0, 0, 1, 1},
		1:  {0, 1, 1, 2},
		2:  {0, 1, 2, 2},
		24: {4, 4, 5, 5}}

	for i, args := range tests {
		got := coordsToIndex(args[0], args[1], args[2], args[3])
		if got != i {
			t.Errorf("cordsToIndex(%d, %d, %d, %d) = %d; want %d",
				args[0], args[1], args[2], args[3], got, i)
		}
	}
}

func TestIndexToCoords(t *testing.T) {
	var tests = map[int][4]int{
		0:  {0, 0, 1, 1},
		1:  {0, 1, 1, 2},
		2:  {0, 1, 2, 2},
		24: {4, 4, 5, 5}}

	for i, args := range tests {
		gotX, gotY := indexToCoords(i, args[2], args[3])
		if gotX != args[0] || gotY != args[1] {
			t.Errorf("indexToCoords(%d, %d, %d) = %d, %d; want %d, %d",
				i, args[2], args[3], gotX, gotY, args[0], args[1])
		}
	}
}
