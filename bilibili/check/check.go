package check

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//检测两次密码是否一致,检测用户是否一致
func Check(password string,password1 string,username string)bool  {
	if password1 == password {
		if username != ""{
			return true
		}
		return false
	}
	return false
}


//检测视频是否为空
func Checkvideo(video string,introduction string,ip string)bool  {
	if video !="" && introduction !=""{
		return true
	}
	return false
}


//检测id是否对应
func Checkid(id, u1id int)bool  {
	if id == u1id{
		return true
	}
	return false
}


//检查cookie是否存在
func Checkcookie(user string,c *gin.Context) bool {
	id,_ := c.Cookie("username")
	fmt.Println(id)
	if id == user{
		return true
		fmt.Println("ok")
	}
	return false
}
