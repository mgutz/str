// Package str is a comprehensive set of string functions to build more
// awesomeness. Str is a port of many of the functions from the JavaScript
// [string.js](http://stringjs.com), which includes my contributions to
// the project.
//
// Package str differs from string.js in that str is based on simple functions
// not an intermediate String object to be consistent with Go.
//
//      str.Between("<a>foo</a>", "<a>", "</a>") == "foo"
//
// Package str is designed to be pipelined.
//
//      s := str.Pipe(
//          "\nabcdef\n",
//          Clean,
//          BetweenF("a", "f"),
//          ChompLeftF("bc")
//      ) // "de"
//
// User defined filters can be added to the pipeline if they have a signature
// of
//
//      func(string) string
//
package str
