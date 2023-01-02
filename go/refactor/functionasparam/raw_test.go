package functionasparam

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	statTypeA := "stat_A"
	statTypeB := "stat_B"
	urlsA := []string{"mad1.com/a", "mad2.com/a", "mad3.com/a"}
	urlsB := []string{"mad1.com/b", "mad2.com/b", "mad3.com/b"}

	testGenRet(new(A).Do(statTypeA, urlsA))
	testGenRet(new(A1).Do(statTypeA, urlsA))
	testGenRet(new(A2).Do(statTypeA, urlsA))

	testGenRet(new(B).Do(statTypeB, urlsB))
	testGenRet(new(B1).Do(statTypeB, urlsB))
	testGenRet(new(B2).Do(statTypeB, urlsB))
}

func testGenRet(rets []map[string]any) string {
	rss := make([]string, 0, len(rets))
	for _, ret := range rets {
		rs := fmt.Sprintf("%s %s %s", ret["type"], ret["url"], ret["ret"])
		rss = append(rss, rs)
	}

	sort.Strings(rss)

	ret := strings.Join(rss, "\t")
	println(ret)

	return ret
}
