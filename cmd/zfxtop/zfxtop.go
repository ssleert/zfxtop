package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ssleert/sterm"
	"github.com/ssleert/zfxtop/internal/conf"
	"github.com/ssleert/zfxtop/internal/conv"
	"github.com/ssleert/zfxtop/internal/data"
	"github.com/ssleert/zfxtop/internal/draw"
	"github.com/ssleert/zfxtop/internal/msg"
	"github.com/ssleert/zfxtop/internal/self"
	"os"
	"runtime"
	"time"
)

var (
	config  string
	version bool
	test    bool

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
	flag.StringVar(&config, "c", "", "set config location")
	flag.BoolVar(&test, "t", false, "test flag")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(100)

	if version {
		versionFunc()
		return
	}

	if test {
		fmt.Println("asd")
		return
	}

	// parse config
	err := conf.Parse(config)
	if err != nil && !errors.Is(err, conf.ErrNoConfFiles) {
		msg.ExitMsg(err)
	}

	// start draw module with args from config
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

	d := data.Start()

	datadyn := &s.DataDyn
	datastat := &s.DataStat

	exitFromDraw := func(err error) {
		s.Stop()
		fmt.Println(err)
		os.Exit(1)
	}

	updateAll := func() {
		err := data.Update(datastat)
		if err != nil {
			exitFromDraw(err)
		}
		err = d.Update(datadyn)
		if err != nil {
			exitFromDraw(err)
		}
		buf := s.Static()
		fmt.Print(buf)
		buf = s.Dynamic()
		fmt.Print(buf)
	}

	updateDyn := func() {
		err := d.Update(datadyn)
		if err != nil {
			exitFromDraw(err)
		}
		buf := s.Dynamic()
		fmt.Print(buf)
	}

	// handle user input
	go func() {
		for {
			ch, err := sterm.GetChar()
			if err != nil {
				panic(err)
			}
			switch ch {
			case 'q', 'Q', 'й', 'Й':
				s.Stop()
				os.Exit(0)
			case 'r', 'R', 'к', 'К':
				fmt.Print(s.Redraw())
			default:
				continue
			}
		}
	}()

	// draw first data
	updateAll()

	ticker := time.NewTicker(conf.Update)

	// main loop
	for range ticker.C {
		updateDyn()
	}
}
