package test

import (
	"fmt"
	"testing"
	"time"
)

func TestConv(t *testing.T) {
	var v1 any = 3

	s, ok := v1.(string)
	fmt.Printf("ret1, s=[%s] ok=[%v]\n", s, ok)

	var v2 any = "3"

	i, ok := v2.(int)
	fmt.Printf("ret2, s=[%v] ok=[%v]\n", i, ok)

	var v3 any = nil

	iv3, ok := v3.(int)
	fmt.Printf("ret3, s=[%v] ok=[%v]\n", iv3, ok)

	iv4, ok := v1.(map[string]any)
	// iv4["1"] = 2 // nil map will panic
	fmt.Printf("ret4, s=[%v] ok=[%v] t=[%T] iv4==nil?[%v]\n", iv4, ok, iv4, iv4 == nil)

	iv5, ok := v1.(time.Time)
	fmt.Printf("ret5, s=[%v] ok=[%v] t=[%T]\n", iv5, ok, iv5)

	println(v3.(int)) // 直接强转而不检查时，nil 就会崩溃

}
