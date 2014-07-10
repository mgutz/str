# str

Go string library to build more awesomeness.

See [godoc](https://godoc.org/github.com/mgutz/str)

## About

Package str is a comprehensive set of string functions to build more
awesomeness. Str is a port of many of the functions from [stringjs](http://stringjs.com),
which includes my contributions to the project.

Package str differs from string.js in that str is based on simple functions
instead of an intermediate String object to be consistent with Go.

```go
     str.Between("<a>foo</a>", "<a>", "</a>") == "foo"
```

Package str is designed to be pipelined

```go
     s := str.Pipe(
         "\nabcdef\n",
         Clean,
         BetweenF("a", "f"),
         ChompLeftF("bc")
     ) // de
```

User defined filters may be added to the pipeline if they meet this
signature

```go
     func(string) string
```

## License

The MIT License
