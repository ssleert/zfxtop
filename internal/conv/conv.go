package conv

import (
	"github.com/ssleert/sterm"
)

func SColorsFg(c [6]int) (cv [6]string) {
	for i, e := range c {
		cv[i] = sterm.Color256Fg(e)
	}
	return
}

func SColorsBg(c [6]int) (cv [6]string) {
	for i, e := range c {
		cv[i] = sterm.Color256Bg(e)
	}
	return
}

func TColorsFg(c [3]int) (cv [3]string) {
	for i, e := range c {
		cv[i] = sterm.Color256Fg(e)
	}
	return
}
