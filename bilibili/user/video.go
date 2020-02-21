package user

import (
	_ "github.com/go-sql-driver/mysql"

	"bilibili/check"
	"database/sql"
	"fmt"
	"time"
)

//添加视频    （已检测）
func Addvideo(video, introduction, ip string,u1id int)bool{
	tm := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(tm)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT videoes SET video=?, intro=?, ip=?, time=?, uid=?")
	if err != nil {
		fmt.Println(err,"..")
		return true
	}
	_ , err = stmt.Exec(video, introduction, ip, tm, u1id)

	if err != nil {
		fmt.Println("fail to insert")
		return true
	}
	return false
}


//删除视频  老版！！！！(表示想用类似日志库，不允许删除，只是做个删除标记！先这样写，到时候再改)  已测试
func Deletevideo(uid, vid int)bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("连接数据库失败，请重试")
	}
	stmt1, err1 := db.Query("SELECT uid FROM videoes WHERE id=?", vid)
	fmt.Println(vid)
	if err1 != nil {
		panic(err)
		fmt.Println("搜索失败，请重试")
	}
	for stmt1.Next() {
		var u1id int
		err = stmt1.Scan(&u1id)
		fmt.Println(u1id)
		if check.Checkid(u1id, uid){
			results, err := db.Exec("DELETE FROM videoes WHERE id=?",vid)
			if err != nil{
				fmt.Println("delete data fail,err:",err)
				return false
			}
			fmt.Println(results)
			return true
		}
	}
	return false
}


//视频点赞`(已检测)
func Bingo(id int)bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		return false
	}
	bingo, err :=db.Query("SELECT bingo FROM videoes where id=?", id)
	if err != nil{
		fmt.Println("select data fail. err:",err)
		return false
	}
	for bingo.Next(){
		var dz int
		num := dz + 1
		err = bingo.Scan(&dz)
		if err != nil {
			fmt.Println("get data failed. error:[%v]", err.Error())
			return false
		}
		fmt.Println("1:", num)
		stmt, err1 :=db.Prepare("UPDATE videoes SET bingo=? where id=?")
		fmt.Println("3:", stmt)
		if err1 != nil{
			fmt.Println("delete data fail,err:",err)
		}
		ss, err2 := stmt.Exec(num, id)
		if err2 != nil{
			fmt.Println("delete data fail,err:",err)
		}
		fmt.Println(ss)
	}
	return true
}


//视频打赏
func Rewards(vid, num int)bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		return false
	}
	bingo, err :=db.Query("SELECT reward FROM videoes where id=?", vid)
	if err != nil{
		fmt.Println("select data fail. err:",err)
		return false
	}
	for bingo.Next(){
		var dz int
		num := dz + num
		err = bingo.Scan(&dz)
		if err != nil {
			fmt.Println("get data failed. error:[%v]", err.Error())
			return false
		}
		fmt.Println("1:", num)
		stmt, err1 :=db.Prepare("UPDATE videoes SET reward=? where id=?")
		fmt.Println("3:", stmt)
		if err1 != nil{
			fmt.Println("delete data fail,err:",err)
		}
		ss, err2 := stmt.Exec(num, vid)
		if err2 != nil{
			fmt.Println("delete data fail,err:",err)
		}
		fmt.Println(ss)
	}
	return true
}

//视频收藏
func Collections(id, vid int)bool{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		return false
	}
	stmt, err1 := db.Prepare("INSERT collection SET uid=?, vid=?")
	if err1 != nil {
		fmt.Println(err,"1")
		return false
	}
	_, err2 := stmt.Exec(id, vid)
	if err2 != nil {
		fmt.Println(err,"fail to insert")
		return false
	}
	return false
}

