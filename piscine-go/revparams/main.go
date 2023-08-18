package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	name := arguments[1:]
	// fmt.Println(len(name))
	// z01.PrintRune('\n')
	for i := len(name) - 1; i >= 0; i-- {
		runes := []rune(name[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}

// ./test_with_docker.sh revparams revparams/main.go --allow-builtin github.com/01-edu/z01.PrintRune os.Args
