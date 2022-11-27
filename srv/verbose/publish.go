package verbosesrv

import (
	"github.com/reiver/finger/arg"

	"fmt"
	"os"
)

func Publish(format string, a ...interface{}) {
	var verbose bool =
		arg.Verbose                     ||
		arg.VeryVerbose                 ||
		arg.VeryVeryVerbose             ||
		arg.VeryVeryVeryVerbose         ||
		arg.VeryVeryVeryVeryVerbose     ||
		arg.VeryVeryVeryVeryVeryVerbose

	if !verbose {
		return
	}

	format = fmt.Sprintf("[%s]\n", format)

	fmt.Fprintf(os.Stdout, format, a...)
}
