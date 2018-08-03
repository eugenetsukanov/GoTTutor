package main

import "fmt"

type persone struct {
	name string
	age int
}

func main() {
	p := persone{name: "Jake", age: 21}

	fmt.Println(p)
}