package data

func clearGraph(g []int) {
	for i := range g {
		g[i] = 0
	}
}

func addToGraph(g []int, n int) {
	for i := 0; i < len(g)-1; i++ {
		g[i] = g[i+1]
	}
	g[len(g)-1] = n
}
