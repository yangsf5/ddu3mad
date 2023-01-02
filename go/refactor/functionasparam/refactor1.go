package functionasparam

import (
	"fmt"
	"sync"
	"time"
)

type A1 struct{}

func (m *A1) Do(statType string, urls []string) []map[string]any {
	chanRet := make(chan map[string]any, len(urls))
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			startTime := time.Now().Unix()
			ret := m.pull(statType, url) // statType 虽然不会用，但也传过去
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

// statType 虽然不会用，但也传过来
func (m *A1) pull(statType, url string) string {
	return fmt.Sprintf("A.pull url=%s", url)
}

type B1 struct{}

func (m *B1) Do(statType string, urls []string) []map[string]any {
	chanRet := make(chan map[string]any, len(urls))
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			startTime := time.Now().Unix()
			ret := m.pull(statType, url)
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

func (m *B1) pull(statType, url string) string {
	return fmt.Sprintf("B.pull statType=%s, url=%s", statType, url)
}
