package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"github.com/ssleert/zfxtop/internal/conf"
	"github.com/ssleert/zfxtop/internal/data"
	"strings"
)

type iconed [3]string

type Info struct {
	x        int
	y        int
	s        sterm.State
	tui      strings.Builder
	DataDyn  data.Dynamic
	DataStat data.Static

	// from user
	icons      bool
	borders    sterm.Borders
	center     bool // center tui
	colorMid   string
	colorFaint string
	colorLoad  [6]string
	colorTempr [6]string
	colorList  [3]string
}

const (
	SizeY = 30
	SizeX = 66

	SizeYCenter = 29
	SizeXCenter = 63
)

func (s *Info) putStr(x, y int, str string) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	s.tui.WriteString(str)
}

func (s *Info) writeIconed(x, y int, d *iconed) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	writeIcon(&s.tui, d[0], d[1], " ")
	s.tui.WriteString(sterm.Reset)
	s.tui.WriteString(d[2])
}

func writeIcon(s *strings.Builder, i ...string) {
	if conf.Icons {
		for _, e := range i {
			s.WriteString(e)
		}
	}
}

func (s *Info) reset() {
	s.tui.WriteString(sterm.Reset)
}

func Start(
	ico bool,
	brd sterm.Borders,
	c bool,
	cm, cf string,
	cl, ct [6]string,
	cli [3]string,
) (*Info, error) {
	tsx, tsy, err := sterm.Size()
	if err != nil {
		return nil, err
	}
	if tsx < SizeX || tsy < SizeY {
		return nil, &TermSizeTooLittle{tsx, tsy}
	}

	fmt.Print(sterm.CursorHide())

	var x, y int
	if !c {
		fmt.Print(sterm.ReserveArea(SizeY))
		x, y, err = sterm.CursorPos()
		if err != nil {
			return nil, err
		}

		x += 3 - 1
		y += 1 - 1
	} else {
		fmt.Print(sterm.SaveAttrs())
	}

	s, err := sterm.GetState()
	if err != nil {
		return nil, err
	}

	info := Info{
		x:          x,
		y:          y,
		s:          s,
		icons:      ico,
		borders:    brd,
		center:     c,
		colorMid:   cm,
		colorFaint: cf,
		colorLoad:  cl,
		colorTempr: ct,
		colorList:  cli,
	}
	return &info, nil
}

// stop tui drawing
func (s *Info) Stop() {
	fmt.Print(sterm.CursorShow())
	if !s.center {
		fmt.Print(sterm.CursorTo(1, s.y))
		fmt.Print(sterm.ClearScreenDown())
	} else {
		fmt.Print(sterm.ClearEntireScreen())
		fmt.Print(sterm.CursorTo(1, 1))
	}

	fmt.Print(sterm.Reset)
	sterm.Restore(s.s)
}

func (s *Info) Redraw() string {
	var buf strings.Builder
	buf.WriteString(sterm.CursorTo(1, s.y+1))
	buf.WriteString(sterm.ClearScreenDown())
	buf.WriteString(s.Static())
	buf.WriteString(s.Dynamic())
	return buf.String()
}

func (s *Info) GetCursorDrawPos() (int, int) {
	return 0, 0
}

// draw static info
// executes on program start
// or on redraw
func (s *Info) Static() string {
	s.tui.Reset()

	if s.center {
		x, y, _ := sterm.Size()
		s.x = x/2 - SizeXCenter/2
		s.y = y/2 - SizeYCenter/2
		s.tui.WriteString(sterm.ClearEntireScreen())
	}

	s.frames()
	s.titles()
	s.infoStatic()
	s.batStatic()
	s.diskStatic()
	s.swapStatic()
	s.cpuStatic()
	s.memStatic()

	return s.tui.String()
}

// draw dinamic info
// like cpu/ram/swap graph
// executes on every tick
func (s *Info) Dynamic() string {
	s.tui.Reset()
	s.cpuDynamic()
	s.memDynamic()
	s.swapDynamic()
	s.diskDynamic()
	s.batDynamic()

	return s.tui.String()
}
