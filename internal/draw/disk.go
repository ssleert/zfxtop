package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) diskStatic() {
	var (
		rootText = iconed{s.colorList[0], "", "root"}
		homeText = iconed{s.colorList[1], "", "/home"}
		userText = iconed{s.colorList[2], "", "/usr"}
	)

	s.writeIconed(s.x+35, s.y+18, &rootText)
	s.writeIconed(s.x+35, s.y+19, &homeText)
	s.writeIconed(s.x+35, s.y+20, &userText)

	s.tui.WriteString(s.colorMid)
	s.putStr(s.x+45, s.y+18, "%")
	s.putStr(s.x+45, s.y+19, "%")
	s.putStr(s.x+45, s.y+20, "%")

	s.putStr(s.x+57, s.y+17, " GiB ")
	s.putStr(s.x+59, s.y+18, "GiB")
	s.putStr(s.x+59, s.y+19, "GiB")
	s.putStr(s.x+59, s.y+20, "GiB")

	s.reset()
	s.putStr(s.x+56, s.y+17, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.DiskTotal)))
}

// draw dynamic info for bat block
func (s *Info) diskDynamic() {
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.RootUsedPerc))
	s.putStr(s.x+44, s.y+18, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.RootUsedPerc)))
	s.reset()
	s.putStr(s.x+57, s.y+18, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.RootUsed)))

	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.HomeUsedPerc))
	s.putStr(s.x+44, s.y+19, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.HomeUsedPerc)))
	s.reset()
	s.putStr(s.x+57, s.y+19, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.HomeUsed)))

	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.UsrUsedPerc))
	s.putStr(s.x+44, s.y+20, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.UsrUsedPerc)))
	s.reset()
	s.putStr(s.x+57, s.y+20, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.UsrUsed)))
}
