package conf

import (
	"errors"
	"fmt"
)

type ColorOutOfRangeError struct {
	Name string
}

func (e *ColorOutOfRangeError) Error() string {
	return fmt.Sprintf("%s var color is out or range. plz check your conf file for correctness", e.Name)
}

var ErrConfFileNotExist = errors.New("conf file not exists")
