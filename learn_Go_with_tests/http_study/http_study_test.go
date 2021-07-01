package http_study

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/* 重构 一： 提取判断 response.Status ，去掉 got 与 want */
func TestPlayerServer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("if returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinReuqest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.wincalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.wincalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

/* 辅助函数 */

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

/*
http.Handler

  - 通过实现这个接口来创建 web 服务器

  - 用 http.HandlerFunc 把普通函数转化为 http.Handler

  - 把 httptest.NewRecorder 作为一个 ResponseWriter 传进去，这样让你可以监视 handler 发送了什么响应

  - 使用 http.NewRequest 构建对服务器的请求

接口，模拟和依赖注入

  - 允许你以小步快速迭代的方式逐步构建系统

  - 允许你开发需要存储 handler 而无需实际存储

  - TDD 驱使你实现需要的接口
*/
