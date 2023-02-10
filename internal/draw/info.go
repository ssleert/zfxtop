package draw

import (
	"github.com/ssleert/sterm"
)

// draw static info for info block
func (s *Info) infoStatic() {
	var (
		distroName = iconed{s.colorList[2], "", "Distro"}
		hostName   = iconed{s.colorList[1], "", "Hostname"}
	)

	s.writeIconed(38, s.y+27, &distroName)
	s.writeIconed(38, s.y+28, &hostName)

	s.putStr(64, s.y+27, sterm.RevPrint(s.DataStat.DistroName))
	s.putStr(64, s.y+28, sterm.RevPrint(s.DataStat.HostName))
}
