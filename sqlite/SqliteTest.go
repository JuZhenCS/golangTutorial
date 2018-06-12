//sqlite是一个嵌入式数据库,嵌入到程序当中的数据库.
//几乎每个发行版都会预装sqlite.
//这是一个sqlite的测试程序.
//首先需要安装sqlite3的驱动
// go get -u github.com/mattn/go-sqlite3
//date:2018-06-11
//by:JuZhen

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//打开数据库,不存在就创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//创建表的命令
	//`在~上面,不是单引号
	sql_table := `
	CREATE TABLE IF NOT EXISTS userinfo(
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NULL,
		departname VARCHAR(64) NULL,
		created DATE NULL
	);
	`
	db.Exec(sql_table)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("王大明", "战忽局", "2018-07-03")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Print(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("李小明", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Print(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var departname string
	var created time.Time
	for rows.Next() {
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid, username, departname, created)
	}
	rows.Close()

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
