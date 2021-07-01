package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
为什么不测试 JSON 字符串？

  你可以争论一个更简单的初始步骤就是断言响应体有一个特定的 JSON 字符串。

  以我的经验，断言 JSON 字符串的测试有以下问题。

    - 脆弱。如果你改了数据模型，测试将会失败。

    - 难以调试。在比较两个 JSON 字符串时，很难理解真正的问题是什么。

	- 意图不佳。当输出应该是 JSON 时，真正重要的是数据究竟是什么，而不是它的编码方式。

	- 重复测试标准库。没有必要测试标准库如何输出 JSON，它已经过测试。不要测试别人的代码。

相反，我们应该将 JSON 解析为与我们测试相关的数据结构。
*/
func TestLrague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}
		assertContentType(t, response, jsonContentType)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertContentType(t, response, jsonContentType)
		assertStatus(t, response.Code, http.StatusOK)

		assertLeague(t, got, wantedLeague)
	})
}

func TestRecordingWinsRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
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

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.HeaderMap)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

/*
总结

  我们继续使用 TDD 安全地迭代了我们的程序，使其能够通过路由以可维护的方式支持新端点，现在它可以为我们的客户返回 JSON。在下一章中，我们将介绍持久化数据和对联盟排序。

  我们所涵盖的内容：

    - 路由。标准库为你提供了易于使用的类型来处理路由。它完全支持 http.Handler 接口，因为你可以将路由分配给 Handler，而路由本身也是 Handler。它没有你可能期望的某些特性，例如路径变量（例如 /users/{id}）。你可以自己轻易地解析这些信息，但如果它成了负担，你可能会考虑查看其它路由库。大多数流行的库都坚持标准库的实现 http.Handler 的理念。

	- 类型嵌入。我们对这项技术略有提及，但你可以从 Effective Go 了解更多信息。如果你应该从中得到一个收获，那就是它极其有用，但是 总是考虑你的公开 API，只有适合被公开的才公开。

	- JSON 反序列化和序列化。标准库使得序列化和反序列化数据变得非常简单。它也是开放的配置，你可以根据需要自定义这些数据转换的工作方式。
*/
