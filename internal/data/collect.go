package data

// os depent functions

import (
	"errors"
	"github.com/ssleert/memory"
	"github.com/ssleert/sfolib"
	"math"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// +================= static info functions =================+

func getCpuModel(ch chan string, errch chan error) {
	var result strings.Builder
	result.Grow(100)

	s, err := sfolib.ReadLines("/proc/cpuinfo", 5)
	if err != nil {
		errch <- err
		ch <- ""
		return
	}

	cpuName := strings.Fields(s[4])[3:]
	cpuVendor := strings.Fields(s[1])[2]
	switch cpuVendor {
	case "AuthenticAMD":
		for _, e := range cpuName {
			switch {
			case strings.Contains(e, "Ryzen"):
				result.WriteRune('R')
				result.WriteString(cpuName[2])
				result.WriteRune('-')
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "FX"):
				result.WriteString("FX-")
				result.WriteString(strings.Split(e, "-")[1])
				break
			case strings.Contains(e, "Athlon"):
				result.WriteString("Athlon ")
				result.WriteString(cpuName[2])
				break
			case strings.Contains(e, "Turion"):
				result.WriteString("Turion ")
				result.WriteString(cpuName[2])
				result.WriteRune(' ')
				result.WriteString(cpuName[3])
				break
			}
		}
		if result.Len() == 0 {
			result.WriteString("AMD")
		}
	case "GenuineIntel":
		for i, e := range cpuName {
			switch {
			case strings.Contains(e, "i3"),
				strings.Contains(e, "i5"),
				strings.Contains(e, "i7"),
				strings.Contains(e, "i9"):
				result.WriteString(cpuName[i])
				break
			case strings.Contains(e, "Xeon"):
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "Celeron"):
				result.WriteString("Celeron ")
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "Duo"):
				result.WriteString("2d-")
				result.WriteString(cpuName[4])
				break
			case strings.Contains(e, "Pentium"):
				result.WriteString("Pentium ")
				result.WriteString(cpuName[2])
				break
			}
		}
		if result.Len() == 0 {
			result.WriteString("INTEL")
		}
	}
	if result.Len() == 0 {
		result.WriteString("UNKNOWN")
	}

	errch <- err
	ch <- result.String()
}

func getMemTotal(ch chan float64, errch chan error) {
	sd, err := memory.GetTotalRam()
	if err != nil {
		errch <- err
		ch <- 0
		return
	}

	errch <- err
	ch <- float64(sd) / 1024
}

func getDiskSize(ch chan float64, errch chan error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		errch <- err
		ch <- 0
		return
	}
	result := float64(stat.Bsize*int64(stat.Blocks)/1024/1024) / 1024

	errch <- nil
	ch <- result
}

func getDistroName(ch chan string, errch chan error) {
	rls := [...][2]string{
		{"/etc/os-release", "NAME"},
		{"/etc/lsb-release", "DISTRIB_DESCRIPTION"},
	}

	var result string
	for _, e := range rls {
		if sfolib.Exists(e[0]) {
			s, err := getValue(e[0], e[1])
			if err != nil {
				errch <- err
				ch <- ""
				return
			}
			result = s[1 : len(s)-1]
		}
	}
	if result == "" {
		result = "UNKNOWN"
	}

	errch <- nil
	ch <- result
}

func getHostName(ch chan string, errch chan error) {
	result := os.Getenv("hostname")
	if result == "" {
		var err error
		result, err = sfolib.ReadFirstLine("/etc/hostname")
		if err != nil {
			errch <- err
			ch <- ""
			return
		}
	}

	errch <- nil
	ch <- result
}

// +=========================================================+

// +================ dynamic info functions =================+

// get cpu ticks sample
// used in getCpuLoad
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

func getCpuLoad(ch chan int, errch chan error) {
	var (
		idle0, idle1   int
		total0, total1 int
		err            error
	)

	for {
		idle0, total0, err = getCPUSample()
		if err != nil {
			errch <- err
			ch <- 0
			return
		}
		time.Sleep(100 * time.Millisecond)
		idle1, total1, err = getCPUSample()
		if err != nil {
			errch <- err
			ch <- 0
			return
		}

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)

		errch <- nil
		load := int(math.Round((100 * (totalTicks - idleTicks) / totalTicks) * 0.9))
		if load < 0 || load > 100 {
			errch <- errors.New("cpu load is higher than 100%")
			ch <- 0
			return
		}
		ch <- load

		idle0, total0, err = getCPUSample()
		if err != nil {
			errch <- err
			ch <- 0
			return
		}
	}
}

// +=========================================================+
