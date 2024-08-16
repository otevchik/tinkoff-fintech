package main

import (
	"fmt"
)

func sandglass(size int, char rune, color string) {
	for i := 0; i < size; i++ {
		start := i
		end := size - 1 - i
		for j := 0; j < size; j++ {
			switch {
			case (i == 0 || i == (size - 1)):
				fmt.Printf(color + "%c", char)
			case ((j == start || j == end) && start != end):
				fmt.Printf(color + "%c", char)
			case (start == end && j == start):
				fmt.Printf(color + "%c", char)
			default:
				fmt.Print(" ")
			}
		}
		if (i != size - 1) {
			fmt.Println()
		}
	}
}

func main () {
	size := 14
	color := "\033[31m"
	char := 'B'
	if size >= 5 {
		sandglass(size, char, color)
	} else {
		fmt.Println("Слишком маленькая высота")
	}
}