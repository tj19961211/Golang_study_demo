package http_study

import (
	"fmt"
	"net/http"
)

/* HTTP 服务器 */

/*
你被要求创建一个 Web 服务器，用户可以在其中跟踪玩家赢了多少场游戏。

 - GET /players/{name} 应该返回一个表示获胜总数的数字

 - POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随着每次 POST 递增

我们将遵循 TDD 方法，尽可能快地让程序先可用，然后进行小步迭代改进，直到我们找到解决方案。通过采取这种方法我们

  - 在任何给定时间保持问题都是小问题

  - 不要陷入陷阱（rabbit holes）

  - 如果我们卡住或迷失了方向，回退不会前功尽弃。
*/

/*
## TDD 理念

### 持续迭代，重构

    在本书中，我们强调了编写测试并观察失败（红色），编写 最少量 代码跑通测试（绿色）然后重构的 TDD 过程。

    就 TDD 的安全性而言，编写最少量代码的这一规则非常重要。你应该尽快摆脱测试失败的状态（红色）。

> Kent Beck 这样描述它：
>
> 快速跑通测试，为满足必要条件暂时犯错亦可。

你可以先写一些存在已知问题的代码，因为后面你会基于 TDD 安全地进行重构。

### 如果不这样做？

  测试未通过前你写得越多，也就引入越来越多测试不能覆盖的问题。

  这样做是为了小步快速迭代，测试驱使你不会掉入陷阱。

#### 先有鸡还是先有蛋

	我们如何逐步建立这个？我们不能在没有数据的前提下 GET 一个玩家，而且似乎很难知道 POST 在没有 GET 的情况下是否工作。

    这就是 模拟 测试的亮点。

      - GET 需要一个类似 PlayerStore 的东西来获得玩家的分数。这应该是一个接口，所以测试时我们可以创建一个简单的存根来测试代码而无需实现任何真实的存储机制。

  	  - 对于 POST，我们可以 监听 PlayerStore 的调用以确保它能正确存储玩家。我们的存储实现不会与检索相关联。

	  - 为了尽快让代码可运行，我们可以先在内存中写一个非常简单的实现，然后我们可以实现一个任何喜欢的存储机制。
*/

/*
## http.HandlerFunc

  之前我们探讨过 Handler 接口是为创建服务器而需要实现的。一般来说，我们通过创建 struct 来实现接口。然而，struct 的用途是用于存储数据，但是目前没有状态可存储，因此创建一个 struct 感觉不太对。

  HandlerFunc(https://golang.org/pkg/net/http/#HandlerFunc) 可以让我们避免这样。

  > HandlerFunc 类型是一个允许将普通函数用作 HTTP handler 的适配器。如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是一个调用 f 的 Handler。
*/

/*
重构 1：通过将分数检索分离为函数来简化 PlayerServer

重构 2： 改变架构 将 PlayerServer 改成 struct

重构 3： 区分 POST 与 GET 分别调用不同的辅助函数
*/
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)

	}

	/* 重构 二： 修改 */
	// if r.Method == http.MethodPost {
	// 	w.WriteHeader(http.StatusAccepted)
	// 	return
	// }

	// player := r.URL.Path[len("/players/"):]

	// score := p.store.GetPlayerScore(player)
	// if score == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// }

	// fmt.Fprint(w, score)

	/* 重构 一 修改 */
	// if player == "Pepper" {
	// 	fmt.Fprint(w, "20")
	// 	return
	// }

	// if player == "Floyd" {
	// 	fmt.Fprint(w, "10")
	// 	return
	// }
}

/* ServeHTTP 辅助函数 */

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	// if name == "Pepper" {
	// 	return "20"
	// }

	// if name == "Floyd" {
	// 	return "10"
	// }
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.wincalls = append(s.wincalls, name)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores   map[string]int
	wincalls []string
}
