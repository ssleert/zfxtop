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
	buf, _ := sterm.FrameArea(s.borders, 4, s.y+2, 66, s.y+10)
	s.tui.WriteString(buf)

	// mem frame
	buf, _ = sterm.FrameArea(s.borders, 4, s.y+11, 34, s.y+29)
	s.tui.WriteString(buf)
	s.separator(4, s.y+17, 31)
	s.separator(4, s.y+23, 31)

	// swap frame
	buf, _ = sterm.FrameArea(s.borders, 36, s.y+11, 66, s.y+17)
	s.tui.WriteString(buf)

	// disk frame
	buf, _ = sterm.FrameArea(s.borders, 36, s.y+18, 66, s.y+22)
	s.tui.WriteString(buf)

	// battery frame
	buf, _ = sterm.FrameArea(s.borders, 36, s.y+23, 66, s.y+25)
	s.tui.WriteString(buf)

	// info frame
	buf, _ = sterm.FrameArea(s.borders, 36, s.y+26, 66, s.y+29)
	s.tui.WriteString(buf)

	// buttons
	s.reset()
	s.tui.WriteString(
		fmt.Sprintf(
			"%s%sQ%s - quit | %sR%s - redraw",
			sterm.CursorTo(26, s.y+30),
			s.colorMid,
			s.colorFaint,
			s.colorMid,
			s.colorFaint,
		),
	)
}
