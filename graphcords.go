package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GraphtoScreenCords(v1 rl.Vector2) rl.Vector2 {
	ww, wh := rl.GetScreenWidth(), rl.GetScreenHeight()
	fww := float32(ww)
	fwh := float32(wh)

	camCords := rl.NewVector2(-CameraOrigin.X, CameraOrigin.Y)

	newVec := rl.NewVector2(v1.X, -v1.Y)
	newVec = rl.Vector2Add(newVec, camCords)
	newVec = rl.Vector2Scale(newVec, Zoom)
	newVec = rl.Vector2Add(newVec, rl.NewVector2(fww/2, fwh/2))
	return newVec
}

func ScreenToGraphCords(screenPos rl.Vector2) rl.Vector2 {
	ww, wh := rl.GetScreenWidth(), rl.GetScreenHeight()
	fww := float32(ww)
	fwh := float32(wh)

	camCords := rl.NewVector2(-CameraOrigin.X, CameraOrigin.Y)

	centeredVec := rl.Vector2Subtract(screenPos, rl.NewVector2(fww/2, fwh/2))
	cameraAdjustVec := rl.Vector2Subtract(centeredVec, camCords)
	scaledVec := rl.Vector2Scale(cameraAdjustVec, 1.0/Zoom)
	graphVec := rl.NewVector2(scaledVec.X, -scaledVec.Y)

	return graphVec
}

func GetScreenSizeGraphCords() rl.Vector2 {
	ww, wh := rl.GetScreenWidth(), rl.GetScreenHeight()

	return ScreenToGraphCords(rl.NewVector2(float32(ww), float32(wh)))
}
