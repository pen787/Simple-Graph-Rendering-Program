package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const StartWindowWidth, StartWindowHeight = 1000., 600.

var (
	Zoom         float32    = 5
	CameraOrigin rl.Vector2 = rl.NewVector2(0, 0)
)

func main() {
	var (
		IsScriptError      bool     = false
		GraphResolution    int32    = 2
		GraphResolutionMin int32    = 1
		GraphResolutionMax int32    = 20
		GraphColor         rl.Color = color.RGBA{}
	)

	// UI varible
	var (
		ResolutionStepValue       int32        = 1
		ResolutionStepBoundingBox rl.Rectangle = rl.Rectangle{X: 100, Y: 220, Width: 70, Height: 30}
		ResolutionStepFocus       bool         = false
	)

	programArgs := os.Args[1:]
	if len(programArgs) < 1 {
		fmt.Printf("Usage : [path]\n")
		return
	}

	rl.InitWindow(StartWindowWidth, StartWindowHeight, "rendering window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Init a embled script
	embledScriptService := NewEmbledScript()
	defer embledScriptService.Close()

	if err := embledScriptService.DoFile(programArgs[0]); err != nil {
		log.Println("Error when script load : ")
		fmt.Println(err)
		IsScriptError = true
	}

	c, err := embledScriptService.CallLoad()
	if err != nil {
		log.Println("Script Load function error : ")
		fmt.Println(err)
		IsScriptError = true
	}
	GraphColor = c

	for !rl.WindowShouldClose() {
		//// [Update]
		// UI
		mousePos := rl.GetMousePosition()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if rl.CheckCollisionPointRec(mousePos, ResolutionStepBoundingBox) {
				ResolutionStepFocus = true
			} else {
				ResolutionStepFocus = false
			}
		}

		//// [Draw]
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Graph
		DrawGrid()

		if !IsScriptError {
			DrawGraph(int(GraphResolution), func(x float32) float32 {
				value, err := embledScriptService.CallRender(x)
				if err != nil {
					IsScriptError = true
					log.Println("Script Render function error : ")
					fmt.Print(err)
				}

				return value
			}, GraphColor)
		}

		// UI
		gui.WindowBox(rl.Rectangle{X: 0, Y: 30, Width: 200, Height: 70}, "#44# Camera Control")
		zoom_slider := gui.Slider(rl.Rectangle{X: 50, Y: 70, Width: 100, Height: 20}, "Zoom", "", Zoom, 0.1, 15)
		Zoom = zoom_slider

		gui.WindowBox(rl.Rectangle{X: 0, Y: 120, Width: 200, Height: 300}, "#63# Graph")
		gui.Label(rl.Rectangle{X: 5, Y: 160, Width: 170, Height: 30}, fmt.Sprintf("Script : %s", programArgs[0]))

		click := gui.Button(rl.Rectangle{X: 50, Y: 185, Width: 100, Height: 30}, "Reset script")
		if click {
			embledScriptService.ResetScript()
			c, err := embledScriptService.CallLoad()
			if err != nil {
				log.Println("Script Load function error : ")
				fmt.Println(err)
				IsScriptError = true
			}
			GraphColor = c
		}

		valueBoxUpdate := gui.ValueBox(
			ResolutionStepBoundingBox,
			"Resolution Step :",
			&ResolutionStepValue,
			int(GraphResolutionMin),
			int(GraphResolutionMax),
			ResolutionStepFocus,
		)
		if valueBoxUpdate {
			GraphResolution = max(min(ResolutionStepValue, GraphResolutionMax), GraphResolutionMin)
		}

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
