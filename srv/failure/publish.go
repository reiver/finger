package failuresrv

import (
	"fmt"
	"os"
)

// Publish is used to output error messages that are sent just before the program exits (due to an error).
//
// For example:
//
//	package main
//	
//	import failure "github.com/reiver/go-finger/srv/failure"
//	
//	// ...
//	
//	err := something()
//	if nil != err {
//		failure.Publish("we have a problem: %s", err)
//		os.Exit(1)
//		return
//	}
func Publish(format string, a ...interface{}) {
	format = fmt.Sprintf("%s\n", format)

	fmt.Fprintf(os.Stderr, format, a...)
}
