package main

import (
	"flag"
	"fmt"
	"github.com/ssleert/sterm"
	"github.com/ssleert/zfxtop/internal/conf"
	"github.com/ssleert/zfxtop/internal/conv"
	"github.com/ssleert/zfxtop/internal/draw"
	"github.com/ssleert/zfxtop/internal/msg"
	"github.com/ssleert/zfxtop/internal/self"
	"os"
)

var (
	config  string
	version bool

	versionFunc = func() {
		logoColors := [...]string{
			sterm.Color256Fg(conf.ColorLoad[5]),
			sterm.Color256Fg(conf.ColorLoad[4]),
			sterm.Color256Fg(conf.ColorLoad[3]),
			sterm.Color256Fg(conf.ColorLoad[2]),
			sterm.Color256Fg(conf.ColorLoad[1]),
			sterm.Color256Fg(conf.ColorLoad[0]),
			sterm.Color256Fg(conf.ColorFaint),
			sterm.Color256Fg(conf.ColorMid),
		}
		fmt.Printf("%s        ____     __             %s╭──────────────────╮\n", logoColors[0], logoColors[6])
		fmt.Printf("%s ____  / __/  __/ /_____  ____  %s│%s  author: %sssleert %s│\n", logoColors[1], logoColors[6], logoColors[4], sterm.Reset, logoColors[6])
		fmt.Printf("%s/_  / / /_|'|/_/ __/ __ \\/ __ \\ %s│%s    lang: %sgo      %s│\n", logoColors[2], logoColors[6], logoColors[3], sterm.Reset, logoColors[6])
		fmt.Printf("%s / /_/ __/>  </ /_/ /_/ / /_/ / %s│%s version: %s%s   %s│\n", logoColors[3], logoColors[6], logoColors[2], sterm.Reset, self.Version, logoColors[6])
		fmt.Printf("%s/___/_/ /_/|_|\\__/\\____/ .___/  %s│%s  commit: %s%s %s│\n", logoColors[4], logoColors[6], logoColors[1], sterm.Reset, self.Commit, logoColors[6])
		fmt.Printf("%s                      /_/       %s╰──────────────────╯%s\n", logoColors[5], logoColors[6], sterm.Reset)
	}
)

func init() {
	flag.BoolVar(&version, "v", false, "get version info")
	flag.StringVar(&config, "c", "", "get version info")
	flag.Parse()
}

func main() {
	if version {
		versionFunc()
		os.Exit(0)
	}

	err := conf.Parse(config)
	if err != nil {
		msg.ExitMsg(err)
	}

	s, err := draw.Start(
		conf.Icons,
		conf.Borders,
		sterm.Color256Fg(conf.ColorMid),
		sterm.Color256Fg(conf.ColorFaint),
		conv.SColorsBg(conf.ColorLoad),
		conv.SColorsFg(conf.ColorTempr),
		conv.TColorsFg(conf.ColorList),
	)
	if err != nil {
		msg.ExitMsg(err)
	}

	s.GetStaticData().Update()
	buf := s.Static()
	fmt.Println(buf)
	fmt.Println()
}
