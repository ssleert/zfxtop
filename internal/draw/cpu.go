package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) cpuStatic() {
	s.reset()
	s.putStr(s.x+61, s.y+1, sterm.RevPrint(" "+s.DataStat.CpuName+" "))
}

// draw dynamic info for bat block
func (s *Info) cpuDynamic() {
	var (
		load = iconed{s.colorList[2], "", "load"}
		freq = iconed{s.colorList[1], "󰓅", "fr"}
		temp = iconed{s.colorList[0], "", "temp"}
	)

	s.reset()
	s.tui.WriteString(sterm.CursorTo(s.x+27, s.y+1))
	s.tui.WriteRune(' ')
	s.tui.WriteString(s.DataDyn.Time.Format("15:04:05"))
	s.tui.WriteRune(' ')

	s.drawGraph(s.x+3, s.y+8, 7, 59, s.DataDyn.Graph.CpuLoad[:])
	s.tui.WriteString(s.colorFaint)
	buf, _ := sterm.CharArea(' ', s.x+3, s.y+2, s.x+17, s.y+6)
	s.tui.WriteString(buf)
	buf, _ = sterm.FrameArea(s.borders, s.x+3, s.y+2, s.x+17, s.y+6)
	s.tui.WriteString(buf)

	s.writeIconed(s.x+5, s.y+3, &load)
	s.writeIconed(s.x+5, s.y+4, &freq)
	s.writeIconed(s.x+5, s.y+5, &temp)

	s.tui.WriteString(sterm.CursorTo(s.x+15, s.y+3))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteRune('%')
	s.tui.WriteString(sterm.CursorLeft(2))
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.CpuLoad))
	s.tui.WriteString(sterm.RevPrint(" " + strconv.Itoa(s.DataDyn.CpuLoad)))

	s.tui.WriteString(sterm.CursorTo(s.x+13, s.y+4))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteString("ghz")
	s.tui.WriteString(sterm.CursorLeft(4))
	s.tui.WriteString(colorForFreq(&s.colorTempr, s.DataDyn.CpuFreq))
	s.tui.WriteString(sterm.RevPrint(fmt.Sprintf("%.1f", s.DataDyn.CpuFreq)))

	s.tui.WriteString(sterm.CursorTo(s.x+14, s.y+5))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteString("°C")
	s.tui.WriteString(sterm.CursorLeft(3))
	s.tui.WriteString(colorForTemp(&s.colorTempr, s.DataDyn.CpuTemp))
	s.tui.WriteString(sterm.RevPrint(strconv.Itoa(s.DataDyn.CpuTemp)))
}
