package draw

import (
	"github.com/ssleert/sterm"
	"math"
	"strconv"
)

// TODO rewrite without hard code))
func (s *Info) batPercLine(x, y, l, p int) {
	s.tui.WriteString(sterm.CursorTo(x, y))

	s.tui.WriteString(s.colorFaint)

	lngt := float64(l)
	perc := float64(p)
	for n := 0.0; n < lngt; n++ {
		switch {
		case n >= math.Round((lngt/6))*5 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[5])
		case n >= math.Round((lngt/6))*4 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[4])
		case n >= math.Round((lngt/6))*3 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[3])
		case n >= math.Round((lngt/6))*2 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[2])
		case n >= math.Round((lngt/6))*1 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[1])
		case n <= math.Round((lngt/6))*1 && n <= math.Round(perc/100*lngt):
			s.tui.WriteString(s.colorTempr[0])
		default:
			s.tui.WriteString(s.colorFaint)
		}
		s.tui.WriteRune('■')
	}
}

// draw static info for bat block
func (s *Info) batStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(64, s.y+24, "%")
}

// draw dynamic info for bat block
func (s *Info) batDynamic() {
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Bat.Perc))
	s.putStr(63, s.y+24, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Bat.Perc)))

	switch {
	case s.DataDyn.Bat.Perc > 99:
		s.batPercLine(38, s.y+24, 22, s.DataDyn.Bat.Perc)
	case s.DataDyn.Bat.Perc > 9:
		s.batPercLine(38, s.y+24, 23, s.DataDyn.Bat.Perc)
	default:
		s.batPercLine(38, s.y+24, 21, s.DataDyn.Bat.Perc)
	}
}
