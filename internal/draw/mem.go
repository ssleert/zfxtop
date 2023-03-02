package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) memStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(s.x+25, s.y+10, " GiB ")
	s.putStr(s.x+27, s.y+11, "GiB")
	s.putStr(s.x+10, s.y+11, "%")
	s.putStr(s.x+27, s.y+17, "GiB")
	s.putStr(s.x+15, s.y+17, "%")
	s.putStr(s.x+27, s.y+23, "GiB")
	s.putStr(s.x+10, s.y+23, "%")

	s.reset()

	s.putStr(s.x+3, s.y+11, "Used")
	s.putStr(s.x+3, s.y+17, "Available")
	s.putStr(s.x+3, s.y+23, "Free")

	s.putStr(s.x+25, s.y+10, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.MemTotal)))
	s.putStr(s.x+25, s.y+10, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.MemTotal)))
}

// draw dynamic info for bat block
func (s *Info) memDynamic() {
	s.putStr(s.x+11, s.y+11, "                ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.UsedPerc))
	s.putStr(s.x+9, s.y+11, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.UsedPerc)))
	s.reset()
	s.putStr(s.x+25, s.y+11, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Used)))

	s.putStr(s.x+16, s.y+17, "           ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.AvailablePerc))
	s.putStr(s.x+14, s.y+17, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.AvailablePerc)))
	s.reset()
	s.putStr(s.x+25, s.y+17, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Available)))

	s.putStr(s.x+11, s.y+23, "                ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.FreePerc))
	s.putStr(s.x+9, s.y+23, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.FreePerc)))
	s.reset()
	s.putStr(s.x+25, s.y+23, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Free)))

	s.drawGraph(s.x+3, s.y+15, 4, 27, s.DataDyn.Graph.MemUsed[:])
	s.drawGraph(s.x+3, s.y+21, 4, 27, s.DataDyn.Graph.MemAvailable[:])
	s.drawGraph(s.x+3, s.y+27, 4, 27, s.DataDyn.Graph.MemFree[:])
}
