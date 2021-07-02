package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	/* 重构 一 ： 创建 NewPlayerServer 用来创建路由规则 */
	/* 重构 二： 在这个函数里面完成所有初始化操作  */
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	router := http.NewServeMux()

// 	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		player := r.URL.Path[len("/players/"):]

// 		switch r.Method {
// 		case http.MethodGet:
// 			p.showScore(w, player)
// 		case http.MethodPost:
// 			p.processWin(w, player)

// 		}
// 	}))
// 	router.ServeHTTP(w, r)
// }

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

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())

	w.WriteHeader(http.StatusOK)
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return []Player{
// 		{"Chris", 20},
// 	}
// }

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)

	}

}

/*
我们更改了 PlayerServer 的第二个属性，删除了命名属性 router http.ServeMux，并用 http.Handler 替换了它；这被称为 嵌入。
高效 Go - 嵌入(https://golang.org/doc/effective_go.html#embedding)
*/
/*
有任何缺点吗？

   你必须小心使用嵌入类型，因为你将公开所有嵌入类型的公共方法和字段。在我们的例子中它是可以的，因为我们只是嵌入了 http.Handler 这个 接口。

   如果我们懒一点，嵌入了 http.ServeMux（混合类型），它仍然可以工作 但 PlayerServer 的用户就可以给我们的服务器添加新路由了，因为 Handle(path, handler) 会公开。

   *****嵌入类型时，真正要考虑的是对你公开的 API 有什么影响。*****

   滥用嵌入最终会污染你的 API，并暴露你的类型的内部信息，这是个常见的错误。
*/
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}
