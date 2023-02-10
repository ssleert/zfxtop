package main

import (
	"fmt"
)

func main() {
	var bsd string
	for {
		var asd string
		fmt.Scanln(&asd)
		if asd == "break" {
			break
		}
		bsd += asd
	}
	fmt.Println(bsd)
	for {

	}
}
