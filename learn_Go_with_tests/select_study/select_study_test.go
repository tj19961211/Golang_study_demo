package select_study

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/* 重构第一次: 改为使用模拟测试，这样我们就可以控制可靠的服务器来测试了

   重构二： 提取创建模拟 http server 函数
*/

/*
进程同步

  - Go 在并发方面很在行，为什么我们要一个接一个地测试哪个网站更快呢？我们应该能够同时测试两个。

  - 我们并不关心请求的 准确响应时间，我们只是需要知道哪个更快返回而已。

想实现这个，我们要介绍一个叫 select 的新构造（construct），它可以帮我们轻易清晰地实现进程同步。
*/
func TestRacer(t *testing.T) {
	t.Run("TestRacer", func(t *testing.T) {
		/*httptest.NewServer 接受一个我们传入的 匿名函数 http.HandlerFunc*/
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer1(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

/*
总结

  select

   - 可帮助你同时在多个 channel 上等待。

   - 有时你想在你的某个「案例」中使用 time.After 来防止你的系统被永久阻塞。

   httptest

   - 一种方便地创建测试服务器的方法，这样你就可以进行可靠且可控的测试。

   - 使用和 net/http 相同的接口作为「真实的」服务器会和真实环境保持一致，并且只需更少的学习。
*/
