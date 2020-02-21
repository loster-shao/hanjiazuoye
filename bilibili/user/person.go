package user

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
)

type Psn struct {
	Id   int    `json:"id"`
	User string `json:"User"`
}

type Friend struct {
	F1   int    `json:"F1"`
	F2   int    `json:"F2"`
	F3   int    `json:"F3"`
	F4   int    `json:"F4"`
	F5   int    `json:"F5"`
	F6   int    `json:"F6"`
	F7   int    `json:"F7"`
	F8   int    `json:"F8"`
	F9   int    `json:"F9"`
	F10  int    `json:"F10"`
}

type Video struct {
	Ip   string `json:"Ip"`
	Vd   string `json:"video"`
	Intr string `json:"introdution"`
	Re   string `json:"reward"`
	Bing string `json:"bingo"`
	Time string `json:"time"`
}

func Pn(id int) ([]Psn, []Video, []Friend){
	fmt.Println(id)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
		fmt.Println("0")
	}
	stmt, err1 := db.Query("SELECT id, user FROM users WHERE id=?", id)
	if err1 != nil {
		panic(err1)
		fmt.Println(err1, "1")
	}
	var person []Psn
	for stmt.Next() {
		var pn Psn
		stmt.Scan(&pn.Id, &pn.User)
		person = append(person, pn)
	}

	stmt1, err2 := db.Query("SELECT ip, intro, video, bingo, reward, time FROM videoes WHERE uid=?", id)
	if err2 != nil {
		fmt.Println(err,"2")
	}
	var video []Video
	for stmt1.Next(){
		var pn Video
		err = stmt1.Scan(&pn.Ip, &pn.Intr, &pn.Vd, &pn.Bing, &pn.Re, &pn.Time) //问题：多个值会不会直接加入
		video = append(video, pn)
	}

	stmt2, err3 := db.Query("SELECT f1, f2, f3, f4, f5, f6, f7, f8, f9, f10 FROM users WHERE id=?", id)
	if err3 != nil {
		panic(err3)
		fmt.Println(err3, "1")
	}
	var friend []Friend
	for stmt2.Next() {
		var pn Friend
		stmt2.Scan(&pn.F1, &pn.F2, &pn.F3, &pn.F4, &pn.F5, &pn.F6, &pn.F7, &pn.F8, &pn.F9, &pn.F10)
		friend = append(friend, pn)
	}
	return person, video,friend
}
