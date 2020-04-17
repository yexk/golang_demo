package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yexk/src/api/server"
	"github.com/yexk/src/mytest"
)

func main() {
	fmt.Println("starting main ....")
	// mysql
	// mysql()

	// redis
	// redis()

	// http
	serverInit()
}

func serverInit() {
	fmt.Println("starting http server ....")
	server.Webserver(8080)
}

func redis() {
	mytest.Demo1()
}

func mysql() {
	mytest.Demo2()
}
