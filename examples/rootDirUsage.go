package main

import (
	"fmt"
	"syscall"
)

func main() {
	var stat syscall.Statfs_t
	syscall.Statfs("/", &stat)

	// Available blocks * size per block = available space in bytes
	fmt.Println(float64(stat.Bavail*uint64(stat.Bsize)) / 1024 / 1024 / 1024)
}
