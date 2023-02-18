package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirSizeMB(path string) (float64, error) {
	var size int64
	err := filepath.Walk("/", func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return float64(size) / 1024 / 1024 / 1024, err
}

func main() {
	fmt.Println(DirSizeMB("/usr"))
}
