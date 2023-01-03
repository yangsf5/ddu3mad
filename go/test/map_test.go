package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	keys := []string{"A", "B", "C"}

	m := make(map[string]int, len(keys))

	for k, v := range keys {
		m[v] = k
	}

	fmt.Printf("%v\n", m)

	m2 := map[string]int{keys[0]: 0}
	fmt.Printf("%v\n", m2)
}
