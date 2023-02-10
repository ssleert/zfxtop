package draw

// draw static info for bat block
func (s *Info) swapStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(60, s.y+11, " GiB ")
	s.putStr(62, s.y+12, "GiB")
	s.reset()
	s.putStr(38, s.y+12, "Free")
}

// draw dynamic info for bat block
func (s *Info) swapDynamic() {
	// no implemented
	return
}
