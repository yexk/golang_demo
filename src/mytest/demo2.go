package mytest

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Mysql 数据库配置
type Mysql struct {
	Host     string
	Username string
	Password string
	Port     int
	DBname   string
	Chartset string
}

// Test2 字段
type Test2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Demo2 测试1
func Demo2() {
	var t Test2
	conf := Mysql{
		Host:     "127.0.0.1",
		DBname:   "test",
		Port:     3306,
		Password: "", // password要携带:
		Username: "root",
		Chartset: "utf8",
	}
	// root:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?chartset=%s", conf.DBname, conf.Password, conf.Host, conf.Port, conf.DBname, conf.Chartset))
	if err != nil {
		fmt.Println("mysql connect error :", err)
	}

	fmt.Println("mysql connect successful~")
	rows, err := db.Query("select * from test2")
	for rows.Next() {
		rows.Scan(&t.ID, &t.Name)
	}

	s, err := json.Marshal(t)
	fmt.Println(string(s))
	// 释放
	rows.Close()
}
