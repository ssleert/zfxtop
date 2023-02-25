package conf

import (
	"errors"
	"fmt"
)

type ColorOutOfRangeError struct {
	Name string
}

type BorderIsIncorrect struct {
	Name string
}

func (e *ColorOutOfRangeError) Error() string {
	return fmt.Sprintf("%s var color is out or range. plz check your conf file for correctness", e.Name)
}

func (e *BorderIsIncorrect) Error() string {
	return fmt.Sprintf("'%s' border type is incorrect", e.Name)
}

var ErrConfFileNotExist = errors.New("conf file not exists")
var ErrNoConfFiles = errors.New("no conf files finded in system")
