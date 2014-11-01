package main

import (
	"github.com/mgutz/goa"
	f "github.com/mgutz/goa/filter"
	. "github.com/mgutz/godo"
	"github.com/mgutz/godo/util"
	"github.com/mgutz/str"
)

// Project is local project.
func Tasks(p *Project) {
	p.Task("default", D{"readme"})

	p.Task("install", func() {
		Run("go get github.com/robertkrimen/godocdown/godocdown")
	})

	p.Task("lint", func() {
		Run("golint .")
		Run("gofmt -w -s .")
		Run("go vet .")
		Run("go test")
	})

	p.Task("readme", W{"**/*.go"}, func() {
		Run("godocdown -output README.md")

		packageName, _ := util.PackageName("doc.go")

		// add godoc link
		goa.Pipe(
			f.Load("./README.md"),
			f.Str(str.ReplaceF("--", "\n[godoc](https://godoc.org/"+packageName+")\n", 1)),
			f.Write(),
		)
	})
}

func main() {
	Godo(Tasks)
}
