package arts

import (
	"fmt"
	"github.com/ssleert/sterm"
	"github.com/ssleert/zfxtop/internal/conf"
	"github.com/ssleert/zfxtop/internal/self"
)

func VersionFunc() {
	colors := [...]string{
		sterm.Color256Fg(conf.ColorLoad[5]),
		sterm.Color256Fg(conf.ColorLoad[4]),
		sterm.Color256Fg(conf.ColorLoad[3]),
		sterm.Color256Fg(conf.ColorLoad[2]),
		sterm.Color256Fg(conf.ColorLoad[1]),
		sterm.Color256Fg(conf.ColorLoad[0]),
		sterm.Color256Fg(conf.ColorFaint),
		sterm.Color256Fg(conf.ColorMid),
	}

	fmt.Printf(
		`
%s          ____     __             %s╭──────────────────╮
%s   ____  / __/  __/ /_____  ____  %s│%s  author: %sssleert %s│
%s  /_  / / /_|'|/_/ __/ __ \/ __ \ %s│%s    lang: %sgo      %s│
%s   / /_/ __/>  </ /_/ /_/ / /_/ / %s│%s version: %s%s   %s│
%s  /___/_/ /_/|_|\__/\____/ .___/  %s│%s  commit: %s%s %s│
%s                        /_/       %s╰──────────────────╯%s

`,
		colors[0], colors[6],
		colors[1], colors[6], colors[4], sterm.Reset, colors[6],
		colors[2], colors[6], colors[3], sterm.Reset, colors[6],
		colors[3], colors[6], colors[2], sterm.Reset, self.Version, colors[6],
		colors[4], colors[6], colors[1], sterm.Reset, self.Commit, colors[6],
		colors[5], colors[6], sterm.Reset,
	)
}

func HelpFunc() {
	colors := [...]string{
		sterm.Color256Fg(conf.ColorLoad[5]),
		sterm.Color256Fg(conf.ColorLoad[4]),
		sterm.Color256Fg(conf.ColorLoad[3]),
		sterm.Color256Fg(conf.ColorLoad[2]),
		sterm.Color256Fg(conf.ColorLoad[1]),
		sterm.Color256Fg(conf.ColorLoad[0]),
		sterm.Color256Fg(conf.ColorFaint),
		sterm.Color256Fg(conf.ColorMid),
	}

	fmt.Printf(
		`

   %sz%sf%sx%st%so%sp %s-%s fetch top for gen Z with X
%s  ╭────────────────────────────────────╮
  │ %sarg          %s│ %sdesc                %s│
  │──────────────┼─────────────────────│
  │ %s-%sc %s--%sconfig  %s│ %sset config location %s│
  │ %s-%sl %s--%sclear   %s│ %sclear cache         %s│
  │ %s-%sv %s--%sversion %s│ %sget version info    %s│
  │ %s-%sh %s--%shelp    %s│ %sget help info       %s│
  ╰────────────────────────────────────╯%s

`,
		colors[5], colors[4], colors[3], colors[2], colors[1], colors[0], colors[7], sterm.Reset,
		colors[6],
		colors[7], colors[6], colors[7], colors[6],
		colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6],
		colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6],
		colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6],
		colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6],
		sterm.Reset,
	)
}
