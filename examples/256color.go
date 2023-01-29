package main

import (
	"fmt"
	"github.com/ssleert/sterm"
)

func main() {
	for i := 0; i < 256; i++ {
		fmt.Print(" ", i, " ")
		sterm.Color256Bg(i)
		fmt.Print("                                 ")
		fmt.Print(sterm.Reset)
		fmt.Print("\n")
	}
}
