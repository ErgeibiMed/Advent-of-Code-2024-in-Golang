package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes_raw, err := (os.ReadFile("./input.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)
	}

	occurences := occurence(string(bytes_raw))
	fmt.Printf("Xmas occured %d times \n", occurences)

}
func occurence(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	occurences := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	word := "XMAS"

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				rowIndex := dir[0]
				colIndex := dir[1]
				isXmas := true

				for charIndex := 0; charIndex < len(word); charIndex++ {
					rowOffset := row + (rowIndex * charIndex)
					colOffset := col + (colIndex * charIndex)

					if rowOffset < 0 || rowOffset >= len(grid) || colOffset < 0 || colOffset >= len(grid[row]) {
						isXmas = false
						break
					}

					if grid[rowOffset][colOffset] != rune(word[charIndex]) {
						isXmas = false
						break
					}
				}

				if isXmas {
					occurences++
				}
			}
		}
	}

	return occurences
}
