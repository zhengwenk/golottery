package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"strings"
)

var conn *sql.DB

func GetConn() *sql.DB {

	if (conn != nil) {
		return conn;
	}

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", "root", "", "golang")

	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Println("get mysql conn error, ", err)
		return nil
	}

	return conn
}

func Save(tableName string, data map[string]interface{}) (int64, error) {

	sqlStr := "INSERT INTO %s(%s) VALUES(%s)"

	var fields []string
	var placeholder []string

	values := make([]interface{}, 0)

	for f, v := range data {
		fields = append(fields, "`" + f + "`")
		placeholder = append(placeholder, "?")
		values = append(values, v)
	}

	sqlStr = fmt.Sprintf(sqlStr, tableName, strings.Join(fields, ","), strings.Join(placeholder, ","))

	db := GetConn()
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)

	if (err != nil) {
		log.Println("prepare insert sql error, ", err)
		return 0, err
	}

	log.Println("sql:", sqlStr)

	result, err := stmt.Exec(values...)

	if err != nil {
		return 0, err
	}

	ins_id, _ := result.LastInsertId()
	log.Println("ins_id:", ins_id)
	return ins_id, nil
}

