package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//a := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", "root", "", "localhost", "3306", "test", "utf-8")

	//panic(a)

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/product?charset=utf8")
	checkErr(err)
	/*insertSql := "INSERT INTO `tp_rabbitmq_test` SET body=?"
	//插入数据
	res, err := InsertData(db, insertSql)
	checkErr(err)
	//获取此次插入的id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	updateSql := "update `tp_rabbitmq_test` set body=? where id=?"

	bodyM := make(map[string]interface{})
	bodyM["body"] = "Go update test"
	bodyM["id"] = id

	res, err = updateData(db, updateSql, bodyM)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)*/

	//查询语句
	selectSql := "SELECT * FROM `product`"

	//查询数据
	rows, err := queryData(db, selectSql)
	checkErr(err)

	for rows.Next() {
		//var id int
		var body string
		var create_time string
		var update_time string
		err = rows.Scan(&body, &create_time, &update_time)
		checkErr(err)
		//fmt.Println(id)
		fmt.Println(body)
		fmt.Println(create_time)
		fmt.Println(update_time)
	}

	/*//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)*/

	db.Close()

}

/**
插入数据库
*/
func InsertData(db *sql.DB, sql string) (sql.Result, error) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	//执行插入
	res, err := stmt.Exec("Go test")
	return res, err
}

/**
更新数据
*/
func updateData(db *sql.DB, sql string, body map[string]interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	//执行更新
	res, err := stmt.Exec(body["body"], body["id"])
	return res, err
}

/**
查询语句
*/
func queryData(db *sql.DB, sql string) (*sql.Rows, error) {
	rows, err := db.Query(sql)
	return rows, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
