package find

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"user"`
}

type Video struct {
	Id    int    `json:"id"`
	Video string `json:"video"`
}


//通过id搜索用户
func Findu(id int) []User {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt, err1 := db.Query("select id, user from users where id=?", id)
	if err1 != nil{
		fmt.Println("select db failed,err:",err)
	}
	fmt.Println(stmt)
	var userers []User
	for stmt.Next(){
		var userer User
		err = stmt.Scan(&userer.Id, &userer.Username)
		fmt.Println("1: ", userer)
		userers = append(userers, userer)
	}
	return userers
}


//通过模糊字段搜索用户(未实现模糊查询)
func Fuser(name string) []User {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt, err1 := db.Query("SELECT id, user FROM users WHERE user LIKE ?","%" + name + "%")
	if err1 != nil{
		fmt.Println("select db failed,err:",err)
	}
	var userers []User
	for stmt.Next(){
		var userer User
		err = stmt.Scan(&userer.Id, &userer.Username)
		fmt.Println("2: ", userer)
		userers = append(userers, userer)
	}
	return userers
}


//通过id搜索视频
func Findv(id int) []Video {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt, err := db.Query("select id, video from videoes where id=?", id)
	if err != nil{
		fmt.Println("select db failed,err:",err)
	}
	var videoes []Video
	for stmt.Next(){
		var video Video
		err = stmt.Scan(&video.Id, &video.Video)
		fmt.Println("3: ", video)
		videoes = append(videoes, video)
	}
	return videoes
}


//通过模糊字段搜索视频(未实现模糊查询)
func Fvideo(name string) []Video {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt, err1 := db.Query("SELECT id, video FROM videoes WHERE video LIKE ?","%" + name + "%")
	if err1 != nil{
		fmt.Println("select db failed,err:",err)
	}
	var videoes []Video
	for stmt.Next(){
		var video Video
		err = stmt.Scan(&video.Id, &video.Video)
		fmt.Println("4: ", video)
		videoes = append(videoes, video)
	}
	return videoes
}


//留言查询
func Findm(id int)  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	message, err := db.Query("select user from messages where id=?", id)
	if err != nil{
		fmt.Println("select db failed,err:",err)
		return
	}
	fmt.Println(message)
	return
}
