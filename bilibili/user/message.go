package user

import (
	"bilibili/check"
	"database/sql"
	"fmt"
)

//储存弹幕  不需要！！！???
func GetMessage(vid, u1id, u2id int, message string)bool{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		println("连 接 数 据 库 失 败 ，请 重 试 ！ ！ ！")
	}
	stmt, err := db.Prepare("INSERT messages SET message=?, u1id=?, u2id=?, vid=?")
	if err != nil {
		fmt.Println(err,"..")
		return true
	}
	_ , err = stmt.Exec(message, u1id, u2id, vid)
	if err != nil {
		fmt.Println("fail to insert")
		return true
	}
	return false
}


//查看弹幕
type Message struct {
	Video   string    `json:"video"`
	Uname   string    `json:"Uname"`
	Fname   string    `json:"Fname"`
	Message string    `json:"message"`
    Msg     []Message `json:"[]message"`
}

func SendMessage(vid, uid, fid int,/*c *gin.Context*/) []Message  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("发送信息连接数据库失败，请重试")
	}
	stmt, err1 := db.Query("SELECT vid, u1id, u2id, message FROM messages WHERE vid=?", vid)
	var msgs []Message
	for stmt.Next(){
		var msg Message
		stmt.Scan(&msg.Video, &msg.Uname, &msg.Fname, &msg.Message)
		msgs = append(msgs, msg)
	}
	if err1 !=nil{
		fmt.Println(err, "显示错误")
		return nil
	}
	fmt.Println(msgs,"szs")
	return msgs
}


//储存评论
func Getcomment(vid, uid, fid, pid int, message string)bool{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		println("连 接 数 据 库 失 败 ，请 重 试 ！ ！ ！")
	}
	stmt, err := db.Prepare("INSERT msg SET vid=?, uid=?, fid=?, pid=?, message=?")
	if err != nil {
		fmt.Println(err,"..")
		return false
	}
	_ , err = stmt.Exec(vid, uid, fid, pid, message)
	if err != nil {
		fmt.Println("fail to insert")
		return false
	}
	return true
}


//查看评论
type Msg struct {
	video string `json:"video"`
	Uname string `json:"Uname"`
	Fname string `json:"Fname"`
	Message string  `json:"message"`
	message []Msg `json:"[]message"`
}

func Sendcomment(vid, uid, fid, pid int,/*c *gin.Context*/) []Msg  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("发送信息连接数据库失败，请重试")
	}
	stmt, err1 := db.Query("SELECT vid, uid, fid, message FROM msg WHERE vid=? AND pid=?", vid, pid)
	if err1 != nil {
		panic(err)
		fmt.Println("搜索失败")
	}
	var msgs []Msg
	for stmt.Next(){
		var msg Msg
		stmt.Scan(&msg.video, &msg.Uname, &msg.Fname, &msg.Message)
		msgs = append(msgs, msg)
	}
	fmt.Println(msgs,"szs")
	return msgs
}


//删除留言     (未检测)应该没问题
func Delcomment(id, ip int) bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt1, _ := db.Query("SELECT uid FROM msg WHERE id=?", id)
	for stmt1.Next() {
		var u1id int
		err = stmt1.Scan(&u1id)
		if check.Checkid(u1id,ip){
			results, err := db.Exec("DELETE from msg where id=?", id)
			fmt.Println(results)
			if err != nil{
				fmt.Println("delete data fail,err:",err)
				return false
			}
		}
	}
	return false
}
