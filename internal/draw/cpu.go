package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) cpuStatic() {
	s.reset()
	s.putStr(64, s.y+2, sterm.RevPrint(" "+s.DataStat.CpuName+" "))
}

// draw dynamic info for bat block
func (s *Info) cpuDynamic() {
	var (
		load = iconed{s.colorList[2], "", "load"}
		freq = iconed{s.colorList[1], "龍", "fr"}
		temp = iconed{s.colorList[0], "", "temp"}
	)

	s.reset()
	s.tui.WriteString(sterm.CursorTo(30, s.y+2))
	s.tui.WriteRune(' ')
	s.tui.WriteString(s.DataDyn.Time.Format("15:04:05"))
	s.tui.WriteRune(' ')

	s.tui.WriteString(s.colorFaint)
	buf, _ := sterm.CharArea(' ', 6, s.y+3, 64, s.y+9)
	s.tui.WriteString(buf)
	buf, _ = sterm.FrameArea(s.borders, 6, s.y+3, 20, s.y+7)
	s.tui.WriteString(buf)

	s.writeIconed(8, s.y+4, &load)

	// nerd fonts FUCK YOU
	// s.writeIconed(8, s.y+5, &freq)
	// '龍' icon weight is bigger than normal char

	s.tui.WriteString(sterm.CursorTo(8, s.y+5))
	if s.icons {
		s.tui.WriteString(freq[0])
		s.tui.WriteString(freq[1])
		s.reset()
	}
	s.tui.WriteString(freq[2])

	// -------------------------------------------

	s.writeIconed(8, s.y+6, &temp)

	s.tui.WriteString(sterm.CursorTo(18, s.y+4))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteRune('%')
	s.tui.WriteString(sterm.CursorLeft(2))
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.CpuLoad))
	s.tui.WriteString(sterm.RevPrint(" " + strconv.Itoa(s.DataDyn.CpuLoad)))

	s.tui.WriteString(sterm.CursorTo(16, s.y+5))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteString("ghz")
	s.tui.WriteString(sterm.CursorLeft(4))
	s.tui.WriteString(colorForFreq(&s.colorTempr, s.DataDyn.CpuFreq))
	s.tui.WriteString(sterm.RevPrint(fmt.Sprintf("%.1f", s.DataDyn.CpuFreq)))

	s.tui.WriteString(sterm.CursorTo(17, s.y+6))
	s.tui.WriteString(s.colorMid)
	s.tui.WriteString("°C")
	s.tui.WriteString(sterm.CursorLeft(3))
	s.tui.WriteString(colorForTemp(&s.colorTempr, s.DataDyn.CpuTemp))
	s.tui.WriteString(sterm.RevPrint(strconv.Itoa(s.DataDyn.CpuTemp)))
}
