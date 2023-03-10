package input

import (
	"github.com/ssleert/sterm"
)

func Handle(input chan rune, stop chan bool) {
	for {
		ch, err := sterm.GetChar()
		if err != nil {
			panic(err)
		}

		input <- ch
		if <-stop {
			return
		}
	}
}
