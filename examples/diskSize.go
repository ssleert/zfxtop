package main

import (
	"fmt"
	"syscall"
)

func main() {
	var stat syscall.Statfs_t
	syscall.Statfs("/", &stat)
	fmt.Println("Disk size:", stat.Bsize*int64(stat.Blocks)/1024/1024/1024, "GB")
}
