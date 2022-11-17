package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var DEBUG = false
var cheated = false

const (
	CELL_HIDDEN = iota
	CELL_REVEALED
	CELL_FLAGGED
)
const (
	GAMESTATE_IN_PROGRESS = iota
	GAMESTATE_GAME_WON
	GAMESTATE_GAME_OVER
	GAMESTATE_CREATE_NEW_GAME
)

var clearColor rl.Color = rl.NewColor(30, 10, 0, 200)

var gameState = GAMESTATE_IN_PROGRESS

var gameSizeSmallBtn *Button = NewButton("small", 42, 120, 70, 30, rl.White, rl.Gold)
var gameSizeMediumBtn *Button = NewButton("medium", 122, 120, 70, 30, rl.White, rl.Gold)
var gameSizeLargeBtn *Button = NewButton("large", 202, 120, 70, 30, rl.White, rl.Gold)
var gameSizeHugeBtn *Button = NewButton("huge", 282, 120, 70, 30, rl.White, rl.Gold)

var bombAmountSmallBtn *Button = NewButton("10%", 42, 240, 70, 30, rl.White, rl.Gold)
var bombAmountMediumBtn *Button = NewButton("25%", 122, 240, 70, 30, rl.White, rl.Gold)
var bombAmountLargeBtn *Button = NewButton("50%", 202, 240, 70, 30, rl.White, rl.Gold)
var bombAmountHugeBtn *Button = NewButton("65%", 282, 240, 70, 30, rl.White, rl.Gold)

var startNewGameBtn *Button = NewButton("Go", 160, 350, 70, 30, rl.White, rl.Gold)

