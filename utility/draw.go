package utility

import rl "github.com/gen2brain/raylib-go/raylib"

const StartWindowWidth, StartWindowHeight = 1000., 600.

func DrawGrid(CameraOrigin rl.Vector2, Zoom float32) {
	// X line
	rl.DrawLine(0, (StartWindowHeight/2)+int32(CameraOrigin.Y*Zoom), StartWindowWidth, StartWindowHeight/2+int32(CameraOrigin.Y*Zoom), rl.Red)

	// Y line
	rl.DrawLine(StartWindowWidth/2-int32(CameraOrigin.X*Zoom), 0, StartWindowWidth/2-int32(CameraOrigin.X*Zoom), StartWindowHeight, rl.Green)
}

func DrawGraph(resolution int, fun func(x float32) float32, color rl.Color, CameraOrigin rl.Vector2, zoom float32) {
	resolution_step := max(resolution, 1)
	points := make([]rl.Vector2, 0, rl.GetScreenWidth()/resolution_step)
	for x := 0; x < rl.GetScreenWidth(); x += resolution_step {
		xgpos := ScreenToGraphCords(rl.NewVector2(float32(x), 0), CameraOrigin, zoom)

		renderpos := GraphtoScreenCords(rl.NewVector2(xgpos.X, fun(xgpos.X)), CameraOrigin, zoom)
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
