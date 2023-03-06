package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ssleert/sterm"
	"github.com/ssleert/zfxtop/internal/arts"
	"github.com/ssleert/zfxtop/internal/conf"
	"github.com/ssleert/zfxtop/internal/conv"
	"github.com/ssleert/zfxtop/internal/data"
	"github.com/ssleert/zfxtop/internal/draw"
	"github.com/ssleert/zfxtop/internal/input"
	"github.com/ssleert/zfxtop/internal/msg"
	"os"
	"time"
)

var (
	version    bool
	config     string
	clearCache bool
	center     bool
)

func init() {
	// version
	flag.BoolVar(&version, "v", false, "get version info")
	flag.BoolVar(&version, "version", false, "get version info")

	// config
	flag.StringVar(&config, "c", "", "set config location")
	flag.StringVar(&config, "config", "", "set config location")

	// clear
	flag.BoolVar(&clearCache, "l", false, "clear cache")
	flag.BoolVar(&clearCache, "clear", false, "clear cache")

	// center tui
	flag.BoolVar(&center, "t", false, "center tui")
	flag.BoolVar(&center, "center", false, "center tui")

	flag.Usage = arts.HelpFunc
	flag.Parse()
}

func main() {
	if version {
		arts.VersionFunc()
		return
	}
	if clearCache {
		fmt.Println("not implemented")
		fmt.Println("dont worry the cache hasnt been implemented yet either")
		return
	}

	// parse config
	err := conf.Parse(config)
	if err != nil && !errors.Is(err, conf.ErrNoConfFiles) {
		msg.ExitMsg(err)
	}

	// set conf value true
	// if flag value true
	// ))
	if center {
		conf.Center = true
	}

	// start draw module with args from config
	s, err := draw.Start(
		conf.Icons,
		conf.Borders,
		conf.Center,
		sterm.Color256Fg(conf.ColorMid),
		sterm.Color256Fg(conf.ColorFaint),
		conv.SColorsFg(conf.ColorLoad),
		conv.SColorsFg(conf.ColorTempr),
		conv.TColorsFg(conf.ColorList),
	)
	if err != nil {
		msg.ExitMsg(err)
	}

	// start data collection
	d := data.Start()
	datadyn := &s.DataDyn
	datastat := &s.DataStat

	// first screen update
	err = data.Update(datastat)
	if err != nil {
		s.Stop()
		fmt.Println(err)
		os.Exit(1)

	}
	fmt.Print(s.Static())
	fmt.Print(s.Redraw())

	// usr input
	inputch := make(chan rune)
	stopch := make(chan bool) // stop msg for input goroutine
	go input.Handle(inputch, stopch)

	// main loop
	ticker := time.NewTicker(conf.Update)
	for {
		select {
		case <-ticker.C:
			err := d.Update(datadyn)
			if err != nil {
				s.Stop()
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Print(s.Dynamic())
		case ch := <-inputch:
			switch ch {
			case 'q', 'Q', 'й', 'Й':
				s.Stop()
				stopch <- true
				return
			case 'r', 'R', 'к', 'К':
				fmt.Print(s.Redraw())
			}
			stopch <- false
		}
	}
}
