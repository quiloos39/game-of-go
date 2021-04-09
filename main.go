package main

import (
	"fmt"
	"os/exec"

	// "time"
	"bufio"
	"os"
)

const WIDTH = 10
const HEIGHT = 5

func getCell(board []int, x int, y int) int {
	if x < 0 || x > WIDTH {
		return 0
	} else if y < 0 || y > HEIGHT {
		return 0
	}
	return board[x+(y*WIDTH)]
}

// Since i am using integers to represent alive ( 1 ) and dead ( 0 ) i can just add them up.
func countNeighbor(board []int, x int, y int) int {
	counter := getCell(board[:], x+1, y) +
		getCell(board[:], x+1, y+1) +
		getCell(board[:], x+1, y-1) +
		getCell(board[:], x, y+1) +
		getCell(board[:], x, y-1) +
		getCell(board[:], x-1, y) +
		getCell(board[:], x-1, y+1) +
		getCell(board[:], x-1, y-1)
	return counter
}

func draw(board []int, cursor []int) {
	var out, _ = exec.Command("clear").Output()
	fmt.Printf("%s", out)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if x == cursor[0] && y == cursor[1] {
				fmt.Print("x")
			} else {
				fmt.Print(getCell(board[:], x, y))
			}
		}
		fmt.Print("\n")
	}
}

func controlCursor(cursor []int) {
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadByte()
	fmt.Printf("%s", key)
}

func main() {
	var board [WIDTH * HEIGHT]int
	cursor := [2]int{0, 0}
	for {
		draw(board[:], cursor[:])
		controlCursor(cursor[:])
		// time.Sleep(1000 * time.Millisecond)
	}
}
