package main

import (
	"github.com/reiver/finger/arg"
	failure "github.com/reiver/finger/srv/failure"
	verbose "github.com/reiver/finger/srv/verbose"

	"github.com/reiver/go-finger"

	"io"
	"net"
	"os"
)

func main() {

	// This ('arguments') holds what was on the command-line when the flags (and program-name) are removed.
	//
	// For example, if this program was called with the following command-line:
	//
	//	finger -v /W dariush@reiver.link
	//
	// Then logically 'arguments' would be:
	//
	//	arguments == []string{"/W", "dariush@reiver.link"}
	//
	// And also, for another example, if this program called with the following command-line:
	//
	//	finger -v dariush@reiver.link
	//
	// Then logically 'arguments' would be:
	//
	//	arguments == []string{"dariush@reiver.link"}
	var arguments []string = arg.Values

	// Here we see if there is a finger-protocol switch (as a string) from the command-line
	// which this program was called with.
	//
	// For example:
	//
	//	finger /W
	//	# switch -> "/W"
	//
	// And, for example:
	//
	//	finger /W dariush@reiver.link
	//	# switch -> "/W"
	//
	// Or, for example:
	//
	//	finger /W dariush@reiver.link@example.com
	//	# switch -> "/W"
	//
	// Or, alternatively, for example:
	//
	//	finger /W joeblow@example.com:1971
	//	# switch -> "/W"
	//
	// Or also, for example:
	//
	//	finger /W janedoe
	//	# switch -> "/W"
	//
	// Or even, for example:
	//
	//	finger /W @example.com
	//	# switch -> "/W"
	//
	// Etc.
	//
	// Although the switch could be something other than "/W".
	// (For example — "/PULL", "/ECHO", "/path/to/a/file/ext", etc.)
	// It just has to become with a "/" character.
	//
	// Also, there might not be a switch. For example:
	//
	//	finger joeblow@example.com
	//	# no switch
	var swtch finger.Switch
	func(){
		if len(arguments) < 1 {
			return
		}

		argument0 := arguments[0]

		var err error

		swtch, err = finger.ParseSwitch(argument0)
		if nil != err {
			return
		}

		arguments = arguments[1:]
	}()
	verbose.Publish("SWITCH: %q", swtch)

	// Get the finger-protocol query (as a string) from the command-line
	// which this program was called with.
	//
	// For example:
	//
	//	finger dariush@reiver.link
	//	# query -> "dariush@reiver.link"
	//
	// Or, for example:
	//
	//	finger dariush@reiver.link@example.com
	//	# query -> "dariush@reiver.link@example.com"
	//
	// Or, alternatively, for example:
	//
	//	finger joeblow@example.com:1971
	//	# query -> "joeblow@example.com:1971"
	//
	// Or also, for example:
	//
	//	finger janedoe
	//	# query -> "janedoe"
	//
	// Or even, for example:
	//
	//	finger @example.com
	//	# query -> "@example.com"
	//
	// Etc.
	//
	// Note that in this type of finger-protocol request:
	//
	//	"/W joeblow@example.com"
	//
	// The "/W" is NOT part of the query.
	//
	// Only the "joeblow@example.com" is the query.
	var q string
	{
		if 1 <= len(arguments) {
			q = arguments[0]
		}
	}
	verbose.Publish("QUERY: %q", q)

	// Here we parse the finger-protocol query string,
	// and if valid, turn it into a finger.Query.
	var query finger.Query
	{
		var err error

		query, err = finger.ParseQuery(q)
		if nil != err {
			failure.Publish("bad query: %s", err)
////////////////////////// EXIT
			os.Exit(1)
			return
		}
	}
	verbose.Publish("COMPILED QUERY: %q", q)

	// Here we infer what TCP address we should make a TCP connection to.
	// (That we will later make a finger-protocol request from.)
	//
	// And what the finger-protocol query we sent will be.
	var address finger.Address
	var target finger.Target
	{
		var subquery finger.Query

		address, subquery = query.ClientParameters()

		target = subquery.Target()
	}
	verbose.Publish("ADDRESS: %q", address)
	verbose.Publish("RESOLVED-ADDRESS: %q", address.Resolve())
	verbose.Publish("TARGET: %#v", target)

	// Create the finger-protocol request.
	//
	// We aren't sending it to the server yet.
	//
	// We are just preparing the request (and will use it a little later).
	var request finger.Request = finger.AssembleRequest(swtch, target)
	verbose.Publish("REQUEST: %q", request)

	// Connect to the finger-protocol server.
	var conn net.Conn
	{
		var err error

		conn, err = net.Dial("tcp", address.Resolve())
		if nil != err {
			failure.Publish("problem connecting to %q: %s", address, err)
////////////////////////// EXIT
			os.Exit(1)
			return
		}
	}
	defer conn.Close()

	// Bind the finger-protocol client to the TCP connection to the finger-protocol server.
	var client finger.Client = finger.AssembleClient(conn)

	// Send the finger-protocol request,
	// and receive finger-protocol response.
	var responseReader finger.ResponseReader
	{
		var err error

		responseReader, err = client.Do(request)
		if nil != err {
			failure.Publish("problem with doing request: %s", err)
////////////////////////// EXIT
			os.Exit(1)
			return
		}
	}

	// Send the contents of the finger-protocol response,
	// from the finger-protocol server, to STDOUT.
	//
	// This way the user can see it.
	// Or it can be piped to another program.
	// Or saved to a file.
	// Or whatever.
	{
		io.Copy(os.Stdout, responseReader)
	}
}
