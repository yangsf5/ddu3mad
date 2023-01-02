package functionasparam

import (
	"fmt"
	"sync"
	"time"
)

type A2 struct{}

func (m *A2) Do(statType string, urls []string) []map[string]any {
	return do(statType, urls, m.pull)
}

// statType 虽然不会用，但也传过来
func (m *A2) pull(statType, url string) string {
	return fmt.Sprintf("A.pull url=%s", url)
}

type B2 struct{}

func (m *B2) Do(statType string, urls []string) []map[string]any {
	return do(statType, urls, m.pull)
}

func (m *B2) pull(statType, url string) string {
	return fmt.Sprintf("B.pull statType=%s, url=%s", statType, url)
}

func do(statType string, urls []string, pullF func(statType, url string) string) []map[string]any {
	chanRet := make(chan map[string]any, len(urls))
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			startTime := time.Now().Unix()
			ret := pullF(statType, url) // pullF 函数由参数传过来
			spend := time.Now().Unix() - startTime

			chanRet <- map[string]any{"type": statType, "url": url, "ret": ret, "spend": spend}
		}(url)
	}

	wg.Wait()
	close(chanRet)

	rets := make([]map[string]any, 0, len(urls))
	for ret := range chanRet {
		rets = append(rets, ret)
	}

	return rets
}
