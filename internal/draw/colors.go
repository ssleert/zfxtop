package draw

// return color for percent
func colorForPercent(c *[6]string, p int) string {
	switch {
	case p > 80:
		return c[5]
	case p > 68:
		return c[4]
	case p > 51:
		return c[3]
	case p > 34:
		return c[2]
	case p > 17:
		return c[1]
	default:
		return c[0]
	}
}

func colorForFreq(c *[6]string, p float64) string {
	switch {
	case p > 5:
		return c[5]
	case p > 4:
		return c[4]
	case p > 3:
		return c[1]
	default:
		return c[0]
	}
}
