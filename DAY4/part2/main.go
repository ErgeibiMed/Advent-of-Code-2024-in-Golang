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

	appears := appearence(string(bytes_raw))
	fmt.Println("counted ", appears)
}

func appearence(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	apprns := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2][2]int{
		{
			{-1, 1},
			{1, -1},
		},
		{
			{-1, -1},
			{1, 1},
		},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] != 'A' {
				continue
			}

			isXmas := true

			for _, dir := range directions {
				row1Index := row + dir[0][0]
				col1Index := col + dir[0][1]

				row2Index := row + dir[1][0]
				col2Index := col + dir[1][1]

				if row1Index < 0 || row1Index >= len(grid) || col1Index < 0 || col1Index >= len(grid[row]) || row2Index < 0 || row2Index >= len(grid) || col2Index < 0 || col2Index >= len(grid[row]) {
					isXmas = false
					break
				}

				if (grid[row1Index][col1Index] == 'M' && grid[row2Index][col2Index] == 'S') || (grid[row1Index][col1Index] == 'S' && grid[row2Index][col2Index] == 'M') {
					continue
				}

				isXmas = false
				break
			}

			if isXmas {
				apprns++
			}
		}
	}

	return apprns
}
