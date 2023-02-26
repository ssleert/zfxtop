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
	fmt.Printf("\n%s          ____     __             %s╭──────────────────╮\n", colors[0], colors[6])
	fmt.Printf("%s   ____  / __/  __/ /_____  ____  %s│%s  author: %sssleert %s│\n", colors[1], colors[6], colors[4], sterm.Reset, colors[6])
	fmt.Printf("%s  /_  / / /_|'|/_/ __/ __ \\/ __ \\ %s│%s    lang: %sgo      %s│\n", colors[2], colors[6], colors[3], sterm.Reset, colors[6])
	fmt.Printf("%s   / /_/ __/>  </ /_/ /_/ / /_/ / %s│%s version: %s%s   %s│\n", colors[3], colors[6], colors[2], sterm.Reset, self.Version, colors[6])
	fmt.Printf("%s  /___/_/ /_/|_|\\__/\\____/ .___/  %s│%s  commit: %s%s %s│\n", colors[4], colors[6], colors[1], sterm.Reset, self.Commit, colors[6])
	fmt.Printf("%s                        /_/       %s╰──────────────────╯%s\n\n", colors[5], colors[6], sterm.Reset)
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
	fmt.Printf("\n   %sz%sf%sx%st%so%sp %s-%s fetch top for gen Z with X\n", colors[5], colors[4], colors[3], colors[2], colors[1], colors[0], colors[7], sterm.Reset)
	fmt.Printf("%s  ╭────────────────────────────────────╮\n", colors[6])
	fmt.Printf("  │ %sarg          %s│ %sdesc                %s│\n", colors[7], colors[6], colors[7], colors[6])
	fmt.Printf("  │──────────────┼─────────────────────│\n")
	fmt.Printf("  │ %s-%sc %s--%sconfig  %s│ %sset config location %s│\n", colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6])
	fmt.Printf("  │ %s-%sl %s--%sclear   %s│ %sclear cache         %s│\n", colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6])
	fmt.Printf("  │ %s-%sv %s--%sversion %s│ %sget version info    %s│\n", colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6])
	fmt.Printf("  │ %s-%sh %s--%shelp    %s│ %sget help info       %s│\n", colors[7], sterm.Reset, colors[7], sterm.Reset, colors[6], sterm.Reset, colors[6])
	fmt.Printf("  ╰────────────────────────────────────╯%s\n\n", sterm.Reset)
}
