package main

import "testing"

func TestCreateGrid(t *testing.T) {
	g, _ := CreateGrid(10, 10, 10)

	var count int
	for _, c := range g.cells {
		if c.revealed {
			count++
		}
	}

	if count > 0 {
		t.Errorf("New grid should not contain revealed cells, want 0 revealed got %d", count)
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
