package datamodels

import (
	"chihuo/comm"
	"database/sql"
	"strconv"
)

//数据库数据结构
type Manu struct {
	ID    int
	Name  string
	Type  int
	Count int
}

//向外提供的接口
type IManu interface {
	Conn() error
	Insert(m *Manu) error
	DeleteOne(m *Manu) error
	SelectAll() (manuArray []*Manu, err error)
	UpdateCount(user *Manu) (err error)
}

//实现接口的数据结构
type UseManu struct {
	table     string
	mysqlConn *sql.DB
}

//创建一个UseManu实例
func NewUseMane(table string, db *sql.DB) IManu {
	return &UseManu{table: table, mysqlConn: db}
}

//连接数据库
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

//搜索所有数据
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
		manu.Name = v["Name"]
		manu.ID, err = strconv.Atoi(v["ID"])
		if err != nil {
			return nil, err
		}
		manu.Count, err = strconv.Atoi(v["Count"])
		if err != nil {
			return nil, err
		}
		manu.Type, err = strconv.Atoi(v["Type"])
		if err != nil {
			return nil, err
		}
		manuArray = append(manuArray, manu)
	}
	return
}

//插入一条新数据
func (u *UseManu) Insert(user *Manu) (err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sql := "INSERT " + u.table + " SET Name=?,Type=?,Count=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Type, user.Count)
	if err != nil {
		return err
	}
	return
}

//删除一条数据
func (u *UseManu) DeleteOne(user *Manu) (err error) {
	if err = u.Conn(); err != nil {
		return
	}

	sql := "DELETE FROM " + u.table + " WHERE ID=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return err
	}
	return
}

func (u *UseManu) UpdateCount(user *Manu) (err error) {
	if err = u.Conn(); err != nil {
		return
	}

	sql := "UPDATE " + u.table + " SET Count=? WHERE ID=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Count, user.ID)
	if err != nil {
		return err
	}
	return
}
