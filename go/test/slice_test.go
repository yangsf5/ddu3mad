package test

import (
	"fmt"
	"testing"
)

func TestAppendNil(t *testing.T) {
	s := []string{"a", "b", "c", "d"}

	ms := make([]map[string]int, 0, len(s))

	for k, v := range s {
		m := map[string]int{v: k}
		ms = append(ms, m)
	}

	ms = append(ms, nil)

	fmt.Printf("%v\n", ms)

	ms2 := make([]map[string]int, 0, len(s))
	ms2 = append(ms2, nil, map[string]int{"x": 1})
	fmt.Printf("len:%d cap:%d slice:%v ms2[0]:%v ms2[1]:%v\n", len(ms2), cap(ms2), ms2, ms2[0], ms2[1])
}
