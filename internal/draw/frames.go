package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
)

func (s *Info) separator(x, y, l int) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	s.tui.WriteRune(s.borders[6])
	for i := 0; i < l-2; i++ {
		s.tui.WriteRune(s.borders[1])
	}
	s.tui.WriteRune(s.borders[7])
}

func (s *Info) frames() {
	s.tui.WriteString(s.colorFaint)

	// cpu frame
	buf, _ := sterm.FrameArea(s.borders, s.x+1, s.y+1, s.x+63, s.y+9)
	s.tui.WriteString(buf)

	// mem frame
	buf, _ = sterm.FrameArea(s.borders, s.x+1, s.y+10, s.x+31, s.y+28)
	s.tui.WriteString(buf)
	s.separator(s.x+1, s.y+16, 31)
	s.separator(s.x+1, s.y+22, 31)

	// swap frame
	buf, _ = sterm.FrameArea(s.borders, s.x+33, s.y+10, s.x+63, s.y+16)
	s.tui.WriteString(buf)

	// disk frame
	buf, _ = sterm.FrameArea(s.borders, s.x+33, s.y+17, s.x+63, s.y+21)
	s.tui.WriteString(buf)

	// battery frame
	buf, _ = sterm.FrameArea(s.borders, s.x+33, s.y+22, s.x+63, s.y+24)
	s.tui.WriteString(buf)

	// info frame
	buf, _ = sterm.FrameArea(s.borders, s.x+33, s.y+25, s.x+63, s.y+28)
	s.tui.WriteString(buf)

	// buttons
	s.reset()
	s.tui.WriteString(
		fmt.Sprintf(
			"%s%sQ%s - quit | %sR%s - redraw",
			sterm.CursorTo(s.x+23, s.y+29),
			s.colorMid,
			s.colorFaint,
			s.colorMid,
			s.colorFaint,
		),
	)
}
