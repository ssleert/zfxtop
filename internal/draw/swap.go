package draw

// draw static info for bat block
func (s *Info) swapStatic() {
	s.tui.WriteString(s.colorMid)
	s.putStr(64, s.y+24, "%")
	return
}

// draw dynamic info for bat block
func (s *Info) swapDynamic() {
	// no implemented
	return
}
