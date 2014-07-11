// Package str is a comprehensive set of string functions to build more
// Go awesomeness. Str is a port of the JavaScript [string.js](http://stringjs.com)
// including my contributions to the project.
//
// Str does not duplicate functionality found in `strings` or `strconv`. Str
// may add filter versions of functions found in those packages for use with
// Pipe.
//
// Str is based on plain functions instead of object-based methods to be more
// consistent with Go standard libraries.
//
//      str.Between("<a>foo</a>", "<a>", "</a>") == "foo"
//
// Str is designed to be pipelined.
//
//      s := str.Pipe("\nabcdef\n", Clean, BetweenF("a", "f"), ChompLeftF("bc"))
//      s == "de"
//
// User-defined filters can be added to the pipeline by creating a function
// or closure that returns a function with this signature
//
//      func(string) string
//
package str
