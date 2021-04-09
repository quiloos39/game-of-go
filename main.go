package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Terminal size
const (
	WIDTH  = 10
	HEIGHT = 5
)

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

func clearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func draw(board []int, cursor []int) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if x == cursor[0] && y == cursor[1] {
				fmt.Print("x")
			} else {
				fmt.Print(getCell(board[:], x, y))
			}
		}
		if y == 0 {
			fmt.Printf("   Simulating: %t", false)
		} else if y == 1 {
			fmt.Printf("   Generation: %d", 0)
		}
		fmt.Print("\n")
	}
}

func control(cursor []int, generation *int, simulation *bool) {
	stdin := make([]byte, 3)
	os.Stdin.Read(stdin)

	if stdin[0] == 27 && stdin[1] == 91 {
		// Right arrow key.
		if stdin[2] == 67 {
			cursor[0] = cursor[0] + 1
		}
		// Left arrow key.
		if stdin[2] == 68 {
			cursor[0] = cursor[0] - 1
		}
		// Down arrow key.
		if stdin[2] == 65 {
			cursor[1] = cursor[1] - 1
		}
		// Up arrow key.
		if stdin[2] == 66 {
			cursor[1] = cursor[1] + 1
		}
	}

	// Make sure cursor is between boundaries.
	if cursor[0] < 0 {
		cursor[0] = 0
	} else if cursor[0] >= WIDTH {
		cursor[0] = WIDTH - 1
	}

	if cursor[1] < 0 {
		cursor[1] = 0
	} else if cursor[1] >= HEIGHT {
		cursor[1] = HEIGHT - 1
	}

	// Clear stdin not sure if necessary but did anyway.
	stdin[0] = 0
	stdin[1] = 0
	stdin[2] = 0
}

func main() {
	// https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter/17278776#17278776
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()              // Do not display entered characters on the screen

	var board [WIDTH * HEIGHT]int
	cursor := [2]int{0, 0}
	generation := 0
	simulating := false

	for {
		clearTerm()
		go control(cursor[:], &generation, &simulating)
		draw(board[:], cursor[:])
		time.Sleep(100 * time.Millisecond)
	}

}
