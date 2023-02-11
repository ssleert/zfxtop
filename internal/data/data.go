package data

import (
	"time"
)

// data struct for dynamic data updates
type Dynamic struct {
	// time on top
	Time time.Time

	// cpu frame info
	CpuLoad      int
	CpuLoadGraph []int

	CpuFreq float64
	CpuTemp int

	// mem fram info
	MemUsed      float64
	MemUsedProc  int
	MemUsedGraph []int

	MemAvailable      float64
	MemAvailableProc  int
	MemAvailableGraph []int

	MemFree      float64
	MemfreeProc  int
	MemFreeGraph []int

	// swap frame info
	SwapTotal      float64
	SwapTotalProc  int
	SwapTotalGraph []int

	SwapFree      float64
	SwapFreeProc  int
	SwapFreeGraph []int

	// disk frame info
	DiskRootUsed     float64
	DiskRootUsedProc int

	DiskHomeUsed     float64
	DiskHomeUsedProc int

	DiskUsrUsed     float64
	DiskUsrUsedProc int

	// battery frame info
	BatCharge int
	BatLife   time.Duration
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

// data channels for goroutines pool
type pool struct {
	start chan struct{}

	// time on top
	time chan time.Time

	// cpu frame info
	cpuLoad      chan int
	cpuLoadGraph chan []int

	cpuFreq chan float64
	cpuTemp chan int

	// mem frame info
	memUsed      chan float64
	memUsedProc  chan int
	MemUsedGraph chan []int

	memAvailable      chan float64
	memAvailableProc  chan int
	memAvailableGraph chan []int

	memFree      chan float64
	memfreeProc  chan int
	memFreeGraph chan []int

	// swap frame info
	swapTotal      chan float64
	swapTotalProc  chan int
	swapTotalGraph chan []int

	swapFree      chan float64
	swapFreeProc  chan int
	swapFreeGraph chan []int

	// disk frame info
	diskRootUsed     chan float64
	diskRootUsedProc chan int

	diskHomeUsed     chan float64
	diskHomeUsedProc chan int

	diskUsrUsed     chan float64
	diskUsrUsedProc chan int

	// battery frame info
	batCharge chan int
	batLife   chan time.Duration

	err chan error
	n   int
}

// start goroutines for data collections
func Start() *pool {
	pool := pool{
		start: make(chan struct{}),
		time:  make(chan time.Time),
		err:   make(chan error),
		n:     1,
	}

	go GetTimeNow(pool.start, pool.err, pool.time)

	return &pool
}

func startPool(start chan struct{}, n int) {
	for i := 0; i < n; i++ {
		start <- struct{}{}
	}
}

func handleErr(errch chan error, n int) error {
	for i := 0; i < n; i++ {
		err := <-errch
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *pool) Update(d *Dynamic) error {
	startPool(p.start, p.n)

	d.Time = <-p.time

	err := handleErr(p.err, p.n)
	if err != nil {
		return err
	}

	return nil
}

func Update(d *Static) error {
	cpuNamech := make(chan string)
	errch := make(chan error)

	go getCpuModel(cpuNamech, errch)

	d.CpuName = <-cpuNamech
	d.MemTotal = 15.55
	d.DiskTotal = 223
	d.DistroName = "Arch Linux"
	d.HostName = "sfome"

	err := <-errch
	if err != nil {
		return err
	}

	return nil
}
