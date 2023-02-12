package data

// os indepent functions
// mostly based on go stdlib

import (
	"github.com/ssleert/ginip"
	"time"
)

func getValue(s string, v string) (string, error) {
	ini, err := ginip.Load(s)
	if err != nil {
		return "", err
	}
	r, err := ini.GetValueString("", v)
	if err != nil {
		return "", err
	}

	return r, nil
}

func GetTimeNow(ch chan time.Time, err chan error) {
	for {
		err <- nil
		ch <- time.Now()
	}
}
