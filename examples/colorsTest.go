package main

import (
	"fmt"
	"github.com/ssleert/sterm"
)

var (
	// #444444
	colorReset = sterm.Reset
	colorFaint = fmt.Sprintf("\x1b[38;5;%dm", 238)
	colorLoad  = [...]string{
		"\x1b[38;5;27m",
		"\x1b[38;5;63m",
		"\x1b[38;5;99m",
		"\x1b[38;5;135m",
		"\x1b[38;5;171m",
		"\x1b[38;5;207m",
	}
	colorTempr = [...]string{
		"\x1b[38;5;49m",
		"\x1b[38;5;79m",
		"\x1b[38;5;109m",
		"\x1b[38;5;139m",
		"\x1b[38;5;169m",
		"\x1b[38;5;199m",
	}
	colorList = [...]string{
		"\x1b[38;5;109m",
		"\x1b[38;5;79m",
		"\x1b[38;5;169m",
	}
	colorMid = "\x1b[38;5;245m"
)

func main() {
	fmt.Println()
	fmt.Println("   " + colorFaint + "╭─" + colorReset + " " + colorMid + "" + colorReset + " CPU " + colorFaint + "─────────────────" + colorReset + " 19:47:27 " + colorFaint + "────────────────" + colorReset + " E5-2660 " + colorFaint + "─╮" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╭─────────────╮" + colorFaint + colorReset + "                                             " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[2] + "" + colorReset + " load  " + colorTempr[4] + "99" + colorMid + "% " + colorFaint + "│" + colorFaint + colorReset + "                         " + colorLoad[5] + "▁▂▃▄▅▆▇████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[1] + "龍" + colorReset + "fr " + colorTempr[1] + "3.3" + colorMid + "ghz " + colorFaint + "│" + colorFaint + colorReset + "                 " + colorLoad[4] + "▁▂▃▄▅▆▇████████████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[0] + " " + colorReset + "temp " + colorTempr[2] + "63" + colorMid + "°C " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[3] + "▁▂▃▄▅▆▇████████████████████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╰─────────────╯" + colorFaint + colorReset + " " + colorLoad[2] + "▁▂▃▄▅▆▇████████████████████████████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[1] + "▁▂▃▄▅▆▇████████████████████████████████████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorLoad[0] + "▁▂▃▄▅▆▇████████████████████████████████████████████████████" + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "╰─────────────────────────────────────────────────────────────╯" + colorReset)
	fmt.Println("   " + colorFaint + "╭─" + colorReset + " " + colorMid + "" + colorReset + " MEM " + colorFaint + "─────────" + colorReset + " 15.55 " + colorMid + "GiB " + colorFaint + "─╮" + colorReset + " " + colorFaint + "╭─" + colorReset + " " + colorMid + "" + colorReset + " SWAP " + colorFaint + "──────────" + colorReset + " 4.0 " + colorMid + "GiB " + colorFaint + "─╮" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " Used               4.46 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " Free               1.23 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                             " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "                             " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                 " + colorLoad[5] + "▁▂▃▄▅▆▇████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "                 " + colorLoad[5] + "▁▂▃▄▅▆▇████ " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[2] + "▁▂▃▄▅▆▇████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[2] + "▁▂▃▄▅▆▇████████████ " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorLoad[0] + "▁▂▃▄▅▆▇████████████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorLoad[0] + "▁▂▃▄▅▆▇████████████████████ " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "├─────────────────────────────┤" + colorFaint + colorReset + " " + colorFaint + "╰─────────────────────────────╯" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " Availible          11.0 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╭─" + colorReset + " " + colorMid + "" + colorReset + " DISK " + colorFaint + "──────────" + colorReset + " 223 " + colorMid + "GiB " + colorFaint + "─╮" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                             " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[0] + "" + colorReset + " root  " + colorTempr[4] + "66" + colorMid + "%       " + colorReset + "140.9 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                 " + colorLoad[5] + "▁▂▃▄▅▆▇████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[1] + "" + colorReset + " /home " + colorTempr[3] + "54" + colorMid + "%       " + colorReset + "114.4 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[2] + "▁▂▃▄▅▆▇████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[2] + "" + colorReset + " /usr  " + colorTempr[0] + "10" + colorMid + "%       " + colorReset + "10.2  " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorLoad[0] + "▁▂▃▄▅▆▇████████████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╰─────────────────────────────╯" + colorReset)
	fmt.Println("   " + colorFaint + "├─────────────────────────────┤" + colorFaint + colorReset + " " + colorFaint + "╭─" + colorReset + "  BAT " + colorFaint + "──────────────" + colorReset + " 3:45 " + colorFaint + "─╮" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " Free               4.10 " + colorMid + "GiB " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorTempr[0] + "■■■■" + colorTempr[1] + "■■■■" + colorTempr[2] + "■■■■" + colorTempr[3] + "■■■■" + colorTempr[4] + "■■■■" + colorTempr[5] + "■■■" + colorReset + " 51% " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                             " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╰─────────────────────────────╯" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "                 " + colorLoad[5] + "▁▂▃▄▅▆▇████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "╭─" + colorReset + " " + colorMid + "" + colorReset + " INFO " + colorFaint + "────────────────────╮" + colorReset)
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + "         " + colorLoad[2] + "▁▂▃▄▅▆▇████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[2] + "" + colorReset + " Distro         Arch Linux " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println("   " + colorFaint + "│" + colorFaint + colorReset + " " + colorLoad[0] + "▁▂▃▄▅▆▇████████████████████ " + colorFaint + "│" + colorFaint + colorReset + " " + colorFaint + "│" + colorFaint + colorReset + " " + colorList[1] + "" + colorReset + " Hostname            sfome " + colorFaint + "│" + colorFaint + colorReset + "")
	fmt.Println(colorFaint + "   ╰─────────────────────────────╯ ╰─────────────────────────────╯" + colorReset)
	fmt.Println("                         " + colorMid + "Q" + colorFaint + " - quit | " + colorMid + "R" + colorFaint + " - redraw" + colorReset)
	fmt.Println()
}
