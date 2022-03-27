package sources

import rl "github.com/gen2brain/raylib-go/raylib"

func Loop() {
	for !rl.WindowShouldClose() && !Terminate {
		beforeWindow()
		if swapMinPos {
			goto skipRender
		}
		drawWindow()
		AppContent()

		afterWindow()
	skipRender:
		if isMinimizing || swapMinPos {
			Minimize()
		}
	}
}

func AppContent() {

}