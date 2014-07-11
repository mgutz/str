// Package str is a comprehensive set of string functions to build more
// awesomeness in Go. Str is a port of the JavaScript [string.js](http://stringjs.com),
// including my contributions to the project.
//
// Str does not duplicate functionality found in `strings` or `strconv`. Str
// may add filter versions of functions found in those packages for use with
// Pipe.
//
// Str is based on simple functions not an intermediate String object to be
// more consistent with the standard libraries Go.
//
//      // "foo"
//      str.Between("<a>foo</a>", "<a>", "</a>")
//
// Str is designed to be pipelined.
//
//      // "de"
//      s := str.Pipe("\nabcdef\n", Clean, BetweenF("a", "f"), ChompLeftF("bc"))
//
// User-defined filters can be added to the pipeline by creating a function
// or closure that returns a function with this signature
//
//      func(string) string
//
package str
