package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scrWidth  = 1000
	scrHeight = 1000
	numCellsX = 50
	numCellsY = 50
	dimCellW  = scrWidth / numCellsX
	dimCellH  = scrHeight / numCellsY
)

type GameState [numCellsY][numCellsX]int

var (
	pauseGame    bool
	initialState GameState
)

func DrawState(state GameState) {
	rl.BeginDrawing()
	for y := 0; y < numCellsY; y++ {
		for x := 0; x < numCellsX; x++ {
			if state[y][x] == 0 {
				rl.DrawRectangle(int32(x*dimCellW), int32(y*dimCellH), dimCellW, dimCellH, rl.Black)
			} else if state[y][x] == 1 {
				rl.DrawRectangle(int32(x*dimCellW), int32(y*dimCellH), dimCellW, dimCellH, rl.White)
			}
			rl.DrawRectangleLines(int32(x*dimCellW), int32(y*dimCellH), dimCellW, dimCellH, rl.Gray)
		}
	}
	rl.EndDrawing()
}

func CheckAlive(state GameState, x int, y int) int {
	var resNum int

	anteriorY := (y - 1) % numCellsY
	if anteriorY < 0 {
		anteriorY += numCellsY
	}
	posteriorY := (y + 1) % numCellsY
	anteriorX := (x - 1) % numCellsX
	if anteriorX < 0 {
		anteriorX += numCellsX
	}
	posteriorX := (x + 1) % numCellsX

	vecinos := state[anteriorY][anteriorX] +
		state[y%numCellsY][anteriorX] +
		state[posteriorY][anteriorX] +
		state[anteriorY][x%numCellsX] +
		state[posteriorY][x%numCellsX] +
		state[anteriorY][posteriorX] +
		state[y%numCellsY][posteriorX] +
		state[posteriorY][posteriorX]

	if state[y][x] == 0 && vecinos == 3 {
		resNum = 1
	} else if state[y][x] == 1 && (vecinos < 2 || vecinos > 3) {
		resNum = 0
	} else if state[y][x] == 1 && (vecinos == 2 || vecinos == 3) {
		resNum = 1
	}
	return resNum
}

func UpdateGame(state GameState) GameState {
	var newState [numCellsY][numCellsX]int
	newState = state
	for y := 0; y < numCellsY; y++ {
		for x := 0; x < numCellsX; x++ {
			newState[y][x] = CheckAlive(state, x, y)

			// Añadir entidades con el mouse
			positionMouse := rl.GetMousePosition()
			Xcoord := int(math.Floor(float64(positionMouse.X) / float64(dimCellW)))
			Ycoord := int(math.Floor(float64(positionMouse.Y) / float64(dimCellH)))

			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				newState[Ycoord][Xcoord] = 1
			}
			if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
				newState[Ycoord][Xcoord] = 0
			}
		}
	}
	return newState
}

func main() {
	rl.InitWindow(scrWidth, scrHeight, "Game Of Life by AdriCort12")
	defer rl.CloseWindow()

	currentState := initialState

	// Automatá móvil
	currentState[13][20] = 1
	currentState[14][20] = 1
	currentState[15][20] = 1
	currentState[15][19] = 1
	currentState[14][18] = 1

	// Automata Palo
	currentState[16][20] = 1
	currentState[16][21] = 1
	currentState[16][22] = 1

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		time.Sleep(100 * time.Millisecond)
		DrawState(currentState)
		nextState := UpdateGame(currentState)

		// Estado de pausa
		if rl.IsKeyDown(rl.KeySpace) {
			pauseGame = !pauseGame
		}
		if !pauseGame {
			currentState = nextState
		}

	}
}
