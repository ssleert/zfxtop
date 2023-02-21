package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) swapStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(60, s.y+11, " GiB ")
	s.putStr(62, s.y+12, "GiB")
	s.putStr(45, s.y+12, "%")
	s.reset()
	s.putStr(38, s.y+12, "Used")
}

// draw dynamic info for bat block
func (s *Info) swapDynamic() {
	s.putStr(46, s.y+12, "               ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.SwapUsedPerc))
	s.putStr(44, s.y+12, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.SwapUsedPerc)))

	s.reset()
	s.putStr(59, s.y+11, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.SwapTotal)))
	s.putStr(60, s.y+12, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.SwapUsed)))

	s.drawGraph(38, s.y+16, 4, 27, s.DataDyn.Graph.SwapUsed[:])
}
