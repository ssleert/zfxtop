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
		cpuTitle  = iconed{s.colorMid, "󰍛", "CPU"}
		memTitle  = iconed{s.colorMid, "", "MEM"}
		swapTitle = iconed{s.colorMid, "󰋊", "SWAP"}
		diskTitle = iconed{s.colorMid, "", "DISK"}
		batTitle  = iconed{s.colorMid, "", "BAT"}
		infoTitle = iconed{s.colorMid, "󰋼", "INFO"}
	)

	s.writeTitle(s.x+3, s.y+1, &cpuTitle)
	s.writeTitle(s.x+3, s.y+10, &memTitle)
	s.writeTitle(s.x+35, s.y+10, &swapTitle)
	s.writeTitle(s.x+35, s.y+17, &diskTitle)
	s.writeTitle(s.x+35, s.y+22, &batTitle)
	s.writeTitle(s.x+35, s.y+25, &infoTitle)
}
