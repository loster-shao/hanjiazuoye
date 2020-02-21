package user

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//添加好友   已检测
func Add(uid, fid int)bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	//查询好友栏
	stmt1, _ := db.Query("SELECT f1, f2, f3, f4, f5, f6, f7, f8, f9, f10 FROM users WHERE id=?", uid)
	for stmt1.Next() {
		//err=stmt1.Scan(&users.friend1,&szsbilibili.friend2 ,&szsbilibili.friend3 )
		var friend1, friend2, friend3, friend4, friend5,
		friend6, friend7, friend8, friend9, friend10 int
		err = stmt1.Scan(&friend1, &friend2, &friend3, &friend4, &friend5,
			&friend6, &friend7, &friend8, &friend9, &friend10)
		fmt.Println(friend1, friend2, friend3, friend4, friend5,
			friend6, friend7, friend8, friend9, friend10)
		if friend1 == 0{
			stmt, _ :=db.Prepare("UPDATE users SET f1=? WHERE id=?")
			_, err := stmt.Exec(fid, uid)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true} else if friend2 == 0 {
			stmt, _ :=db.Prepare("UPDATE users SET f2=? WHERE id=?")
			_, err := stmt.Exec(fid, uid)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true} else if friend3 == 0{
			stmt, _ :=db.Prepare("UPDATE users SET f3=? WHERE id=?")
			_, err := stmt.Exec(fid, uid)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true
		}
	}
	//if friend1 != "" {
	//	db.Exec("UPDATE users SET friend1=? WHERE username=?", friends, username)
	//	return true
	//} else if friend2 != "" {
	//	db.Exec("UPDATE users SET friend2=? WHERE username=?", friends, username)
	//	return true
	//} else if friend3 != "" {
	//	db.Exec("UPDATE users SET friend3=? WHERE username=?", friends, username)
	//	return true
	//}
	return false
}
//判断好友栏是否已满？？？表示这个不知道在哪？


//删除好友  未测试（应该可以）    到时候换switch函数，bingo
func Deletef(friend, id int)bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	//查询好友栏
	stmt1, err1 := db.Query("SELECT f1, f2, f3, f4, f5, f6, f7, f8, f9, f10 FROM users WHERE id=?", id)
	if err1 != nil {
		panic(err)
		fmt.Println("查询失败，请重试")
	}
	for stmt1.Next() {

		var friend1, friend2, friend3, friend4, friend5,
		friend6, friend7, friend8, friend9, friend10 int

		err = stmt1.Scan(&friend1, &friend2, &friend3, &friend4, &friend5,
			&friend6, &friend7, &friend8, &friend9, &friend10)

		fmt.Println(friend1, friend2, friend3, friend4, friend5,
			friend6, friend7, friend8, friend9, friend10)

		if friend1 == friend{
			stmt, _ :=db.Prepare("UPDATE users SET f1=? WHERE id=?")
			_, err := stmt.Exec(0, id)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true} else if friend2 == friend {
			stmt, _ :=db.Prepare("UPDATE users SET f2=? WHERE id=?")
			_, err := stmt.Exec(0, id)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true} else if friend3 == friend{
			stmt, _ :=db.Prepare("UPDATE users SET f3=? WHERE id=?")
			_, err := stmt.Exec(0, id)
			if err != nil {
				fmt.Println("fail to insert")
				return false
			}
			return true
		}
	}
	return false
}