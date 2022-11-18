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

var bombAmount10Btn *Button = NewButton("10%", 42, 240, 70, 30, rl.White, rl.Gold)
var bombAmount15Btn *Button = NewButton("15%", 122, 240, 70, 30, rl.White, rl.Gold)
var bombAmount20Btn *Button = NewButton("20%", 202, 240, 70, 30, rl.White, rl.Gold)
var bombAmount25Btn *Button = NewButton("25%", 282, 240, 70, 30, rl.White, rl.Gold)

var bombAmount30Btn *Button = NewButton("30%", 42, 280, 70, 30, rl.White, rl.Gold)
var bombAmount35Btn *Button = NewButton("35%", 122, 280, 70, 30, rl.White, rl.Gold)
var bombAmount40Btn *Button = NewButton("40%", 202, 280, 70, 30, rl.White, rl.Gold)
var bombAmount45Btn *Button = NewButton("45%", 282, 280, 70, 30, rl.White, rl.Gold)

var startNewGameBtn *Button = NewButton("Go", 160, 350, 70, 30, rl.White, rl.Gold)

func main() {
	rl.InitWindow(800, 450, "MineCreeper v1.0.5")
	rl.SetTargetFPS(20)
	loadAssets()
	gameSizeSmallBtn.Selected = true
	bombAmount10Btn.Selected = true

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
			bombAmount10Btn.Draw()
			bombAmount15Btn.Draw()
			bombAmount20Btn.Draw()
			bombAmount25Btn.Draw()
			bombAmount30Btn.Draw()
			bombAmount35Btn.Draw()
			bombAmount40Btn.Draw()
			bombAmount45Btn.Draw()

			if bombAmount10Btn.Pressed() {
				bombAmount10Btn.Selected = true
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount15Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = true
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount20Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = true
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount25Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = true
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount30Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = true
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount35Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = true
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = false
			}
			if bombAmount40Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = true
				bombAmount45Btn.Selected = false
			}
			if bombAmount45Btn.Pressed() {
				bombAmount10Btn.Selected = false
				bombAmount15Btn.Selected = false
				bombAmount20Btn.Selected = false
				bombAmount25Btn.Selected = false
				bombAmount30Btn.Selected = false
				bombAmount35Btn.Selected = false
				bombAmount40Btn.Selected = false
				bombAmount45Btn.Selected = true
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
	fontSize := int32(board.cellSize / 2)
	// Number of bombs
	numBombs := board.numBombs - board.getNumFlaggedCells()
	rl.DrawText(strconv.Itoa(numBombs), 100-rl.MeasureText(strconv.Itoa(numBombs), fontSize)/2, int32(fontSize/2), fontSize, rl.Gold)
	// Time
	timeX := int32(board.width*board.cellSize) - 100 + int32(board.xOffset*2)
	rl.DrawText(strconv.Itoa(int(board.elapsedSeconds)), timeX-rl.MeasureText(strconv.Itoa(int(board.elapsedSeconds)), fontSize)/2, int32(fontSize/2), fontSize, rl.Gold)
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
		boardWidth = 7
		boardHeight = 7
	} else if gameSizeMediumBtn.Selected {
		boardWidth = 12
		boardHeight = 12
	} else if gameSizeLargeBtn.Selected {
		boardWidth = 17
		boardHeight = 17
	} else if gameSizeHugeBtn.Selected {
		boardWidth = 33
		boardHeight = 19
	}

	if bombAmount10Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.10)
	} else if bombAmount15Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.15)
	} else if bombAmount20Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.20)
	} else if bombAmount25Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.25)
	} else if bombAmount30Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.30)
	} else if bombAmount35Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.35)
	} else if bombAmount40Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.40)
	} else if bombAmount45Btn.Selected {
		bombAmount = int(float32(boardWidth*boardHeight) * 0.45)
	}

	gameState = GAMESTATE_IN_PROGRESS
	return newBoard(boardWidth, boardHeight, bombAmount)
}
