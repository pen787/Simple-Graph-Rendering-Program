package main

import (
	"container/list"
	"fmt"
	"image/color"
	"log"
	"os"
	"penrenderingmethod/components"
	"penrenderingmethod/mappers"
	"penrenderingmethod/systems"
	"penrenderingmethod/utility"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

var (
	Zoom         float32    = 5
	CameraOrigin rl.Vector2 = rl.NewVector2(0, 0)
)

func main() {

	programArgs := os.Args[1:]
	if len(programArgs) < 1 {
		fmt.Printf("Usage : [path..]\n")
		return
	}

	rl.InitWindow(utility.StartWindowWidth, utility.StartWindowHeight, "rendering window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	world := ecs.NewWorld()
	systemList := list.New()
	systemList.PushBack(systems.NewRendering(&world, &CameraOrigin, &Zoom))

	graphMap := mappers.NewGraphMapper(&world)

	for _, strpath := range programArgs {
		graphMap.NewEntity(components.NewEmbledScript(), &color.RGBA{}, &strpath)
	}

	for e := systemList.Front(); e != nil; e = e.Next() {
		v, complete := e.Value.(systems.BaseSystem)
		if !complete {
			log.Fatalln("Something not right in render!")
		}
		v.Init()
	}

	for !rl.WindowShouldClose() {
		//// [Update]
		for e := systemList.Front(); e != nil; e = e.Next() {
			v, complete := e.Value.(systems.BaseSystem)
			if !complete {
				log.Fatalln("Something not right in update!")
			}
			v.Update()
		}

		//// [Draw]
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Graph
		utility.DrawGrid(CameraOrigin, Zoom)

		// render system
		for e := systemList.Front(); e != nil; e = e.Next() {
			v, complete := e.Value.(systems.BaseSystem)
			if !complete {
				log.Fatalln("Something not right in render!")
			}
			v.Render()
		}

		// UI
		gui.WindowBox(rl.Rectangle{X: 0, Y: 30, Width: 200, Height: 70}, "#44# Camera Control")
		zoom_slider := gui.Slider(rl.Rectangle{X: 50, Y: 70, Width: 100, Height: 20}, "Zoom", "", Zoom, 0.1, 15)
		Zoom = zoom_slider

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	for e := systemList.Front(); e != nil; e = e.Next() {
		v, complete := e.Value.(systems.BaseSystem)
		if !complete {
			log.Fatalln("Something not right in render!")
		}
		v.Done()
	}
}
