package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/opentracing/opentracing-go/log"

	"chihuo/comm"
	"chihuo/datamodels"
)

type ChihuoController struct {
	Ctx     iris.Context
	Session *sessions.Session
	UseManu datamodels.IManu
}

func (c *ChihuoController) Get() mvc.View {
	// sucai = []string{0: "胡萝卜", 1: "豆角", 2: "青菜", 3: "土豆", 4: "南瓜", 5: "花菜", 6: "西兰花"}
	// huncai = []string{0: "白切鸡", 1: "烧鸭", 2: "卤鸭", 3: "菠菜炒鸭", 4: "手撕鸡", 5: "辣子鸡", 6: "牛肉炒芹菜"}

	array, err := c.UseManu.SelectAll()
	if err != nil {
		fmt.Println(err)
	}
	//创建两组slice，两个以类型作为区分的slice，最后将两slice里面各一个值返回到页面
	var slice1, slice2 []string
	for _, v := range array {
		if v.Type == 1 {
			slice1 = append(slice1, v.Name)
		}
		if v.Type == 2 {
			slice2 = append(slice2, v.Name)
		}
	}

	return mvc.View{
		Data: iris.Map{
			"sucai":  slice1,
			"huncai": slice2,
		},
		Name: "index.html",
	}
}

func (c *ChihuoController) Post() mvc.Response {
	var (
		vegetarian = c.Ctx.FormValue("vegetarian")
		meatDish   = c.Ctx.FormValue("meatDish")
	)
	if vegetarian != "" {
		vt := insertVar(vegetarian, 1)
		err := c.UseManu.Insert(vt)
		if err != nil {
			fmt.Println(err)
		}
	}
	if meatDish != "" {
		md := insertVar(meatDish, 2)
		err := c.UseManu.Insert(md)
		if err != nil {
			fmt.Println(err)
		}
	}
	return mvc.Response{
		Path: "/",
	}
}

func insertVar(name string, tp int) *datamodels.Manu {
	s := &datamodels.Manu{}
	s.Name = name
	s.Type = tp
	s.Count = 0
	return s
}

// func (c *ChihuoController) PostUpload() mvc.View {

// }

func (c *ChihuoController) GetCaidan() mvc.View {
	array, err := c.UseManu.SelectAll()
	if err != nil {
		fmt.Println(err)
	}
	//创建两组slice，两个以类型作为区分的slice，最后将两slice里面各一个值返回到页面
	var slice1, slice2 []*datamodels.Manu
	for _, v := range array {
		if v.Type == 1 {
			slice1 = append(slice1, v)
		}
		if v.Type == 2 {
			slice2 = append(slice2, v)
		}
	}
	number1 := rand.Intn(len(slice1))
	number2 := rand.Intn(len(slice2))
	targetOne := slice1[number1]
	targetOne.Count += 1
	targetTwo := slice2[number2]
	targetTwo.Count += 1
	c.UseManu.UpdateCount(targetOne)
	c.UseManu.UpdateCount(targetTwo)

	return mvc.View{
		Data: iris.Map{
			"sucai":  targetOne.Name,
			"huncai": targetTwo.Name,
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

	chihuoService := datamodels.NewUseMane("manu", db)
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
