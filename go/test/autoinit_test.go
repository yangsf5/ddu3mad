package test

import (
	"fmt"
	"testing"
)

func TestMapAutoInit(t *testing.T) {
	ms := map[string][]string{}
	ms["a"] = append(ms["a"], "a1") // 这种不崩，因为已经初始化的 map，取值时，会自动给自己的内层成员初始化
	// ms["a"][2] = "a2" // 这种会崩溃，因为虽然 map 给内存的初始化了，但并有给扩充 len，所以 append 可以，但直接数据下边不可以

	fmt.Printf("ms:%v\n", ms)

	mm := map[string]map[string]string{}
	mm["l1"]["l2"] = "l3"                    // 这种会崩溃
	mm["l1"] = map[string]string{"l2": "l3"} // 这种不崩，是因为 mm 最外层在创建时就初始化了，然后按照规则，没有的 "l1" 会自动创建 map[string]string
	fmt.Printf("mm:%v\n", mm)

	// msm := map[string][]map[string]string{}
	// msm["1m"][]
}

func TestSliceAutoInit(t *testing.T) {
}
