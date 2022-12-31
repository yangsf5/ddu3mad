package test

import (
	"fmt"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	rets := make([]int, 0, 3)
	fmt.Printf("after make, len:%d cap:%d slice:%v\n", len(rets), cap(rets), rets)

	for i := 0; i < 9; i++ {
		rets = append(rets, i)
		fmt.Printf("after append %d, len:%d cap:%d slice:%v\n", i+1, len(rets), cap(rets), rets)
	}
}
