package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// will be set in "initBoard"
type Board struct {
	width          int // In cells
	height         int // In cells
	numBombs       int
	cells          []int // -1 bomb, 0 empty, 1: one bomb nearby, 2: two bombs nearby etc..
	cellStates     []int // 0: hidden, 1: revealed, 2: flagged
	cellSize       int
	bombCell       int
	emptyCell      int
	xOffset        int
	yOffset        int
	elapsedSeconds float32
	lastTime       time.Time
}

func newBoard(width, height, numBombs int) *Board {
	// Make sure window does not get to big
	if width > 80 {
		width = 80
	}
	if height > 50 {
		height = 50
	}
	if width < 7 {
		width = 7
	}
	if height < 7 {
		height = 7
	}

	board := &Board{width: width, height: height, cellSize: 32, bombCell: -1, emptyCell: 0, xOffset: 10, yOffset: 50}

	// Number of bombs cannot be more than 60% and less thatn 10% of the board
	if numBombs > int(float32(width*height)*0.60) {
		numBombs = int(float32(width*height) * 0.60)
	} else if numBombs < int(float32(width*height)*0.10) {
		numBombs = int(float32(width*height) * 0.10)
	}
	board.numBombs = numBombs

	board.cells = make([]int, width*height)      // -1 bomb, 0 empty, 1: one bomb nearby, 2: two bombs nearby etc..
	board.cellStates = make([]int, width*height) // 0: hidden, 1: revealed, 2: flagged
	rl.SetWindowSize(width*board.cellSize+board.xOffset*2, height*board.cellSize+board.yOffset+board.xOffset)

	// Placing bombs
	rand.Seed(int64(time.Now().UnixMilli()))
	randomCellIndices := rand.Perm(len(board.cells))[:numBombs]
	for i := range randomCellIndices {
		board.cells[randomCellIndices[i]] = board.bombCell
	}
	for i := range board.cellStates {
		board.cellStates[i] = CELL_HIDDEN
	}
	board.initNumbers()
	board.resetTime()
	return board
}

func (b *Board) screenToCell(screenX, screenY int32) int {
	if screenY < int32(b.yOffset) {
		return -1
	}
	cellX := int(screenX-int32(b.xOffset)) / b.cellSize
	cellY := int(screenY-int32(b.yOffset)) / b.cellSize
	return cellX + cellY*b.width
}

func (b *Board) update() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cellIdx := b.screenToCell(rl.GetMouseX(), rl.GetMouseY())
		if cellIdx < 0 {
			return
		}
		b.reveal(cellIdx)
		if b.isGameWon() {
			gameState = GAMESTATE_GAME_WON
		}
	} else if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		cellIdx := b.screenToCell(rl.GetMouseX(), rl.GetMouseY())
		if cellIdx < 0 {
			return
		}
		if b.cellStates[cellIdx] == CELL_REVEALED {
			return
		}
		if b.cellStates[cellIdx] == CELL_HIDDEN {
			b.cellStates[cellIdx] = CELL_FLAGGED
		} else if b.cellStates[cellIdx] == CELL_FLAGGED {
			b.cellStates[cellIdx] = CELL_HIDDEN
		}
	}
}

func (b *Board) isGameWon() bool {
	numUnrevealedCells := 0
	for i := 0; i < len(b.cellStates); i++ {
		if b.cellStates[i] == CELL_HIDDEN || b.cellStates[i] == CELL_FLAGGED {
			numUnrevealedCells++
		}
	}
	return numUnrevealedCells == b.numBombs
}

