package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //引用的开源postsql接口。
)

func main() {
	db, err := sql.Open("postgres",
		"host='172.17.0.2' port=5432 "+
			"user=root password=1 dbname=root") //开启连接
	if err != nil {
		fmt.Println("开启数据库失败！", err.Error())
		return
	}
	defer db.Close()
	//创建账号密码表
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS account_password" +
		"(account VARCHAR(80), password VARCHAR(80))") //创建表用来存储账号密码
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result.LastInsertId())
	//插入数据
	stmt, err := db.Prepare("INSERT INTO account_password(account, password) " +
		"VALUES($1, $2)") //插入数据预处理(可以提高效率和防止注入)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	account := "test001"
	password := "test001"
	result, err = stmt.Exec(account, password) //执行插入
	fmt.Println(result.LastInsertId())
	//删除数据
	stmt, err = db.Prepare("DELETE FROM account_password WHERE account=$1") //删除预处理
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec("test002")
	fmt.Println(result.LastInsertId())
	//修改数据
	stmt, err = db.Prepare("UPDATE account_password SET password=$2 WHERE account=$1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec("test001", "password1")
	fmt.Println(result.LastInsertId())
	//读取数据
	rows, err := db.Query("SELECT * FROM account_password")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&account, &password)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(account, password)
	}
}
