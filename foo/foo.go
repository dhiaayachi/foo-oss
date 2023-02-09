package foo

import "fmt"

type Foo interface {
	Name() string
	License() string
}

type OssFoo struct {
}

func (o OssFoo) Name() string {
	return "Foo 5"
}

func (o OssFoo) License() string {
	return "OSS"
}

func PrintFoo(f Foo) {
	fmt.Printf("%s:%s\n", f.Name(), f.License())
}
