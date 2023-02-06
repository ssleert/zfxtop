package msg

import (
	"fmt"
	"os"
)

func ExitMsg(err error) {
	fmt.Println(err)
	os.Exit(1)
}
