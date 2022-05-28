package main

import (
	"fmt"
	"github.com/unsafe-risk/go-safe"
)

var _ safe.Safety[*Person] = (*Person)(nil)

type Person struct {
	Id   int64
	Name string
	Age  int16
}

func (p *Person) IsNil() bool {
	return p == nil
}

func (p *Person) Default() *Person {
	return &Person{
		Id:   1,
		Name: "empty name",
		Age:  111,
	}
}

func main() {
	var p *Person
	fmt.Println(p)
	p = safe.Safe(p)
	fmt.Println(p)

	var i *int
	fmt.Println(safe.NumOrZero(i))
	fmt.Println(safe.NumOrDefault(i, func() int {
		return 123
	}))

	var f *float64
	fmt.Println(safe.NumOrZero(f))
	fmt.Println(safe.NumOrDefault(f, func() float64 {
		return 3.141592
	}))

	var str *string
	fmt.Println(safe.StrOrZero(str))
	fmt.Println(safe.StrOrDefault(str, func() string {
		return "hello world"
	}))
}
