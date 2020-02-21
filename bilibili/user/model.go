package user

import (
	"bilibili/check"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)


func Register(c *gin.Context){
	username  := c.PostForm("username")
	password1 := c.PostForm("password1")
	password  := c.PostForm("password")
	fmt.Println("user:" + username + password + password1)
	if check.Check(password, password1, username) {
		if Signup(username, password){
			c.JSON(500,gin.H{"status":http.StatusInternalServerError , "message":"数据库Insert报错"})
		}else {
			c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
		}
	}else{
		c.JSON(500,gin.H{"message":"两次输入密码不一致或用户名为空,请核对后重新输入"})
	}
}


//注册   已检测
func Signup(username, password string)bool{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT users SET user=?, password=?, pid=?, f1=0, f2=0, f3=0")
	if err != nil {
		fmt.Println(err,"..")
		return true
	}
	_ , err = stmt.Exec(username, password, 0)
	if err != nil {
		fmt.Println("fail to insert")
		return true
	}
	return false
}


////删除好友
//func Delete(friends string,username string) bool {
//	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
//	if err != nil {
//		panic(err)
//		fmt.Println("连接数据库失败，请重试")
//	}
//	stmt1,_:=db.Query("SELECT friend1 FROM     WHERE username=?",username)
//	stmt2,_:=db.Query("SELECT friend2 FROM     WHERE username=?",username)
//	stmt3,_:=db.Query("SELECT friend3 FROM     WHERE username=?",username)
//	fmt.Println(stmt1,stmt2,stmt3)
//	if stmt1 == friends {
//		db.Exec("DELETE FROM bilibili WHERE ",friends,username)
//		return true
//	}else {
//		if stmt2 == friends {
//			db.Exec("DELETE FROM szs20191206 WHERE ",friends,username)
//			return true
//		}else {
//			if stmt3 == friends {
//				db.Exec("DELETE FROM szs20191206 WHERE ",friends,username)
//				return true
//			}
//		}
//	}
//	return false
//}
//——————————————————————————————————————————————————未完待续————————————————————————————————————————————————————————————//