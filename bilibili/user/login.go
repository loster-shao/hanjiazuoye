package user

import "database/sql"


//登录   已检测
func Signin(username int, password string)bool  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(" SELECT password FROM users where id=?",username)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var pw string
		err = rows.Scan(&pw)
		if pw == password{
			return true
		}
	}
	return false
}
