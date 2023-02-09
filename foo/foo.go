package foo

import "fmt"

type Foo interface {
	Name() string
	License() string
}

type OssFoo struct {
}

func (o OssFoo) Name() string {
	return "Foo 8"
}

func (o OssFoo) License() string {
	return "OSS 24"
}

func PrintFoo(f Foo) {
	fmt.Printf("%s:%s\n", f.Name(), f.License())
}
