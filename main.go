package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(800, 450, "anderer Text")
	defer rl.CloseWindow()

	var posY float64 = 225
	var velocityY float64 = 0
	const gravity = 1200
	const radius = 50
	var minY float64 = radius

	for !rl.WindowShouldClose() {
		deltaTime := float64(rl.GetFrameTime())
		if rl.IsKeyPressed(rl.KeySpace) {
			velocityY = -500
		}
		velocityY += gravity * deltaTime
		yDirection := velocityY
		posY += yDirection * deltaTime

		if posY < minY {
			posY = minY
		} else {
			var maxY float64 = float64(rl.GetScreenHeight()) - radius
			if posY > maxY {
				posY = maxY
				velocityY = 0
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Lime)
		rl.DrawCircle(200, int32(posY), radius,
			rl.SkyBlue)
		rl.EndDrawing()
	}
}
