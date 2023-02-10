package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
)

// draw static info for bat block
func (s *Info) cpuStatic() {
	s.reset()
	s.putStr(64, s.y+2, sterm.RevPrint(fmt.Sprintf(" %s ", s.DataStat.CpuName)))
}

// draw dynamic info for bat block
func (s *Info) cpuDynamic() {
	s.reset()
	s.tui.WriteString(sterm.CursorTo(30, s.y+2))
	s.tui.WriteRune(' ')
	s.tui.WriteString(s.DataDyn.Time.Format("15:04:05"))
	s.tui.WriteRune(' ')
	return
}
