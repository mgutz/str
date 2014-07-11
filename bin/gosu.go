package main

import (
	"github.com/mgutz/gosu"
	"github.com/mgutz/gosu/util"
)

func Project(p *gosu.Project) {
	p.Task("default", []string{"readme"})

	p.Task("install", func() {
		util.Exec("go get github.com/robertkrimen/godocdown/godocdown")
	})

	p.Task("readme", func() {
		util.Exec("godocdown -output README.md")
	})
}

func main() {
	gosu.Run(Project)
}
