package draw

import (
	"github.com/ssleert/sterm"
)

// draw static info for info block
func (s *Info) infoStatic() {
	var (
		distroName = iconed{s.colorList[2], "󰏖", "Distro"}
		hostName   = iconed{s.colorList[1], "", "Hostname"}
	)

	s.writeIconed(s.x+35, s.y+26, &distroName)
	s.writeIconed(s.x+35, s.y+27, &hostName)

	s.putStr(s.x+61, s.y+26, sterm.RevPrint(s.DataStat.DistroName))
	s.putStr(s.x+61, s.y+27, sterm.RevPrint(s.DataStat.HostName))
}
