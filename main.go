package main

import (
	"fmt"
	"os"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const StartWindowWidth, StartWindowHeight = 800., 450.

var (
	Zoom         float32    = 5
	CameraOrigin rl.Vector2 = rl.NewVector2(0, 0)
)

func DrawGrid() {
	// X line
	rl.DrawLine(0, (StartWindowHeight/2)+int32(CameraOrigin.Y*Zoom), StartWindowWidth, StartWindowHeight/2+int32(CameraOrigin.Y*Zoom), rl.Red)

	// Y line
	rl.DrawLine(StartWindowWidth/2-int32(CameraOrigin.X*Zoom), 0, StartWindowWidth/2-int32(CameraOrigin.X*Zoom), StartWindowHeight, rl.Green)
}

func DrawGraph(resolution int, fun func(x float32) float32, color rl.Color) {
	resolution_step := resolution
	points := make([]rl.Vector2, 0, rl.GetScreenWidth()/resolution_step)
	for x := 0; x < rl.GetScreenWidth(); x += resolution_step {
		xgpos := ScreenToGraphCords(rl.NewVector2(float32(x), 0))

		renderpos := GraphtoScreenCords(rl.NewVector2(xgpos.X, fun(xgpos.X)))
		points = append(points, renderpos)

	}
	lastPoint := rl.NewVector2(0, 0)
	for i, point := range points {
		if i == 0 {
			lastPoint = point
			continue
		}

		rl.DrawLineV(lastPoint, point, color)
		lastPoint = point
	}
}

func main() {
	programArgs := os.Args[1:]
	if len(programArgs) < 1 {
		fmt.Printf("Usage : [path]\n")
		return
	}

	rl.InitWindow(StartWindowWidth, StartWindowHeight, "rendering window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Init a embled script
	NewEmbledScript(programArgs[0])

	for !rl.WindowShouldClose() {
		// Update

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		DrawGrid()

		DrawGraph(2, func(x float32) float32 {
			return x * x
		}, rl.DarkGreen)

		// UI

		gui.WindowBox(rl.Rectangle{X: 0, Y: 30, Width: 200, Height: 70}, "#44# Camera Control")
		zoom_slider := gui.Slider(rl.Rectangle{X: 50, Y: 70, Width: 100, Height: 20}, "Zoom", "", Zoom, 0.1, 15)
		Zoom = zoom_slider

		gui.WindowBox(rl.Rectangle{X: 0, Y: 120, Width: 200, Height: 300}, "#63# Graph")
		gui.Label(rl.Rectangle{X: 5, Y: 150, Width: 190, Height: 250}, "Hello")

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
