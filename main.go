package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	fmt.Println("minesweeper by eczek.")

	const WINGDOW_WIDTH = 800
	const WINDOW_HEIGHT = 600

	const COLS = 30
	const ROWS = 30
	const BOMBS = 10

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Minesweeper", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, WINGDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	err = ttf.Init()
	if err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("OpenSans-Regular.ttf", 24)
	if err != nil {
		panic(err)
	}

	vp := sdl.Rect{0, 0, WINGDOW_WIDTH, WINDOW_HEIGHT}
	g, _ := CreateGrid(COLS, ROWS, BOMBS)

	for {
		sdl.Delay(100)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch ev := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("exit")
				return
			case *sdl.MouseButtonEvent:
				if ev.GetType() == sdl.MOUSEBUTTONUP {
					coords := getCellCoordsPosFromMouseCoords(
						ev.X, ev.Y, WINGDOW_WIDTH, WINDOW_HEIGHT, COLS, ROWS)

					switch ev.Button {
					case sdl.BUTTON_LEFT:
						g.RevealOn(coords)
					case sdl.BUTTON_RIGHT:
						g.ToggleMarkOn(coords)
					case sdl.BUTTON_MIDDLE:
						g.Reset()
					}

				}
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 255)
		g.Draw(renderer, font, vp)

		renderer.Present()
	}
}

func getCellCoordsPosFromMouseCoords(x, y int32, w, h int, c, r int) Coords {
	X := int(x) / (w / c)
	Y := int(y) / (h / r)

	return Coords{X, Y}
}
