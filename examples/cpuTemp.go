package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func getTemp() int {
	hwms, _ := os.ReadDir("/sys/class/hwmon")
	for _, e := range hwms {
		hwmn := "/sys/class/hwmon/" + e.Name() + "/"
		f, err := os.ReadFile(hwmn + "name")
		if err != nil {
			panic(err)
		}
		if string(f) == "coretemp\n" {
			d, err := os.ReadDir(hwmn)
			if err != nil {
				panic(err)
			}

			var (
				temp  int
				count float64
			)
			for _, e := range d {
				fn := e.Name()
				if strings.HasPrefix(fn, "temp") && strings.HasSuffix(fn, "_input") {
					f, err := os.ReadFile(hwmn + fn)
					if err != nil {
						panic(err)
					}
					t, err := strconv.Atoi(string(f[:len(f)-1]))
					if err != nil {
						panic(err)
					}
					temp += t
					count++
				}
			}
			temp = int(math.Round(float64(temp)/count)) / 1000
			return temp
		}
	}
	return 0
}

func main() {
	print(getTemp())
	print("\n")
}
