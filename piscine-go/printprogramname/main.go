package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	name := arguments[0]

	runes := []rune(name)
	beg := 0
	for i := range runes {
		if runes[i] == '/' {
			beg = i + 1
		}
	}
	for ; beg < len(runes); beg++ {
		z01.PrintRune(runes[beg])
	}
	z01.PrintRune('\n')
}

// ./test_with_docker.sh printprogramname printprogramname/main.go --allow-builtin github.com/01-edu/z01.PrintRune os.Args
