package slice

import "fmt"

func RangeWrite() []string {
	s1 := []string{"a", "b", "c", "d"}
	for k, v := range s1 {
		v = fmt.Sprintf("%d-%s", k, v)
	}

	return s1
}

func RangeWriteMap() []map[string]string {
	s1 := []map[string]string{{"i": "a"}, {"i": "b"}, {"i": "c"}, {"i": "d"}}
	for _, v := range s1 {
		v["j"] = fmt.Sprintf("j-%v", v["i"])
	}

	return s1
}
