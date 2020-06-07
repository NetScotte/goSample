package mysql

/*
安装驱动
go get github.com/go-sql-driver/mysql
创建表
create table `user` (
	`id` bigint(20) not null auto_increment,
	`name` varchar(20) default '',
	`age` int(11) default '0',
	primary key (`id`)
)engine=InnoDB auto_increment=1 default charset=utf8mb4

// 该库会卡在sql.go的1177的	ci, err := db.connector.Connect(ctx)那里很久，直到报net dial time out
 */

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"time"
)

var (
	ctx context.Context
	DB *sql.DB
)

type User struct {
	Id int64 `db:"id"`
	// 如果name为空，那么需要使用sql.NullString类型
	Name string `db:"name"`
	Age int `db:"age"`
}

// 测试数据库连接功能
func initDB(dbtype, dsn string) (err error){
	// 只会验证参数，而不会验证dsn的可用性
	DB, err = sql.Open(dbtype, dsn)
	DB.SetMaxOpenConns(100)
	fmt.Printf("success get db, start to ping...\n")
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	fmt.Printf("get ctx: %+v\n", ctx)
	conn, err := DB.Conn(ctx)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		err	= fmt.Errorf("Failed connect to %v with dsn: %v, %+v\n", dbtype, dsn, err)
		return
	}

	return
}

// 测试查询单行功能
func queryRowById(id int) (user User, err error) {
	sqlstr := "select id, name, age from user where id=?"
	row := DB.QueryRow(sqlstr, id)
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil || reflect.ValueOf(user).IsNil(){
		err = fmt.Errorf("Failed queryRowById %v: %v\n", id, err)
	}
	return user, err
}


// 测试查询多行功能
func queryData(id int) (err error) {
	sqlstr := "select name, age from user where id=?"
	row, err := DB.Query(sqlstr, id)
	if err != nil {
		err = fmt.Errorf("Failed queryData by id: %v\n", id)
	}
	defer func() {
		if row != nil {
			err = row.Close()
		}
	}()
	for row.Next() {
		var user User
		err = row.Scan(&user.Name, &user.Age)
		if err != nil || reflect.ValueOf(user).IsNil() {
			err = fmt.Errorf("Failed scan data: %v\n", err)
			return
		}
		fmt.Printf("%+v", user)
	}
	return
}

// 测试插入功能
func insertData(name string, age int) (err error) {
	sqlstr := "insert into user(name, age) values(?, ?)"
	_, err = DB.Exec(sqlstr, name, age)
	if err != nil {
		err = fmt.Errorf("Failed insert data(name: %v, age: %v): %v\n", name, age, err)
		return
	}
	fmt.Printf("success insert\n")
	return
}

// 测试更新功能
func updateData(name string, age int) (err error) {
	sqlstr := "update user set age=? where name=?"
	_, err = DB.Exec(sqlstr, age, name)
	if err != nil {
		err = fmt.Errorf("Failed updateData(name: %v, age: %v): %v\n", name, age, err)
		return
	}
	fmt.Printf("success updateData(name: %v, age: %v)\n", name, age)
	return
}

// 测试删除功能
func deleteData(id int) (err error) {
	sqlstr := "delete from user where id=?"
	_, err = DB.Exec(sqlstr, id)
	if err != nil {
		err = fmt.Errorf("Failed deleteData(id: %v): %v\n", id, err)
		return
	}
	fmt.Printf("success delete\n")
	return
}

// 测试预处理功能
func prepareQueryData(id int) (err error) {
	sqlstr := "select name, age from user where id=?"

	// 需要关闭
	stat, err := DB.Prepare(sqlstr)
	if err != nil {
		err = fmt.Errorf("Failed prepare: %v\n", err)
		return
	}
	defer func(){
		if stat != nil {
			stat.Close()
		}
	}()
	var user User
	row := stat.QueryRow(id)
	err = row.Scan(&user.Name, &user.Age)
	if err != nil {
		err = fmt.Errorf("Failed Scan: %v\n", err)
		return
	}
	fmt.Printf("get user info: %+v\n", user)
	return
}

// 事务的特点： ACID
// 测试事务功能
func transaction() (err error) {
	conn, err := DB.Begin()
	if err != nil {
		conn.Rollback()
		err = fmt.Errorf("Failed begin a transaction: %v\n", err)
		return
	}
	sqlstr := "update user set age = 13 where id = 2"
	_, err = conn.Exec(sqlstr)
	if err != nil {
		conn.Rollback()
		err = fmt.Errorf("Failed exec sql, %+v, %v\n", err, sqlstr)
		return
	}
	sqlstr = "update user set age = 15 where id = 2"
	_, err = conn.Exec(sqlstr)
	if err != nil {
		conn.Rollback()
		err = fmt.Errorf("Failed exec sql, %+v, %v\n", err, sqlstr)
		return
	}
	err = conn.Commit()
	if err != nil {
		err = fmt.Errorf("Failed commit: %v\n", err)
		conn.Rollback()
		return
	}
	fmt.Printf("success exec transaction\n")
	return

}



