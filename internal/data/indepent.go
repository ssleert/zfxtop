package data

// os indepent functions
// mostly based on go stdlib

import (
	"time"
)

func GetTimeNow(start chan struct{}, err chan error, ch chan time.Time) {
	for range start {
		ch <- time.Now()
		err <- nil
	}
}
