package main

import "fmt"

func main() {
	typeAssert()
}

func typeAssert() {
	var v1 any = 3

	s, ok := v1.(string)
	fmt.Printf("ret1, s=[%s] ok=[%v]\n", s, ok)

	var v2 any = "3"

	i, ok := v2.(int)
	fmt.Printf("ret1, s=[%v] ok=[%v]\n", i, ok)
}
