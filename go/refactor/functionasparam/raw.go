package functionasparam

import (
	"fmt"
	"sync"
	"time"
)

type A struct{}

func (m *A) Do(statType string, urls []string) []map[string]any {
	chanRet := make(chan map[string]any, len(urls))
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			startTime := time.Now().Unix()
			ret := m.pull(url) // 唯一区别，这里不传 statType
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

func (m *A) pull(url string) string {
	return fmt.Sprintf("A.pull url=%s", url)
}

type B struct{}

func (m *B) Do(statType string, urls []string) []map[string]any {
	chanRet := make(chan map[string]any, len(urls))
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			startTime := time.Now().Unix()
			ret := m.pull(statType, url) // 唯一区别，这里要传 statType
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

func (m *B) pull(statType, url string) string {
	return fmt.Sprintf("B.pull statType=%s, url=%s", statType, url)
}
