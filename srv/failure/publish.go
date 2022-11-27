package failuresrv

import (
	"fmt"
	"os"
)

func Publish(format string, a ...interface{}) {
	format = fmt.Sprintf("%s\n", format)

	fmt.Fprintf(os.Stderr, format, a...)
}
