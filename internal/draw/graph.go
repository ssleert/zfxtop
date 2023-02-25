package draw

import (
	// "fmt"
	"github.com/ssleert/sterm"
	"math"
	"strings"
)

var (
	blocks = [8]rune{
		'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█',
	}
)

func (s *Info) colorForBar(d float64, p int) {
	pf := float64(p)

	if d == 0.5 {
		switch {
		case pf <= d:
			s.tui.WriteString(s.colorLoad[0])
		case pf <= d*2:
			s.tui.WriteString(s.colorLoad[2])
		case pf <= d*3:
			s.tui.WriteString(s.colorLoad[2])
		case pf <= d*4:
			s.tui.WriteString(s.colorLoad[4])
		case pf <= d*5:
			s.tui.WriteString(s.colorLoad[4])
		case pf <= d*6:
			s.tui.WriteString(s.colorLoad[5])
		}
	} else {
		switch {
		case pf <= d:
			s.tui.WriteString(s.colorLoad[0])
		case pf <= d*2:
			s.tui.WriteString(s.colorLoad[1])
		case pf <= d*3:
			s.tui.WriteString(s.colorLoad[2])
		case pf <= d*4:
			s.tui.WriteString(s.colorLoad[3])
		case pf <= d*5:
			s.tui.WriteString(s.colorLoad[4])
		case pf <= d*6:
			s.tui.WriteString(s.colorLoad[5])
		}
	}

}

func (s *Info) drawGraph(x, y, h, l int, in []int) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	for i := 0; i < h; i++ {
		s.tui.WriteString(strings.Repeat(" ", l))
		s.tui.WriteString(sterm.CursorUp(1))
		s.tui.WriteString(sterm.CursorLeft(l))
	}
	s.tui.WriteString(sterm.CursorTo(x, y))

	// delimiter for colors
	clds := float64(h) / 8

	// 1 / 8
	hp := 0.125
	ds := float64(h) / hp
	pds := 100.0 / ds

	for _, e := range in[len(in)-l:] {
		if e < 0 {
			s.tui.WriteRune(' ')
			continue
		} else if e == 0 {
			e = 1
		} else if e > 100 {
			e = 100
		}

		ost := math.Mod(float64(e), pds*8)
		col := int(math.Floor(float64(e) / (pds * 8)))
		if e >= 100 {
			col--
		}

		var i int
		for ; i < col; i++ {
			s.colorForBar(clds, i)
			s.tui.WriteRune(blocks[7])
			s.tui.WriteString(sterm.CursorUp(1))
			s.tui.WriteString(sterm.CursorLeft(1))
		}

		s.colorForBar(clds, i)
		if ost > 0 {
			switch {
			case ost <= pds:
				s.tui.WriteRune(blocks[0])
			case ost <= pds*2:
				s.tui.WriteRune(blocks[1])
			case ost <= pds*3:
				s.tui.WriteRune(blocks[1])
			case ost <= pds*4:
				s.tui.WriteRune(blocks[2])
			case ost <= pds*5:
				s.tui.WriteRune(blocks[3])
			case ost <= pds*6:
				s.tui.WriteRune(blocks[4])
			case ost <= pds*7:
				s.tui.WriteRune(blocks[5])
			case ost <= pds*8:
				s.tui.WriteRune(blocks[6])
			}
		} else {
			s.tui.WriteRune(' ')
		}

		for i := 0; i < col; i++ {
			s.tui.WriteString(sterm.CursorDown(1))
		}
	}
}
