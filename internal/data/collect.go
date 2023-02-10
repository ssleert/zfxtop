package data

import (
	"time"
)

func GetTimeNow(start chan struct{}, err chan error, ch chan time.Time) {
	for range start {
		ch <- time.Now()
	}
}
