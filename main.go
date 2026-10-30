package main

import (
	"image/color"
	"raylib-example/simplexnoise"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const playerRadius = 20
const gravity = 1200
const jumpStrength = -400
const playerX = 200
const minY = playerRadius

var playerY float64 = 225
var velocityY float64
var gameOver = false
var deltaTime float64
var maxY float64
var gameWorldX float64

var obstacles = []*obstacle{
	{500, 100, 100, 250, rl.Brown},
	{900, 200, 100, 250, rl.Brown},
}

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

func worldNoise(x float64) float64 {
	return (simplexnoise.Noise1((x)*0.002) + 1) / 2 * float64(rl.GetScreenHeight())
}

func (o *obstacle) PositionY() {
	y := worldNoise(gameWorldX + o.x)
	if y < float64(rl.GetScreenHeight())/2 {
		o.y = y + playerRadius*4
	} else {
		o.y = y - playerRadius*4 - o.h
	}
}

func drawGameScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Lime)

	for _, o := range obstacles {
		o.Draw()
	}

	rl.DrawCircle(
		int32(playerX),
		int32(playerY),
		playerRadius,
		rl.SkyBlue,
	)

	if gameOver {
		rl.DrawText("GAME OVER", 270, 170, 30, rl.Red)
		rl.DrawText(
			"Drücke R zum Neustarten",
			220, 215, 20, rl.Black,
		)
	}

	linePoints := []rl.Vector2{}
	for x := 0.0; x < float64(rl.GetScreenWidth()); x++ {
		y := worldNoise(gameWorldX + x)
		linePoints = append(linePoints, rl.Vector2{float32(x), float32(y)})
	}
	rl.DrawLineStrip(linePoints, rl.DarkPurple)

	rl.EndDrawing()
}

func handleRestart() {
	if gameOver && rl.IsKeyPressed(rl.KeyR) {
		playerY = 225
		velocityY = 0
		gameOver = false
		gameWorldX = 0

		obstacles[0].x = 500
		obstacles[0].y = 100

		obstacles[1].x = 900
		obstacles[1].y = 200
	}
}

func handleControl() {
	if !gameOver {
		// Fliegen mit der Leertaste
		if rl.IsKeyPressed(rl.KeySpace) {
			velocityY = jumpStrength
		}

		// Physik: Schwerkraft und Bewegung
		velocityY += gravity * deltaTime
		playerY += velocityY * deltaTime

		// Oberen und unteren Bildschirmrand prüfen
		if playerY < minY || playerY > maxY {
			gameOver = true
		}
	}
}

func gameWorldControl() {
	if gameOver {
		return
	}

	gameWorldX += deltaTime * 100

	// Hindernisse bewegen und Kollision prüfen
	for _, o := range obstacles {
		o.x -= deltaTime * 100

		if o.x+o.w < 0 {
			o.x = 800
		}

		o.PositionY()

		if o.IsColliding(playerX, playerY) {
			gameOver = true
		}
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(800, 450, "Mein Flappy Bird")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	velocityY = float64(0)

	for _, o := range obstacles {
		o.PositionY()
	}

	for !rl.WindowShouldClose() {
		deltaTime = float64(rl.GetFrameTime())
		maxY = float64(rl.GetScreenHeight()) - playerRadius

		handleRestart()
		handleControl()
		gameWorldControl()

		drawGameScreen()
	}
}
