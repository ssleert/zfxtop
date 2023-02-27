package data

// os depent functions

import (
	"errors"
	"github.com/ssleert/memory"
	"github.com/ssleert/sfolib"
	"math"
	"os"
	"path/filepath"
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
				switch cpuName[2] {
				case "D", "1", "2", "3", "4":
					result.WriteString(cpuName[2])
				case "CPU":
					result.WriteString(cpuName[3])
				}
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
	}

	errch <- err
	ch <- float64(sd) / GB
}

func getDiskSize(ch chan float64, errch chan error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		errch <- err
		ch <- 0
	}
	result := float64(uint64(stat.Bsize)*uint64(stat.Blocks)) / (GB * MB)

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
			for i, e := range fields[1:] {
				val, err := strconv.Atoi(e)
				if err != nil {
					return 0, 0, err
				}
				total += val
				if i == 3 {
					idle = val
				}
			}
			return idle, total, nil
		}
	}
	panic("unreachable")
}

func getCpuLoad(ch chan int, errch chan error) {
	idle0, total0, err := getCPUSample()
	if err != nil {
		errch <- err
		ch <- 0
	}

	for {
		time.Sleep(100 * time.Millisecond)

		idle1, total1, err := getCPUSample()
		if err != nil {
			errch <- err
			ch <- 0
		}

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)

		load := int(math.Round(
			sfolib.Perc(totalTicks-idleTicks, totalTicks),
		))
		if load < 0 || load > 100 {
			errch <- errors.New("cpu load is higher than 100%")
			ch <- 0
		}
		errch <- nil
		ch <- load

		idle0, total0, err = getCPUSample()
		if err != nil {
			errch <- err
			ch <- 0
		}
	}
}

func getCpuFreq(ch chan float64, errch chan error) {
	for {
		lines, err := sfolib.LoadFile("/proc/cpuinfo")
		if err != nil {
			errch <- err
			ch <- 0
		}

		var (
			cores int
			freq  float64
		)
		for _, line := range lines {
			if strings.HasPrefix(line, "cpu MHz") {
				s := strings.Fields(line)
				sd, err := strconv.ParseFloat(s[3], 64)
				if err != nil {
					errch <- err
					ch <- 0
				}

				freq += sd
				cores++
			}
		}

		result := freq / float64(cores) / 1000

		errch <- nil
		ch <- result
	}
}

func getCpuTemp(ch chan int, errch chan error) {
	var (
		isCoreTemp  bool
		coreTempDir string
	)

	hwms, err := os.ReadDir("/sys/class/hwmon")
	if err != nil {
		errch <- err
		ch <- 0
	}
	for _, e := range hwms {
		hwmn := "/sys/class/hwmon/" + e.Name() + "/"
		f, err := os.ReadFile(hwmn + "name")
		if err != nil {
			errch <- err
			ch <- 0
		}
		if string(f) == "coretemp\n" || string(f) == "k10temp\n" {
			isCoreTemp = true
			coreTempDir = hwmn
		}
	}

	if isCoreTemp {
		d, err := os.ReadDir(coreTempDir)
		if err != nil {
			errch <- err
			ch <- 0
		}

		for {
			var (
				temp  int
				count float64
			)
			for _, e := range d {
				fn := e.Name()
				if strings.HasPrefix(fn, "temp") && strings.HasSuffix(fn, "_input") {
					f, err := os.ReadFile(coreTempDir + fn)
					if err != nil {
						errch <- err
						ch <- 0
					}
					t, err := strconv.Atoi(string(f[:len(f)-1]))
					if err != nil {
						errch <- err
						ch <- 0
					}
					temp += t
					count++
				}
			}
			errch <- nil
			ch <- int(math.Round(float64(temp)/count)) / 1000
		}
	}

	for {
		errch <- nil
		ch <- 0
	}
}

func getMem(memch chan memoryInfo, errch chan error) {
	totalint, err := memory.GetTotalRam()
	if err != nil {
		errch <- err
		memch <- memoryInfo{}
	}
	total := float64(totalint)

	for {
		ram, err := memory.GetRam()
		if err != nil {
			errch <- err
			memch <- memoryInfo{}
		}

		var (
			SwapTotal    float64
			SwapUsed     float64
			SwapUsedPerc int
		)
		if ram.SwapTotal != 0 {
			SwapTotal = float64(ram.SwapTotal) / GB
			SwapUsed = (float64(ram.SwapTotal) - float64(ram.SwapFree)) / GB
			SwapUsedPerc = int(math.Round(
				sfolib.Perc(float64(ram.SwapTotal)-float64(ram.SwapFree), float64(ram.SwapTotal)),
			))
		}

		errch <- nil
		memch <- memoryInfo{
			Used: (total - float64(ram.Available)) / GB,
			UsedPerc: int(math.Round(
				sfolib.Perc(total-float64(ram.Available), total),
			)),
			Available: float64(ram.Available) / GB,
			AvailablePerc: int(math.Round(
				sfolib.Perc(float64(ram.Available), total),
			)),
			Free: float64(ram.Free) / GB,
			FreePerc: int(math.Round(
				sfolib.Perc(float64(ram.Free), total),
			)),
			SwapTotal:    SwapTotal,
			SwapUsed:     SwapUsed,
			SwapUsedPerc: SwapUsedPerc,
		}
	}
}

type dirInfo struct {
	Used     float64
	UsedPerc int
}

