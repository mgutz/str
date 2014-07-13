package tasks

import (
	"github.com/mgutz/goa"
	f "github.com/mgutz/goa/filter"
	"github.com/mgutz/gosu"
	"github.com/mgutz/gosu/util"
	"github.com/mgutz/str"
)

// Project is local project.
func Project(p *gosu.Project) {
	p.Task("default", []string{"readme"})

	p.Task("install", func() {
		util.Exec("go get github.com/robertkrimen/godocdown/godocdown")
	})

	p.Task("lint", func() {
		util.Exec("golint .")
		util.Exec("gofmt -w -s .")
		util.Exec("go vet .")
		util.Exec("go test")
	})

	p.Task("readme", gosu.Files{"**/*.go"}, func() {
		util.Exec("godocdown -output README.md")

		packageName, _ := util.PackageName("doc.go")

		// add godoc link
		goa.Pipe(
			f.Load("./README.md"),
			f.Str(str.ReplaceF("--", "\n[godoc](https://godoc.org/"+packageName+")\n", 1)),
			f.Write(),
		)
	})
}
