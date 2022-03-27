package sources

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Size = [2]int32{800, 600}
var Terminate = false
var mouse rl.Vector2
var exitIcons = [3]rl.Vector2{{20, 12}, {40, 12}, {60, 12}}
var Freeze = false
var Frame = 0
var isMinimizing = false
var shouldTerminate = false
var TitleBarSize = [2]int32{Size[0], 24}
var AreaSize = [2]int32{Size[0], Size[1] - TitleBarSize[1]}

func Setup() {
	rl.SetConfigFlags(rl.FlagWindowAlwaysRun)
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Size[0], Size[1], "")
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	for !rl.WindowShouldClose() && !Terminate {
		Frame++
		if Frame >= 60 {
			Frame = 0
		}
		if !Freeze {
			mouse = rl.GetMousePosition()
		}
		if swapMinPos {
			goto skipRender
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.Color{})
		drawDecorations()
		drawBlankArea()
		rl.EndDrawing()

		if !Freeze {
			handleMenuBar()
		}

	skipRender:
		if isMinimizing || swapMinPos {
			Minimize()
		}
	}

	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func drawDecorations() {
	rl.DrawRectangle(10, 0, Size[0]-20, 24, rl.Color{R: 205, G: 205, B: 205, A: 255})
	rl.DrawCircle(10, 12, 12, rl.Color{R: 205, G: 205, B: 205, A: 255})
	rl.DrawCircle(Size[0]-10, 12, 12, rl.Color{R: 205, G: 205, B: 205, A: 255})
	rl.DrawRectangle(0, 12, Size[0], 12, rl.Color{R: 205, G: 205, B: 205, A: 255})
	rl.DrawCircle(int32(exitIcons[0].X), int32(exitIcons[0].Y), 6, rl.Color{R: 255, G: 95, B: 87, A: 255})
	rl.DrawCircle(int32(exitIcons[1].X), int32(exitIcons[1].Y), 6, rl.Color{R: 255, G: 190, B: 47, A: 255})
	rl.DrawCircle(int32(exitIcons[2].X), int32(exitIcons[2].Y), 6, rl.Color{R: 41, G: 204, B: 65, A: 255})
}

func drawBlankArea() {
	rl.DrawRectangle(0, TitleBarSize[1], AreaSize[0], AreaSize[1], rl.RayWhite)
}

var tmpFrame = 0
var winMinPos rl.Vector2
var winMinSize = rl.Vector2{X: float32(Size[0]), Y: float32(Size[1])}
var swapMinPos = false

func handleMenuBar() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointCircle(mouse, exitIcons[0], 6) {
			Freeze = true
			isMinimizing = true
			tmpFrame = 0
			swapMinPos = false
			winMinPos = rl.GetWindowPosition()
			shouldTerminate = true
		}
		if rl.CheckCollisionPointCircle(mouse, exitIcons[1], 6) {
			Freeze = true
			isMinimizing = true
			tmpFrame = 0
			swapMinPos = false
			winMinPos = rl.GetWindowPosition()
			shouldTerminate = false
		}
		if rl.CheckCollisionPointCircle(mouse, exitIcons[2], 6) {
			rl.ToggleFullscreen()
		}
	}
}

func Minimize() {
	tmpFrame++
	if swapMinPos {
		swapMinPos = false
		rl.MinimizeWindow()
		if shouldTerminate {
			Terminate = true
		}
		return
	}
	rl.SetWindowPosition(int(winMinPos.X)+tmpFrame*50, int(winMinPos.Y)+tmpFrame*100)
	rl.SetWindowSize(int(winMinSize.X/(float32(tmpFrame/(15/10)))), int(winMinSize.Y/(float32(tmpFrame/(15/10)))))
	if tmpFrame >= 20 {
		isMinimizing = false
		swapMinPos = true
		Freeze = false
		rl.SetWindowSize(int(winMinSize.X), int(winMinSize.Y))
		rl.SetWindowPosition(int(winMinPos.X), int(winMinPos.Y))
	}
}
