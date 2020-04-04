package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	fmt.Println("minesweeper by eczek.")

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Mnesweeper", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 300, 300, sdl.WINDOW_OPENGL)
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

	font, err := ttf.OpenFont("Hack-Regular.ttf", 24)
	if err != nil {
		panic(err)
	}

	vp := sdl.Rect{0, 0, 300, 300}
	g := CreateGrid(10, 10)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("exit")
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 255)
		g.Draw(renderer, font, vp)

		renderer.Present()
	}
}
