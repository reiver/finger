package arg

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	Values []string
)

var (
	Verbose bool
	VeryVerbose bool
	VeryVeryVerbose bool
	VeryVeryVeryVerbose bool
	VeryVeryVeryVeryVerbose bool
	VeryVeryVeryVeryVeryVerbose bool
)

var (
	Help bool
)


func init() {
	flag.Usage = func() {
		var w io.Writer = flag.CommandLine.Output()

		fmt.Fprintf(w, "%s — a finger-protocol client"                            +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "Usage of %s:"                                             +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s [/switch] [user][@host[:port]…]"                  +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Typical Usage:"                                            +"\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s user@host"                                        +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s joeblow@example.com"                              +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s charles@reiver.ca"                                +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s malekeh@esphahan.ir"                              +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Typical Usage:"                                            +"\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s user@host:port"                                   +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s dariush@changelog.ca:1971"                        +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s elizabeth@glasgow.scot:7979"                      +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Typical Advanced Usage:"                                   +"\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s /switch user@host"                                +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s /W joeblow@example.com"                           +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s /W charles@reiver.ca"                             +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s /W malekeh@esphahan.ir"                           +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Typical Advanced Usage:"                                   +"\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s /switch user@host"                                +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s /W dariush@changelog.ca:1971"                     +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s /W elizabeth@glasgow.scot:7979"                   +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Other Usages:"                                             +"\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  "+"%s"                                                  +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s joeblow"                                          +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s joeblow@example.com@changelog.ca"                 +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s joeblow@example.com:7979@changelog.ca"            +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s joeblow@example.com:7979@changelog.ca:1079"       +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s janedoe@once.com@twice.net@thrice.org@fource.dev" +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @example.com"                                     +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @example.com:7979"                                +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @example.com@changelog.ca"                        +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @example.com:7979@changelog.ca"                   +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @example.com:7979@changelog.ca:1079"              +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @once.com@twice.net@thrice.org@fource.dev"        +"\n", os.Args[0])
		fmt.Fprintf(w, "  "+"%s @changelog.ca:1971"                               +"\n", os.Args[0])
		fmt.Fprintln(w)
		fmt.Fprint(w, "Flags:"                                                    +"\n")
		fmt.Fprintln(w)
		flag.PrintDefaults()
	}

	flag.BoolVar(&Verbose,                     "v",      false,                          "verbose logs outputted")
	flag.BoolVar(&VeryVerbose,                 "vv",     false,                     "very verbose logs outputted")
	flag.BoolVar(&VeryVeryVerbose,             "vvv",    false,                "very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVerbose,         "vvvv",   false,           "very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVerbose,     "vvvvv",  false,      "very very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVeryVerbose, "vvvvvv", false, "very very very very very verbose logs outputted")
	
	flag.BoolVar(&Help, "help", false, "outputs help message")

	flag.Parse()
	
	Values = flag.Args()

	// --help
	if Help {
		flag.Usage()
		os.Exit(0)
	}
}