func (b *Board) draw() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			idx := x + y*b.width
			if b.cellStates[idx] == CELL_HIDDEN {
				if b.cells[idx] == b.bombCell && gameState == GAMESTATE_GAME_OVER {
					rl.DrawTexture(bombBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(200, 200, 200, 200))
				} else if b.cells[idx] == b.bombCell && gameState == GAMESTATE_GAME_WON {
					rl.DrawTexture(bombBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(200, 200, 200, 200))
				} else {
					if b.cells[idx] == b.bombCell && DEBUG {
						if idx < b.width {
							rl.DrawTexture(grassBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(255, 120, 120, 200))
						} else {
							rl.DrawTexture(dirtBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(255, 120, 120, 200))
						}
					} else {
						if idx < b.width {
							rl.DrawTexture(grassBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(200, 200, 200, 200))
						} else {
							rl.DrawTexture(dirtBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(200, 200, 200, 200))
						}
					}
				}
			} else if b.cellStates[idx] == CELL_FLAGGED {
				rl.DrawTexture(trapBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(200, 200, 200, 200))
			} else if b.cellStates[idx] == CELL_REVEALED {
				if idx < b.width {
					rl.DrawTexture(grassBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(100, 100, 100, 200))
				} else {
					rl.DrawTexture(dirtBlock, int32(x*b.cellSize+b.xOffset), int32(y*b.cellSize+b.yOffset), rl.NewColor(100, 100, 100, 200))
				}
				if b.cells[idx] > b.emptyCell {
					nx := int32(x*b.cellSize+(b.cellSize/2)) - rl.MeasureText(strconv.Itoa(b.cells[idx]), 20)/2 + int32(b.xOffset)
					ny := int32(y*b.cellSize+10) + int32(b.yOffset)

					if b.cells[idx] <= 1 {
						rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(200, 200, 200, 255))
					} else if b.cells[idx] <= 2 {
						rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(0, 180, 255, 255))
					} else if b.cells[idx] <= 3 {
						rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(0, 255, 180, 255))
					} else if b.cells[idx] <= 4 {
						rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(255, 180, 0, 255))
					} else if b.cells[idx] <= 8 {
						rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(255, 80, 80, 255))
					}
					// rl.DrawText(strconv.Itoa((b.cells[idx])), nx-2, ny-4, 20, rl.NewColor(255, 203, 0, 255))
				}
			}
		}
	}
}

func (b *Board) reveal(cellIdx int) {
	if b.cellStates[cellIdx] == CELL_REVEALED || b.cellStates[cellIdx] == CELL_FLAGGED {
		return
	}

	if b.cells[cellIdx] == b.bombCell {
		gameState = GAMESTATE_GAME_OVER
		return
	}

	b.cellStates[cellIdx] = CELL_REVEALED

	if b.cells[cellIdx] > b.emptyCell {
		return
	}

	neighbors := b.getNeighbors(cellIdx)
	for i := 0; i < len(neighbors); i++ {
		if b.cellStates[neighbors[i]] == CELL_REVEALED || b.cellStates[neighbors[i]] == CELL_FLAGGED || b.cells[neighbors[i]] == b.bombCell {
			continue
		}

		b.reveal(neighbors[i])
	}
}

func (b *Board) initNumbers() {
	for i := 0; i < len(b.cells); i++ {
		if b.cells[i] == b.bombCell {
			continue
		}
		numBombs := 0
		neighbors := b.getNeighbors(i)
		for j := 0; j < len(neighbors); j++ {
			if b.cells[neighbors[j]] == b.bombCell {
				numBombs++
			}
		}
		b.cells[i] = numBombs
	}
}

func (b *Board) getNumFlaggedCells() int {
	numFlaggedCells := 0
	for i := 0; i < len(b.cells); i++ {
		if b.cellStates[i] == CELL_FLAGGED {
			numFlaggedCells++
		}
	}
	return numFlaggedCells
}

func (b *Board) getNeighbors(cellIdx int) []int {
	indices := make([]int, 0)
	top := cellIdx - b.width
	topLeft := top - 1
	topRight := top + 1
	bottom := cellIdx + b.width
	bottomLeft := bottom - 1
	bottomRight := bottom + 1
	left := cellIdx - 1
	right := cellIdx + 1

	if top >= 0 {
		indices = append(indices, top)
		if cellIdx%b.width != 0 {
			indices = append(indices, topLeft)
		}
		if (cellIdx+1)%b.width != 0 {
			indices = append(indices, topRight)
		}
	}
	if bottom <= b.width*b.height-1 {
		indices = append(indices, bottom)
		if cellIdx%b.width != 0 {
			indices = append(indices, bottomLeft)
		}
		if (cellIdx+1)%b.width != 0 {
			indices = append(indices, bottomRight)
		}
	}
	if left >= 0 {
		if cellIdx%b.width != 0 {
			indices = append(indices, left)
		}
	}
	if right <= b.width*b.height-1 {
		if (cellIdx+1)%b.width != 0 {
			indices = append(indices, right)
		}
	}

	return indices
}

func (b *Board) resetTime() {
	b.elapsedSeconds = 0
	b.lastTime = time.Now()
}

func (b *Board) updateTime() {
	elapsedTime := time.Since(b.lastTime).Milliseconds()
	b.lastTime = time.Now()
	b.elapsedSeconds += float32(elapsedTime) / 1000
}