func getDiskInfo(ch chan diskInfo, errch chan error) {
	rootDir := make(chan dirInfo)
	homeDir := make(chan dirInfo)
	usrDir := make(chan dirInfo)
	errRoot := make(chan error)
	errHome := make(chan error)
	errUsr := make(chan error)

	go func(ch chan dirInfo, errch chan error) {
		var stat syscall.Statfs_t

		for {
			err := syscall.Statfs("/", &stat)
			if err != nil {
				errch <- err
				ch <- dirInfo{}
			}

			errch <- nil
			ch <- dirInfo{
				Used: float64((stat.Blocks-stat.Bfree)*uint64(stat.Bsize)) / (GB * MB),
				UsedPerc: int(math.Round(
					sfolib.Perc(
						float64((stat.Blocks-stat.Bfree)*uint64(stat.Bsize)),
						float64(stat.Blocks*uint64(stat.Bsize)),
					),
				)),
			}
		}
	}(rootDir, errRoot)

	go func(ch chan dirInfo, errch chan error) {
		dir := "/home"

		var stat syscall.Statfs_t
		err := syscall.Statfs(dir, &stat)
		if err != nil {
			errch <- err
			ch <- dirInfo{}
		}

		var used int64
		filepath.Walk(dir,
			func(_ string, info os.FileInfo, _ error) error {
				if !info.IsDir() {
					used += info.Size()
				}
				return nil
			},
		)

		result := float64(used) / (GB * MB)
		all := float64(stat.Blocks * uint64(stat.Bsize) / (GB * MB))

		for {
			errch <- nil
			ch <- dirInfo{
				Used: result,
				UsedPerc: int(math.Round(
					sfolib.Perc(result, all),
				)),
			}
		}
	}(homeDir, errHome)

	go func(ch chan dirInfo, errch chan error) {
		dir := "/usr"

		var stat syscall.Statfs_t
		err := syscall.Statfs(dir, &stat)
		if err != nil {
			errch <- err
			ch <- dirInfo{}
		}

		var used int64
		filepath.Walk(dir,
			func(_ string, info os.FileInfo, _ error) error {
				if !info.IsDir() {
					used += info.Size()
				}
				return nil
			},
		)

		result := float64(used) / (GB * MB)
		all := float64(stat.Blocks * uint64(stat.Bsize) / (GB * MB))

		for {
			errch <- nil
			ch <- dirInfo{
				Used: result,
				UsedPerc: int(math.Round(
					sfolib.Perc(result, all),
				)),
			}
		}
	}(usrDir, errUsr)

	ticker := time.NewTicker(100 * time.Millisecond)
	var (
		rootDirTmp dirInfo
		homeDirTmp dirInfo
		usrDirTmp  dirInfo
		err        error
	)
	for {
		select {
		case <-ticker.C:
			errch <- nil
			ch <- diskInfo{
				RootUsed:     rootDirTmp.Used,
				RootUsedPerc: rootDirTmp.UsedPerc,
				HomeUsed:     homeDirTmp.Used,
				HomeUsedPerc: homeDirTmp.UsedPerc,
				UsrUsed:      usrDirTmp.Used,
				UsrUsedPerc:  usrDirTmp.UsedPerc,
			}
			continue
		case err = <-errRoot:
			if err != nil {
				errch <- err
				ch <- diskInfo{
					RootUsed:     0,
					RootUsedPerc: 0,
					HomeUsed:     homeDirTmp.Used,
					HomeUsedPerc: homeDirTmp.UsedPerc,
					UsrUsed:      usrDirTmp.Used,
					UsrUsedPerc:  usrDirTmp.UsedPerc,
				}
				continue
			}
			rootDirTmp = <-rootDir
		case err = <-errHome:
			if err != nil {
				errch <- err
				ch <- diskInfo{
					RootUsed:     rootDirTmp.Used,
					RootUsedPerc: rootDirTmp.UsedPerc,
					HomeUsed:     0,
					HomeUsedPerc: 0,
					UsrUsed:      usrDirTmp.Used,
					UsrUsedPerc:  usrDirTmp.UsedPerc,
				}
				continue
			}
			homeDirTmp = <-homeDir
		case err = <-errUsr:
			if err != nil {
				errch <- err
				ch <- diskInfo{
					RootUsed:     rootDirTmp.Used,
					RootUsedPerc: rootDirTmp.UsedPerc,
					HomeUsed:     homeDirTmp.Used,
					HomeUsedPerc: homeDirTmp.UsedPerc,
					UsrUsed:      0,
					UsrUsedPerc:  0,
				}
				continue
			}
			usrDirTmp = <-usrDir
		}
	}
}

func getBat(ch chan batInfo, errch chan error) {
	file := "/sys/class/power_supply/"
	dr, err := os.ReadDir(file)
	if err != nil {
		errch <- err
		ch <- batInfo{}
	}

	bats := make([]string, 0, 5)
	for _, e := range dr {
		if strings.HasPrefix(e.Name(), "BAT") {
			bats = append(bats, file+e.Name()+"/capacity")
		}
	}

	if len(bats) < 1 {
		for {
			errch <- nil
			ch <- batInfo{
				Perc: 101,
				Life: 100 * time.Minute,
			}
		}
	}

	for {
		var sum int
		for _, e := range bats {
			f, err := sfolib.ReadFirstLine(e)
			if err != nil {
				errch <- err
				ch <- batInfo{}
			}
			charge, err := strconv.Atoi(string(f))
			if err != nil {
				errch <- err
				ch <- batInfo{}
			}

			sum += charge
		}

		errch <- nil
		ch <- batInfo{
			Perc: sum / len(bats),
			Life: 100 * time.Minute,
		}
	}
}

// +=========================================================+
