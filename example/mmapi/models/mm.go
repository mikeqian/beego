package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikeqian/beego/orm"
)

var (
	MMs map[string]*MM
)

type MM struct {
	Id   string
	Name string
}

func init() {
	MMs = make(map[string]*MM)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	o := orm.NewOrm()

	dr := o.Driver()
	fmt.Println(dr.Name() == "default")
	fmt.Println(dr.Type() == orm.DR_Sqlite)
}
