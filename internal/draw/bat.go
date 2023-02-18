package draw

import (
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) batStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(64, s.y+24, "%")
}

// draw dynamic info for bat block
func (s *Info) batDynamic() {
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Bat.Perc))
	s.putStr(63, s.y+24, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Bat.Perc)))
}
