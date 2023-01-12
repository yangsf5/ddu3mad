package slice

import (
	"fmt"
	"testing"
)

func TestRangeWrite(t *testing.T) {
	ret1 := RangeWrite()
	fmt.Printf("%v\n", ret1)

	ret2 := RangeWriteMap()
	fmt.Printf("%v\n", ret2)
}

func TestAppend(t *testing.T) {
	// s := append(nil, []string{}) // 不行
	s := append([]string(nil), "test")
	fmt.Printf("%v\n", s)
}
