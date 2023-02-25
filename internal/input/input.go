package input

import (
	"github.com/ssleert/sterm"
)

func Handle(input chan rune) {
	for {
		ch, err := sterm.GetChar()
		if err != nil {
			panic(err)
		}
		input <- ch
	}
}
