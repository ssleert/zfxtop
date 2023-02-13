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
	colorMid   string
	colorFaint string
	colorLoad  [6]string
	colorTempr [6]string
	colorList  [3]string
}

const (
	Size = 31
)

func (s *Info) putStr(x, y int, str string) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	s.tui.WriteString(str)
}

func (s *Info) writeIconed(x, y int, d *iconed) {
	s.tui.WriteString(sterm.CursorTo(x, y))
	writeIcon(&s.tui, d[0], d[1], sterm.Reset, " ")
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
	cm, cf string,
	cl, ct [6]string,
	cli [3]string,
) (*Info, error) {
	fmt.Print(sterm.ReserveArea(Size))
	fmt.Print(sterm.CursorHide())
	x, y, err := sterm.CursorPos()
	if err != nil {
		return nil, err
	}
	s, err := sterm.GetState()
	info := Info{
		x:          x - 1,
		y:          y - 1,
		s:          s,
		icons:      ico,
		borders:    brd,
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
	fmt.Print(sterm.CursorTo(1, s.y+1))
	fmt.Print(sterm.ClearScreenDown())
	fmt.Print(sterm.Reset)
	sterm.Restore(s.s)
}

func (s *Info) Redraw() string {
	var buf strings.Builder
	buf.WriteString(sterm.ClearScreenDown())
	buf.WriteString(s.Static())
	buf.WriteString(s.Dynamic())
	return buf.String()
}

// draw static info
// executes on program start
// or on redraw
func (s *Info) Static() string {
	s.tui.Reset()
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

	return s.tui.String()
}
