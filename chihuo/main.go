package main

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/opentracing/opentracing-go/log"

	"chihuo/comm"
)

var sucai []string
var huncai []string

type ChihuoController struct {
	Ctx     iris.Context
	Session *sessions.Session
	UseManu IManu
}

func (c *ChihuoController) Get() mvc.View {
	// sucai = sucai[0:0]
	// huncai = huncai[0:0]
	sucai = []string{0: "胡萝卜", 1: "豆角", 2: "青菜", 3: "土豆", 4: "南瓜", 5: "花菜", 6: "西兰花"}
	huncai = []string{0: "白切鸡", 1: "烧鸭", 2: "卤鸭", 3: "菠菜炒鸭", 4: "手撕鸡", 5: "辣子鸡", 6: "牛肉炒芹菜"}
	list, err := c.UseManu.SelectAll()
	//fmt.Println(list)
	if err != nil {
		log.Error(err)
	}

	for _, v := range list {
		a := *v
		//fmt.Println(v.HuncaiName)
		sucai = append(sucai, a.SucaiName)
		huncai = append(huncai, a.HuncaiName)
	}

	return mvc.View{
		Data: iris.Map{
			"sucai":  sucai,
			"huncai": huncai,
		},
		Name: "index.html",
	}
}

func (c *ChihuoController) Post() mvc.View {
	var (
		vegetarian = c.Ctx.FormValue("vegetarian")
		meatDish   = c.Ctx.FormValue("meatDish")
	)
	if vegetarian != "" {
		sucai = append(sucai, vegetarian)
	}
	if meatDish != "" {
		huncai = append(huncai, meatDish)
	}
	manu := &Manu{
		HuncaiName: meatDish,
		SucaiName:  vegetarian,
	}

	err := c.UseManu.Insert(manu)
	c.Ctx.Application().Logger().Debug(err)
	if err != nil {
		log.Error(err)
	}
	return mvc.View{
		Data: iris.Map{
			"sucai":  sucai,
			"huncai": huncai,
		},
		Name: "index.html",
	}
}

type Manu struct {
	HuncaiName string
	SucaiName  string
}

type IManu interface {
	Conn() error
	Insert(m *Manu) error
	SelectAll() (manuArray []*Manu, err error)
}

type UseManu struct {
	table     string
	mysqlConn *sql.DB
}

func NewUseMane(table string, db *sql.DB) IManu {
	return &UseManu{table: table, mysqlConn: db}
}

func (u *UseManu) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, err := comm.NewMysqlConn()
		if err != nil {
			return err
		}

		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "manu"
	}
	return
}

func (u *UseManu) SelectAll() (manuArray []*Manu, err error) {
	if err := u.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT * FROM " + u.table
	rows, err := u.mysqlConn.Query(sql)
	//fmt.Println(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := comm.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}

	for _, v := range result {
		manu := &Manu{}
		//comm.DataToStructByTagSql(v, manu)
		//fmt.Println(v)
		manu.HuncaiName = v["huncaiName"]
		manu.SucaiName = v["sucaiName"]
		manuArray = append(manuArray, manu)
	}
	return
}

func (u *UseManu) Insert(user *Manu) (err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sql := "INSERT " + u.table + " SET huncaiName=?,sucaiName=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.HuncaiName, user.SucaiName)
	if err != nil {
		return err
	}
	return
}

func (c *ChihuoController) GetCaidan() mvc.View {
	number1 := rand.Intn(len(sucai))
	number2 := rand.Intn(len(huncai))
	targetOne := sucai[number1]
	targetTwo := huncai[number2]

	return mvc.View{
		Data: iris.Map{
			"sucai":  targetOne,
			"huncai": targetTwo,
		},
		Name: "caidan.html",
	}
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML(".", ".html").Reload(true))
	app.StaticWeb("./image", "./image")

	db, err := comm.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chihuoService := NewUseMane("manu", db)
	chihuoParty := app.Party("/")
	chihuo := mvc.New(chihuoParty)
	chihuo.Register(ctx, chihuoService)
	chihuo.Handle(new(ChihuoController))

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
