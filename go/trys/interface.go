package trys

import (
	"fmt"
)

type vapi interface {
	sayHello()
}

type myapi struct {
	Hstring string
}

func (v *myapi) sayHello() {
	fmt.Println(v.Hstring)
}

type Proxy struct {
	api vapi
}

func Show() {
	a := myapi{"interface"}
	p := Proxy{&a}

	p.api.sayHello()
}
