package main

import (
	"fmt"
	"github.com/ssleert/zfxtop/internal/self"
)

func main() {
	fmt.Println(self.Version)
	fmt.Println(self.Commit)
}
