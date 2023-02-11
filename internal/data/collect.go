package data

// os depent functions

import (
	"github.com/ssleert/memory"
	"github.com/ssleert/sfolib"
	"os"
	"strings"
	"syscall"
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
			case strings.Contains(e, "Turion"):
				result.WriteString("Turion ")
				result.WriteString(cpuName[2])
				result.WriteRune(' ')
				result.WriteString(cpuName[3])
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
