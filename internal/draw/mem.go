package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) memStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(28, s.y+11, " GiB ")
	s.putStr(30, s.y+12, "GiB")
	s.putStr(13, s.y+12, "%")
	s.putStr(30, s.y+18, "GiB")
	s.putStr(18, s.y+18, "%")
	s.putStr(30, s.y+24, "GiB")
	s.putStr(13, s.y+24, "%")

	s.reset()

	s.putStr(6, s.y+12, "Used")
	s.putStr(6, s.y+18, "Available")
	s.putStr(6, s.y+24, "Free")

	s.putStr(27, s.y+11, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.MemTotal)))
	s.putStr(27, s.y+11, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.MemTotal)))
}

// draw dynamic info for bat block
func (s *Info) memDynamic() {
	s.putStr(14, s.y+12, "                ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.UsedPerc))
	s.putStr(12, s.y+12, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.UsedPerc)))
	s.reset()
	s.putStr(28, s.y+12, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Used)))

	s.putStr(19, s.y+18, "           ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.AvailablePerc))
	s.putStr(17, s.y+18, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.AvailablePerc)))
	s.reset()
	s.putStr(28, s.y+18, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Available)))

	s.putStr(14, s.y+24, "                ")
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Mem.FreePerc))
	s.putStr(12, s.y+24, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Mem.FreePerc)))
	s.reset()
	s.putStr(28, s.y+24, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Mem.Free)))
}
