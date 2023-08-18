package main

import (
	"os"

	"github.com/01-edu/z01"
)

func sort(array []rune) []rune {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func charSort(arr1 []string) []rune {
	basestr := ""
	for _, el := range arr1 {
		basestr += el
	}
	base := []rune(basestr)
	return sort(base)
}

func sum(str string) int {
	base := []rune(str)
	if 
}

func main() {
	arguments := os.Args
	names := arguments[1:]

	name := perform(names)

	for i := range name {
		z01.PrintRune(name[i])
		z01.PrintRune('\n')
	}
}

// ./test_with_docker.sh sortparams sortparams/main.go --allow-builtin github.com/01-edu/z01.PrintRune os.Args
