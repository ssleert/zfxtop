package draw

import (
	"github.com/ssleert/sterm"
	"math"
	"strings"
)

var (
	blocks = [8]rune{
		'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█',
	}
)

func (s *Info) drawGraph(x, y, h, l int, in []int) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	for i := 0; i < h; i++ {
		s.tui.WriteString(strings.Repeat(" ", l))
		s.tui.WriteString(sterm.CursorUp(1))
		s.tui.WriteString(sterm.CursorLeft(l))
	}

	// 1 / 8
	hp := 0.125
	ds := float64(h) / hp
	// clds := int(math.Round(float64(h) / 6))
	pds := 100.0 / ds

	s.tui.WriteString(sterm.CursorTo(x, y))
	for _, e := range in {
		ost := math.Mod(float64(e), pds*8)
		col := int(math.Floor(float64(e) / (pds * 8)))

		for i := 0; i < col; i++ {
			s.tui.WriteRune(blocks[7])
			s.tui.WriteString(sterm.CursorUp(1))
			s.tui.WriteString(sterm.CursorLeft(1))
		}

		if ost != 0 {
			switch {
			case ost <= pds:
				s.tui.WriteRune(blocks[0])
			case ost <= pds*2:
				s.tui.WriteRune(blocks[1])
			case ost <= pds*3:
				s.tui.WriteRune(blocks[2])
			case ost <= pds*4:
				s.tui.WriteRune(blocks[3])
			case ost <= pds*5:
				s.tui.WriteRune(blocks[4])
			case ost <= pds*6:
				s.tui.WriteRune(blocks[5])
			case ost <= pds*7:
				s.tui.WriteRune(blocks[6])
			}
		}

		for i := 0; i < col; i++ {
			s.tui.WriteString(sterm.CursorDown(1))
		}
	}
}
