package main

import (
	"fmt"
)

func main() {
	typeAssert()
}

func typeAssert() {
	var v1 any = 3

	s, ok := v1.(string)
	fmt.Printf("ret1, s=[%s] ok=[%v]\n", s, ok)

	var v2 any = "3"

	i, ok := v2.(int)
	fmt.Printf("ret2, s=[%v] ok=[%v]\n", i, ok)

	var v3 any = nil

	iv3, ok := v3.(int)
	fmt.Printf("ret3, s=[%v] ok=[%v]\n", iv3, ok)

	println(v3.(int)) // 直接强转而不检查时，nil 就会崩溃
}
