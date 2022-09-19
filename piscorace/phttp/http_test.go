package phttp

import (
	"net/http"
	"sync"
	"testing"
	"time"
)

var urlList = []string{
	"https://www.hao123.com",
	"https://www.baidu.com",
	"https://cloud.tencent.com",
	"https://kaifa.baidu.com",
}

func TestHttp(t *testing.T) {
	var wg = sync.WaitGroup{}
	h := http.Client{
		Timeout: time.Second * 20,
	}
	num := 1024
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			url := urlList[i%len(urlList)]
			resp, err := h.Get(url)
			if err != nil {
				t.Error(err)
			}
			t.Log(resp.StatusCode == http.StatusOK)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log("完成")
}
