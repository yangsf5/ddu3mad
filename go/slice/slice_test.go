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
