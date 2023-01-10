package v03

import (
	"fmt"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	input := []int{1, 3, 99, 100}

	rets := Slice(input, func(k int, v int) int {
		return v * 10
	})

	var sum int
	for _, ret := range rets {
		sum += ret
	}

	fmt.Printf("Slice ints rets:%v\n", sum)

	input2 := []string{"a", "b", "z", "x"}

	rets2 := Slice(input2, func(k int, v string) string {
		return fmt.Sprintf("k:%d v:%s", k, v)
	})

	fmt.Printf("Slice strings rets:\n%v\n", strings.Join(rets2, "\n"))
}

func TestMap(t *testing.T) {
	input := map[string]string{"a": "1", "b": "2", "c": "3"}

	rets := Map(input, func(k string, v string) string {
		return fmt.Sprintf("k:%s v:%s", k, v)
	})

	fmt.Printf("Map rets:\n%s\n", strings.Join(rets, "\n"))
}
