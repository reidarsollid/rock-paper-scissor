package main

import (
	"fmt"
	"os"
)

func main() {
	s, b := Play(os.Args[1])
	fmt.Printf("%s %s", s, b)
	fmt.Println()
}
