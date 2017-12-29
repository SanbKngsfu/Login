package dao

import (
	"fmt"
	"os"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

func InitDB() (conn mysql.Conn) {
	addr := "127.0.0.1:3306"
	user := "root"
	passwd := "root"
	dbname := "SanbKngsfu"
	conn = CreateCon(addr, user, passwd, dbname)
	if conn == nil {
		panic("conn is nil!")
	}
	conn.Query("set names utf8")
	return
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CheckedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row, mysql.Result) {
	CheckError(err)
	return rows, res
}

func CreateCon(addr string, user string, passwd string, dbname string) (conn mysql.Conn) {
	proto := "tcp"
	conn = mysql.New(proto, "", addr, user, passwd, dbname)
	CheckError(conn.Connect())
	return
}
