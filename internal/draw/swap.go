package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) swapStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(s.x+57, s.y+10, " GiB ")
	s.putStr(s.x+59, s.y+11, "GiB")
	s.putStr(s.x+42, s.y+11, "%")
	s.reset()
	s.putStr(s.x+35, s.y+11, "Used")
}

// draw dynamic info for bat block
func (s *Info) swapDynamic() {
	s.putStr(s.x+43, s.y+11, "               ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.SwapUsedPerc))
	s.putStr(s.x+41, s.y+11, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.SwapUsedPerc)))

	s.reset()
	s.putStr(s.x+56, s.y+10, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.SwapTotal)))
	s.putStr(s.x+57, s.y+11, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.SwapUsed)))

	s.drawGraph(s.x+35, s.y+15, 4, 27, s.DataDyn.Graph.SwapUsed[:])
}
