package tictactoe

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

// Constant declaration
const size = 3

// Define 2D board as a table of strings
// Player 1 = X; Player 2 = O
var board [size][size]string

func TicTacToe() {
	// Get standard input from the player
	scanner := bufio.NewScanner(os.Stdin)

	// Just "for" without anything else is an infinite loop, like "while True:" in Python
	for {

		var mode bool = startMenu() // false = person; true = computer

		// Initial state of the game - clear the board and start as player 1 'X'
		clearBoard()
		currentPlayer := "X"

		// Nested forever loops for each turn, which waits for input and breaks out of it when it's received
		for {
			drawBoard()            // Draw the board in its current state
			var newCol, newRow int // Initialize new row and column

			if !mode || currentPlayer == "X" {
				fmt.Printf("Player %s <col row>: ", currentPlayer) // Input prompt
				scanner.Scan()                                     // Scan for input
				move := strings.TrimSpace(scanner.Text())          // Trim whitespace around input

				// Sscanf:
				// First input - number of scanned items
				// Second input - error reported during scanning

				// This essentially converts the input to valid integers
				_, err := fmt.Sscanf(move, "%d %d", &newRow, &newCol)

				// Check boundaries on the input and respond appropriately
				if err != nil || newRow > (size-1) || newRow < 0 || newCol > (size-1) || newCol < 0 {
					fmt.Println("Invalid move")
					continue // Go back to the top of the loop
				}

				// Check if the move is already taken
				if board[newCol][newRow] != " " {
					fmt.Println("Position already taken")
					continue // Go back to the top of the loop
				}
			} else {
				// Computer moves
				newCol, newRow = chooseMove()
				fmt.Printf("The computer moved to: %d %d\n", newRow, newCol)
			}

			// Now there are no errors, the move is in bounds, and the spot is free; update the game board
			board[newCol][newRow] = currentPlayer

			// Check for a win
			if checkWin(currentPlayer) {
				fmt.Printf("Player %s won!\n", currentPlayer)
				drawBoard()
				break
			}

			// If there is no win, check that the board isn't full
			if fullBoard() {
				fmt.Println("Draw!")
				break
			}

			// Switch player
			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		}

		// Prompt for user to play again
		fmt.Println("Press P to play again; press any other key to quit: ")
		scanner.Scan() // Scan for input

		// Get the response
		response := strings.TrimSpace(scanner.Text())

		// If the response is to play again, go back to the top
		if response == "P" {
			continue
		}
		break // Otherwise, break out
	}
}

func checkWin(player string) bool {
	// Check horizontal/vertical win (not sure if this is efficient)
	for col := range board {
		if board[col][0] == player && board[col][1] == player && board[col][2] == player {
			return true
		}
		for row := range board[col] {
			// Horizontal win - this is redundant but I don't care it's tic tac toe
			if board[0][row] == player && board[1][row] == player && board[2][row] == player {
				return true
			}
		}
	}

	// Check diagonal win
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	} else if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

func startMenu() bool {
	scanner := bufio.NewScanner(os.Stdin) // local scanner (regular scanner out of scope) - this kinda sucks but it is what it is

	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("- Press <1> to play against another human")
	fmt.Println("- Press any other key to play against the computer")
	scanner.Scan()

	response := strings.TrimSpace(scanner.Text())
	return !(response == "1") // Only set to "true" if the response is not 1
}

// Computer chooses a move; choose completely random open space
func chooseMove() (int, int) {
	var col, row int = rand.IntN(3), rand.IntN(3)

	// Keep choosing while not empty
	for board[col][row] != " " {
		col, row = rand.IntN(3), rand.IntN(3)
	}

	// Return location to move
	return col, row
}

func fullBoard() bool {
	// Just iterate over all spaces and check for a blank
	for col := range board {
		for row := range board[col] {
			if board[row][col] == " " {
				return false
			}
		}
	}
	return true
}

func drawBoard() {
	// Iterate over the game board using this paradigm
	fmt.Println("   0   1   2") // Print column numbers
	for col, row := range board {
		// Formatted print - similar to C
		// Print row contents and row numbers
		fmt.Printf("%d  %s | %s | %s\n", col, row[0], row[1], row[2])

		// Print line to separate board rows
		if col < size-1 {
			fmt.Println("  -----------")
		}
	}
	fmt.Println() // Newline
}

func clearBoard() {
	for col := range board {
		for row := range board[col] {
			board[col][row] = " "
		}
	}
}
