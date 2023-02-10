package draw

// import (
// "fmt"
// "github.com/ssleert/sterm"
// )

// draw static info for bat block
func (s *Info) memStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(28, s.y+11, " GiB ")
	s.putStr(30, s.y+12, "GiB")
	s.putStr(30, s.y+18, "GiB")
	s.putStr(30, s.y+24, "GiB")
	s.reset()
	s.putStr(6, s.y+12, "Used")
	s.putStr(6, s.y+18, "Available")
	s.putStr(6, s.y+24, "Free")

}

// draw dynamic info for bat block
func (s *Info) memDynamic() {
	// no implemented
	return
}
