package draw

import (
	"github.com/ssleert/sterm"
)

func (s *Info) writeTitle(x, y int, d *iconed) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	s.tui.WriteRune(' ')
	writeIcon(&s.tui, d[0], d[1], " ")
	s.tui.WriteString(sterm.Reset)
	s.tui.WriteString(d[2])
	s.tui.WriteRune(' ')
}

func (s *Info) titles() {
	var (
		cpuTitle  = iconed{s.colorMid, "", "CPU"}
		memTitle  = iconed{s.colorMid, "", "MEM"}
		swapTitle = iconed{s.colorMid, "", "SWAP"}
		diskTitle = iconed{s.colorMid, "", "DISK"}
		batTitle  = iconed{s.colorMid, "", "BAT"}
		infoTitle = iconed{s.colorMid, "", "INFO"}
	)

	s.writeTitle(6, s.y+2, &cpuTitle)
	s.writeTitle(6, s.y+11, &memTitle)
	s.writeTitle(38, s.y+11, &swapTitle)
	s.writeTitle(38, s.y+18, &diskTitle)
	s.writeTitle(38, s.y+23, &batTitle)
	s.writeTitle(38, s.y+26, &infoTitle)
}
