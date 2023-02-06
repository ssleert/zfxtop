package data

import (
	"time"
)

type Dynamic struct {
	// time on top
	time time.Time

	// cpu frame info
	CpuLoad int
	CpuFreq float64
	CpuTemp int

	// mem fram info
	MemUsed          float64
	MemUsedProc      int
	MemAvailable     float64
	MemAvailableProc int
	MemFree          float64
	MemfreeProc      int

	// swap frame info
	SwapTotal     float64
	SwapTotalProc int
	SwapFree      float64
	SwapFreeProc  int

	// disk frame info
	DiskRootUsed     float64
	DiskRootUsedProc int
	DiskHomeUsed     float64
	DiskHomeUsedProc int
	DiskUsrUsed      float64
	DiskUsrUsedProc  int

	// battery frame info
	BatCharge int
	BatLife   time.Time
}

type Static struct {
	// static mem frame info
	CpuName string

	// static mem frame info
	MemTotal float64

	// static disk frame info
	DiskTotal int

	// static info frame info
	DistroName string
	HostName   string
}

func (d *Static) Update() {
	d.CpuName = "E5-2660"
	d.MemTotal = 15.55
	d.DiskTotal = 223
	d.DistroName = "Arch Linux"
	d.HostName = "sfome"
}
