package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const playerRadius = 20

type obstacle struct {
	x float64
	y float64
	w float64
	h float64
	c color.RGBA
}

func (o obstacle) Draw() {
	rl.DrawRectangle(
		int32(o.x),
		int32(o.y),
		int32(o.w),
		int32(o.h),
		o.c,
	)
}

func (o obstacle) IsColliding(playerX, playerY float64) bool {
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
	rl.InitWindow(800, 450, "Mein Flappy Bird")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	posY := float64(225)
	velocityY := float64(0)

	const gravity = 1200.0
	const jumpStrength = -500.0
	const playerX = 200.0

	minY := float64(playerRadius)
	maxY := float64(450 - playerRadius)

	obstacles := []*obstacle{
		{500, 100, 100, 150, rl.Brown},
		{900, 200, 100, 150, rl.Brown},
	}

	gameOver := false

	for !rl.WindowShouldClose() {
		deltaTime := float64(rl.GetFrameTime())

		// Neustart mit R
		if gameOver && rl.IsKeyPressed(rl.KeyR) {
			posY = 225
			velocityY = 0
			gameOver = false

			obstacles[0].x = 500
			obstacles[0].y = 100

			obstacles[1].x = 900
			obstacles[1].y = 200
		}

		if !gameOver {
			// Fliegen mit der Leertaste
			if rl.IsKeyPressed(rl.KeySpace) {
				velocityY = jumpStrength
			}

			// Physik: Schwerkraft und Bewegung
			velocityY += gravity * deltaTime
			posY += velocityY * deltaTime

			// Oberen und unteren Bildschirmrand prüfen
			if posY < minY || posY > maxY {
				gameOver = true
			}

			// Hindernisse bewegen und Kollision prüfen
			for _, o := range obstacles {
				o.x -= deltaTime * 100

				if o.x+o.w < 0 {
					o.x = 800
				}

				if o.IsColliding(playerX, posY) {
					gameOver = true
				}
			}
		}

		// Zeichnen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Lime)

		for _, o := range obstacles {
			o.Draw()
		}

		rl.DrawCircle(
			int32(playerX),
			int32(posY),
			playerRadius,
			rl.SkyBlue,
		)

		if gameOver {
			rl.DrawText("GAME OVER", 270, 170, 30, rl.Red)
			rl.DrawText(
				"Druecke R zum Neustarten",
				220, 215, 20, rl.Black,
			)
		}

		rl.EndDrawing()
	}
}

