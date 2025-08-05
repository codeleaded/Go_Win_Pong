package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

const screenWidth = 800
const screenHeight = 600

func main(){
	rl.InitWindow(screenWidth, screenHeight, "Pong in Go mit raylib")
	defer rl.CloseWindow()

	paddleSpeed := float32(600)
	ballSpeed := float32(400)
	player := rl.NewRectangle(50, float32(screenHeight/2-60), 15, 120)
	computer := rl.NewRectangle(float32(screenWidth-65), float32(screenHeight/2-60), 15, 120)
	ball := rl.NewRectangle(float32(screenWidth/2-10), float32(screenHeight/2-10), 20, 20)
	ballVelocity := rl.Vector2{X: ballSpeed, Y: ballSpeed}

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()

		if rl.IsKeyDown(rl.KeyUp) && player.Y > 0 {
			player.Y -= paddleSpeed * deltaTime
		}
		if rl.IsKeyDown(rl.KeyDown) && player.Y < float32(screenHeight)-player.Height {
			player.Y += paddleSpeed * deltaTime
		}

		if ball.Y+ball.Height/2 > computer.Y+computer.Height/2 {
			computer.Y += paddleSpeed * deltaTime
		} else if ball.Y+ball.Height/2 < computer.Y+computer.Height/2 {
			computer.Y -= paddleSpeed * deltaTime
		}

		ball.X += ballVelocity.X * deltaTime
		ball.Y += ballVelocity.Y * deltaTime

		if ball.Y <= 0 || ball.Y+ball.Height >= float32(screenHeight) {
			ballVelocity.Y = -ballVelocity.Y
		}

		if rl.CheckCollisionRecs(ball, player) {
			ballVelocity.X = -ballVelocity.X
		}
		if rl.CheckCollisionRecs(ball, computer) {
			ballVelocity.X = -ballVelocity.X
		}

		if ball.X <= 0 || ball.X+ball.Width >= float32(screenWidth) {
			ball.X = float32(screenWidth / 2)
			ball.Y = float32(screenHeight / 2)
			ballVelocity.X = -ballVelocity.X

			if math.Signbit(float64(ballVelocity.Y)) {
				ballVelocity.Y = -ballSpeed
			} else {
				ballVelocity.Y = ballSpeed
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangleRec(player,rl.White)
		rl.DrawRectangleRec(computer,rl.White)
		rl.DrawRectangleRec(ball,rl.White)

		rl.DrawText("Pong", 10, 10, 20, rl.White)

		rl.EndDrawing()
	}
}
