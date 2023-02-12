package main

import (
	"fmt"
	"github.com/ssleert/sfolib"
	"strconv"
	"strings"
	"time"
)

func getCPUSample() (int, int, error) {
	lines, err := sfolib.LoadFile("/proc/stat")
	if err != nil {
		return 0, 0, err
	}

	var (
		total int
		idle  int
	)
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.Atoi(fields[i])
				if err != nil {
					return 0, 0, err
				}
				total += val
				if i == 4 {
					idle = val
				}
			}
			return idle, total, nil
		}
	}
	panic("unreachable")
}

func main() {
	idle0, total0, err := getCPUSample()
	if err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	idle1, total1, err := getCPUSample()
	if err != nil {
		panic(err)
	}

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}
