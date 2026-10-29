package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const playerRadius = 50

type obstacle struct {
	x float64
	y float64
	w float64
	h float64
	c color.RGBA
}

func (o obstacle) Draw() {
	rl.DrawRectangle(int32(o.x), int32(o.y), int32(o.w), int32(o.h), o.c)
}

func (o obstacle) IsColliding(playerX float64, playerY float64) bool {
	playerTop := playerY - playerRadius
	playerBottom := playerY + playerRadius
	playerLeft := playerX - playerRadius
	playerRight := playerX + playerRadius

	return playerRight > o.x &&
		playerLeft < o.x+o.w &&
		playerBottom > o.y &&
		playerTop < o.y+o.h
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(800, 450, "anderer Text")
	defer rl.CloseWindow()

	var posY float64 = 225
	var velocityY float64 = 0
	const gravity = 1200
	var minY float64 = playerRadius

	obstacles := []*obstacle{
		&obstacle{500, 100, 100, 150, rl.Brown},
		&obstacle{1000, 200, 100, 150, rl.Brown},
	}

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
			var maxY float64 = float64(rl.GetScreenHeight()) - playerRadius
			if posY > maxY {
				posY = maxY
				velocityY = 0
			}
		}

		for _, o := range obstacles {
			o.x -= deltaTime * 100
			o.Draw()
			if o.IsColliding(200, posY) {
				rl.DrawText("Game over", 270, 100, 24, rl.Red)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Lime)
		rl.DrawCircle(200, int32(posY), playerRadius,
			rl.SkyBlue)
		rl.EndDrawing()
	}
}
