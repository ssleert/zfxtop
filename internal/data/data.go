package data

import (
	"time"
)

// mem frame info
type memoryInfo struct {
	Used      float64
	UsedPerc  int
	UsedGraph []int

	Available      float64
	AvailablePerc  int
	AvailableGraph []int

	Free      float64
	FreePerc  int
	FreeGraph []int

	// swap frame info
	SwapTotal      float64
	SwapTotalGraph []int

	SwapUsed      float64
	SwapUsedPerc  int
	SwapUsedGraph []int
}

// disk frame info
type diskInfo struct {
	RootUsed     float64
	RootUsedPerc int

	HomeUsed     float64
	HomeUsedPerc int

	UsrUsed     float64
	UsrUsedPerc int
}

// battery frame info
type batInfo struct {
	Perc int
	Life time.Duration
}

// battery
// data struct for dynamic data updates
type Dynamic struct {
	// time on top
	Time time.Time

	// cpu frame info

	// why not make another type for cpu info?
	// coz, frequency and temperature find very different values.
	// and don't depend on each other
	CpuLoad      int
	CpuLoadGraph []int

	CpuFreq float64
	CpuTemp int

	Mem  memoryInfo
	Disk diskInfo
	Bat  batInfo
}

// data channels for goroutines pool
type pool struct {
	// time on top
	time chan time.Time

	// cpu frame info
	cpuLoad      chan int
	cpuLoadGraph chan []int

	cpuFreq chan float64
	cpuTemp chan int

	mem  chan memoryInfo
	disk chan diskInfo
	bat  chan batInfo

	err chan error
	n   int
}

type Static struct {
	// static mem frame info
	CpuName string

	// static mem frame info
	MemTotal float64

	// static disk frame info
	DiskTotal float64

	// static info frame info
	DistroName string
	HostName   string
}

// start goroutines for data collections
func Start() *pool {
	pool := pool{
		time:    make(chan time.Time),
		cpuLoad: make(chan int),
		cpuFreq: make(chan float64),
		cpuTemp: make(chan int),
		mem:     make(chan memoryInfo),
		disk:    make(chan diskInfo),
		bat:     make(chan batInfo),
		err:     make(chan error),
		n:       7,
	}

	go getTimeNow(pool.time, pool.err)
	go getCpuLoad(pool.cpuLoad, pool.err)
	go getCpuFreq(pool.cpuFreq, pool.err)
	go getCpuTemp(pool.cpuTemp, pool.err)
	go getMem(pool.mem, pool.err)
	go getDiskInfo(pool.disk, pool.err)
	go getBat(pool.bat, pool.err)

	return &pool
}

func handleErr(errch chan error, n int) (err error) {
	for i := 0; i < n; i++ {
		err = <-errch
		if err != nil {
			return
		}
	}
	return
}

func (p *pool) Update(d *Dynamic) error {
	err := handleErr(p.err, p.n)
	if err != nil {
		return err
	}

	d.Time = <-p.time
	d.CpuLoad = <-p.cpuLoad
	d.CpuFreq = <-p.cpuFreq
	d.CpuTemp = <-p.cpuTemp
	d.Mem = <-p.mem
	d.Disk = <-p.disk
	d.Bat = <-p.bat

	return nil
}

func Update(d *Static) error {
	cpuNamech := make(chan string)
	memTotalch := make(chan float64)
	diskTotalch := make(chan float64)
	distroNamech := make(chan string)
	hostNamech := make(chan string)
	errch := make(chan error)

	go getCpuModel(cpuNamech, errch)
	go getMemTotal(memTotalch, errch)
	go getDiskSize(diskTotalch, errch)
	go getDistroName(distroNamech, errch)
	go getHostName(hostNamech, errch)

	err := handleErr(errch, 5)
	if err != nil {
		return err
	}

	d.CpuName = <-cpuNamech
	d.MemTotal = <-memTotalch
	d.DiskTotal = <-diskTotalch
	d.DistroName = <-distroNamech
	d.HostName = <-hostNamech

	return nil
}
