package main

import "fmt"

func main() {
	i := 9
	
	inc(&i)

	fmt.Println(i)
}

func inc(x *int) {
	*x++
}