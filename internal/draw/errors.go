package draw

import (
	"fmt"
)

type TermSizeTooLittle struct {
	X int
	Y int
}

func (err *TermSizeTooLittle) Error() string {
	return fmt.Sprintf(
		"%d,%d term size too small",
		err.X,
		err.Y,
	)
}