func main() {
	rl.InitWindow(800, 450, "MineCreeper v1.0")
	rl.SetTargetFPS(20)
	loadAssets()
	gameSizeSmallBtn.Selected = true
	bombAmountSmallBtn.Selected = true

	board := createNewBoard()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyF9) {
			DEBUG = !DEBUG
			cheated = true
		}

		checkNewGameButtonClick(board)
		if gameState == GAMESTATE_IN_PROGRESS {
			board.update()
		}
		rl.BeginDrawing()
		rl.ClearBackground(clearColor)

		switch gameState {
		case GAMESTATE_IN_PROGRESS:
			board.draw()
			drawGameGui(board)
			board.updateTime()
		case GAMESTATE_GAME_OVER:
			board.draw()
			x := int32(rl.GetScreenWidth()/2 - int(rl.MeasureText("GAME LOST", 30)/2))
			y := int32(rl.GetScreenHeight()/2) - 15
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.NewColor(120, 60, 60, 80))
			rl.DrawText("GAME LOST", x+2, y+2, 30, rl.Black)
			rl.DrawText("GAME LOST", x, y, 30, rl.Red)
			if cheated {
				rl.DrawText("CHEATER!", x+2+5, y+52, 30, rl.Black)
				rl.DrawText("CHEATER!", x+5, y+50, 30, rl.Red)
			}
			drawGameGui(board)
		case GAMESTATE_GAME_WON:
			board.draw()
			x := int32(rl.GetScreenWidth()/2 - int(rl.MeasureText("GAME WON", 30)/2))
			y := int32(rl.GetScreenHeight()/2) - 15
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.NewColor(60, 120, 60, 80))
			rl.DrawText("GAME WON", x+2, y+2, 30, rl.Black)
			rl.DrawText("GAME WON", x, y, 30, rl.Green)
			if cheated {
				rl.DrawText("CHEATER!", x+2+5, y+52, 30, rl.Black)
				rl.DrawText("CHEATER!", x+5, y+50, 30, rl.Red)
			}
			drawGameGui(board)
		case GAMESTATE_CREATE_NEW_GAME:
			rl.DrawText("NEW GAME", int32(rl.GetScreenWidth()/2)-rl.MeasureText("NEW GAME", 30)/2, 20, 30, rl.Gold)
			// Board size
			rl.DrawText("Board size:", 20, 80, 22, rl.Gold)
			gameSizeSmallBtn.Draw()
			gameSizeMediumBtn.Draw()
			gameSizeLargeBtn.Draw()
			gameSizeHugeBtn.Draw()
			if gameSizeSmallBtn.Pressed() {
				gameSizeSmallBtn.Selected = true
				gameSizeMediumBtn.Selected = false
				gameSizeLargeBtn.Selected = false
				gameSizeHugeBtn.Selected = false
			}
			if gameSizeMediumBtn.Pressed() {
				gameSizeMediumBtn.Selected = true
				gameSizeSmallBtn.Selected = false
				gameSizeLargeBtn.Selected = false
				gameSizeHugeBtn.Selected = false
			}
			if gameSizeLargeBtn.Pressed() {
				gameSizeLargeBtn.Selected = true
				gameSizeMediumBtn.Selected = false
				gameSizeSmallBtn.Selected = false
				gameSizeHugeBtn.Selected = false
			}
			if gameSizeHugeBtn.Pressed() {
				gameSizeHugeBtn.Selected = true
				gameSizeLargeBtn.Selected = false
				gameSizeMediumBtn.Selected = false
				gameSizeSmallBtn.Selected = false
			}

			// Board size
			rl.DrawText("Creeper amount:", 20, 200, 22, rl.Gold)
			bombAmountSmallBtn.Draw()
			bombAmountMediumBtn.Draw()
			bombAmountLargeBtn.Draw()
			bombAmountHugeBtn.Draw()

			if bombAmountSmallBtn.Pressed() {
				bombAmountSmallBtn.Selected = true
				bombAmountMediumBtn.Selected = false
				bombAmountLargeBtn.Selected = false
				bombAmountHugeBtn.Selected = false
			}
			if bombAmountMediumBtn.Pressed() {
				bombAmountMediumBtn.Selected = true
				bombAmountSmallBtn.Selected = false
				bombAmountLargeBtn.Selected = false
				bombAmountHugeBtn.Selected = false
			}
			if bombAmountLargeBtn.Pressed() {
				bombAmountLargeBtn.Selected = true
				bombAmountSmallBtn.Selected = false
				bombAmountMediumBtn.Selected = false
				bombAmountHugeBtn.Selected = false
			}
			if bombAmountHugeBtn.Pressed() {
				bombAmountHugeBtn.Selected = true
				bombAmountLargeBtn.Selected = false
				bombAmountSmallBtn.Selected = false
				bombAmountMediumBtn.Selected = false
			}

			startNewGameBtn.Draw()
			if startNewGameBtn.Pressed() {
				board = createNewBoard()
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func drawGameGui(board *Board) {
	// Number of bombs
	rl.DrawTexture(signTex, 0, int32(board.yOffset-board.cellSize), rl.White)
	numBombs := board.numBombs - board.getNumFlaggedCells()
	rl.DrawText(strconv.Itoa(numBombs), int32(board.xOffset*2)+int32(board.cellSize)-rl.MeasureText(strconv.Itoa(numBombs), 24)/2, int32(24/2), 24, rl.Gold)
	// Time
	timeSignX := int32((board.width*board.cellSize - board.cellSize*2) - board.xOffset*3)
	rl.DrawTexture(signTex, timeSignX, int32(board.yOffset-board.cellSize), rl.White)
	rl.DrawText(strconv.Itoa(int(board.elapsedSeconds)), timeSignX+int32(board.xOffset*2)+int32(board.cellSize)-rl.MeasureText(strconv.Itoa(int(board.elapsedSeconds)), 24)/2, int32(24/2), 24, rl.Gold)
	// New game button
	rl.DrawTexture(headTex, int32(((board.width*board.cellSize)/2)+board.xOffset-board.cellSize/2), int32(board.yOffset-board.cellSize)-5, rl.White)
}

func checkNewGameButtonClick(board *Board) {
	posX := int32(((board.width * board.cellSize) / 2) + board.xOffset - board.cellSize/2)
	posY := int32(board.yOffset-board.cellSize) - 5

	clicked := false
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.GetMouseX() > posX && rl.GetMouseX() < posX+int32(board.cellSize) {
			if rl.GetMouseY() > posY && rl.GetMouseY() < posY+int32(board.cellSize) {
				clicked = true
			}
		}
	}

	if clicked {
		rl.SetWindowSize(400, 400)
		gameState = GAMESTATE_CREATE_NEW_GAME
	}
}

func createNewBoard() *Board {
	DEBUG = false
	cheated = false
	boardWidth := 0
	boardHeight := 0
	bombAmount := 0

	if gameSizeSmallBtn.Selected {
		boardWidth = 8
		boardHeight = 8
	} else if gameSizeMediumBtn.Selected {
		boardWidth = 12
		boardHeight = 12
	} else if gameSizeLargeBtn.Selected {
		boardWidth = 20
		boardHeight = 20
	} else if gameSizeHugeBtn.Selected {
		boardWidth = 45
		boardHeight = 25
	}

	if bombAmountSmallBtn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.10)
	} else if bombAmountMediumBtn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.25)
	} else if bombAmountLargeBtn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.50)
	} else if bombAmountHugeBtn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.65)
	}

	gameState = GAMESTATE_IN_PROGRESS
	return newBoard(boardWidth, boardHeight, bombAmount)
}
