package main

import (
	"bilibili/check"
	"bilibili/find"
	"bilibili/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)


//各种接口
func main(){
	r:=gin.Default()
	//r.GET("/bilibili/1.enter/load.html", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "http://localhost:63342/bilibili/1.enter/load.html")
	//	//c.JSON(200,gin.H{"status": http.StatusOK, "message":"nb"})
	//})
	r.POST("/register", user.Register)//注册 OK
	r.POST("/login", Login)//登录  OK
	r.POST("/logout", Logout) //不知道能不能用    OK？
	r.POST("/GetComment", GetComment)//评论储存  OK
	r.POST("/SendComment", SendComment)//评论发送  OK
	r.POST("/DelComment", DelComment)//删评论  OK
	r.POST("/SendMsg", SendMsg)//弹幕发送  OK
	r.POST("/GetMsg", GetMsg)//弹幕储存  OK
	r.POST("/add", Add)//加好友
	r.POST("/delete", Delete)//删除好友
	r.POST("/Addvideo", Addvideoes)//储存视频
	r.POST("/Sendvideoes", Sendvideoes)//发送视频
	r.POST("/Bingoes", Bingoes)//视频点赞
	r.POST("/Rewards", Rewards)//视频投币
	r.POST("/collection", Collection)// 视频收藏
	r.POST("/Delvideo", Deletevideos)//删除视频
	r.POST("/person", Person)//个人主页 OK
	r.POST("/FIND", Find)//寻找
	//w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//w.Header().Set("content-type", "application/json") //返回数据格式是json
	////值可以设为星号,也可以指定具体主机地址,可设置多个地址用逗号隔开,设为指定主机地址第三项才有效
	//ctx.Header("Access-Control-Allow-Origin", "http://192.168.5.106:8080")
	////允许请求头修改的类容
	//ctx.Header("Access-Control-Allow-Headers", "Content-Type")
	////允许使用cookie
	//ctx.Header("Access-Control-Allow-Credentials", "true")
	//http.HandleFunc("/", Entrance)
	//http.HandleFunc("/ajax", TestCrossOrigin)
	//http.ListenAndServe(":8000", nil)
	_ = r.Run()
}


/*func Entrance(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("templates/ajax.html")
	t.Execute(w, nil)
}

func TestCrossOrigin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var message Message
		message.Name = "benben_2015"
		message.Msg = "success"

		result, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		ResponseWithOrigin(w, r, http.StatusOK, result)
		return
	}*/


////注册(已检测)OK
//func Register(c *gin.Context){
//	username  := c.PostForm("username")
//	password1 := c.PostForm("password1")
//	password  := c.PostForm("password")
//	fmt.Println("user:" + username + password + password1)
//	if check.Check(password, password1, username) {
//		if user.Signup(username, password){
//			c.JSON(500,gin.H{"status":http.StatusInternalServerError , "message":"数据库Insert报错"})
//		}else {
//			c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
//		}
//	}else{
//		c.JSON(500,gin.H{"message":"两次输入密码不一致或用户名为空,请核对后重新输入"})
//	}
//}


//登录（已检测）OK
func Login(c *gin.Context) {
	id0 := c.PostForm("name")
	password := c.PostForm("password")
	id,_ := strconv.Atoi(id0)
	if user.Signin(id ,password) {
		c.SetCookie(/*1*/"username", id0,
			/*2*/6000,
			/*3*/"/",
			/*4*/"localhost",
			/*5*/false,
			/*6*/false)
		    //第一个参数为 cookie 名；
		    // 第二个参数为 cookie 值；
		    // 第三个参数为 cookie 有效时长；
		    // 第四个参数为 cookie 所在的目录；
		    // 第五个为所在域，表示我们的 cookie 作用范围；
		    // 第六个表示是否只能通过 https 访问；
		    // 第七个表示 cookie 是否支持HttpOnly属性。
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败"})
	}
}


//退出登录
func Logout(c *gin.Context)  {
	id,_ := c.Cookie("username")
	c.SetCookie(/*1*/"username", id,
		/*2*/0,
		/*3*/"/",
		/*4*/"localhost",
		/*5*/false,
		/*6*/false)
	//第一个参数为 cookie 名；第二个参数为 cookie 值；
	// 第三个参数为 cookie 有效时长；第四个参数为 cookie 所在的目录；
	// 第五个为所在域，表示我们的 cookie 作用范围；
	// 第六个表示是否只能通过 https 访问；第七个表示 cookie 是否支持HttpOnly属性。
	username,_ := c.Cookie("username")
	fmt.Println(username,123)
	if check.Checkcookie(id, c) {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "退出成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message":"退出失败"})
	}
}


//发送弹幕OK
func SendMsg(c *gin.Context) {
	cookie, err := c.Request.Cookie("username")
	id    := cookie.Value
	uid,_ := strconv.Atoi(id)
	fid0  := c.PostForm("friend")
	vid0  := c.PostForm("video")
	fid,_ := strconv.Atoi(fid0)
	vid,_ := strconv.Atoi(vid0)
	fmt.Println(uid, "message", fid, vid)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	msgs := user.SendMessage(vid, uid, fid)
	fmt.Println(msgs)
	if msgs != nil{
		c.JSON(200,gin.H{"status": http.StatusOK,"message:":msgs})
	}else {
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"留言为空or发送失败"})
	}
}


//储存弹幕OK
func GetMsg(c *gin.Context){
	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	u1id,_ := strconv.Atoi(id)
	fmt.Println("username" + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}

	vid ,_  := strconv.Atoi(c.PostForm("vid"))
	u2id,_  := strconv.Atoi(c.PostForm("u2id")) //strconv.Atoi(c.PostForm(""))为string类型转int （未测试是否成功）
	message := c.PostForm("message")

	if user.GetMessage(vid, u1id, u2id, message) {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})
	} else {
		c.JSON(200, gin.H{"内容": message, "用户名": id})
	}
}


//获取评论
func GetComment(c *gin.Context){
	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username: " + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}

	vid,_  := strconv.Atoi(c.PostForm("vid"))
	fid,_  := strconv.Atoi(c.PostForm("fid"))
	pid,_  := strconv.Atoi(c.PostForm("pid"))//strconv.Atoi(c.PostForm(""))为string类型转int （未测试是否成功）
	message:= c.PostForm("message")
	fmt.Println(vid, fid, pid, message)

	if user.Getcomment(vid, uid, fid, pid, message) {
		c.JSON(200, gin.H{"内容": message, "用户名": id})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})
	}
}


//发送评论
func SendComment(c *gin.Context)  {
	vid0 := c.PostForm("vid")
	uid0 := c.PostForm("uid")
	fid0 := c.PostForm("fid")
	pid0 := c.PostForm("pid")
	vid,_:= strconv.Atoi(vid0)
	uid,_:= strconv.Atoi(uid0)
	fid,_:= strconv.Atoi(fid0)
	pid,_:= strconv.Atoi(pid0)
	fmt.Println(vid, uid, fid)
	msgs := user.Sendcomment(vid, uid, fid, pid)
	fmt.Println(msgs)
	if msgs != nil{
		c.JSON(200,gin.H{"status": http.StatusOK,"message:":msgs})
	}else {
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"留言为空or发送失败"})
	}
}


//删除评论
func DelComment(c *gin.Context){
	cookie,err:=c.Request.Cookie("username")
	id    := cookie.Value
	uid,_ := strconv.Atoi(id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	ip0  := c.PostForm("ip")
	ip,_ := strconv.Atoi(ip0)
	fmt.Println(uid)
	if user.Delcomment(uid, ip){
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "删除评论成功"})
	}else {
		c.JSON(500,gin.H{"status":http.StatusInternalServerError , "message":"数据库Insert报错"})
	}
}


//发送视频 有问题
func Sendvideoes(c *gin.Context)  {
	video := c.PostForm("video")
	videoes := "E:/gotest/"+ video+ ".png"
	fmt.Println(videoes)
	fp, err := os.Open(videoes)
	if err != nil{
		c.JSON(403, gin.H{"status": http.StatusBadRequest, "message": "error"} )
	}else {
		return
	}
	defer fp.Close()
	data := make([]byte, 100)
	count, _ := fp.Read(data)
	fmt.Println(data[:count])
	c.JSON(500, gin.H{"status": http.StatusOK, "ip": videoes, "message": data[:count] })
}


//添加视频OK
func Addvideoes(c *gin.Context){
	//存视频
	f, err := c.FormFile("f1")
	if err != nil{
		c.JSON(502, gin.H{"status": http.StatusBadGateway,"message":"上传失败"})
	}
	dst := path.Join("E:/gotest", f.Filename)
	c.SaveUploadedFile(f, dst)
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB 我也没懂这个。。。

	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	u1id,_ := strconv.Atoi(id)
	fmt.Println("username" + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	video        := c.PostForm("video")
	introduction := c.PostForm("introduction")
	ip           := c.PostForm("ip")
	fmt.Println("video:"+ video + introduction + ip /*+ time*/)
	if check.Checkvideo(video, introduction, ip) {
		if user.Addvideo(video, introduction, ip, u1id){
			c.JSON(500,gin.H{"status":http.StatusInternalServerError , "message" : "数据库Insert报错"})
		}else {
			c.JSON(200, gin.H{"status": http.StatusOK, "message": "发布视频成功"})
		}
	}else{
		c.JSON(500,gin.H{"message":"名字或内容为空,请核对后重新输入"})
	}
}


//删除视频OK
func Deletevideos(c *gin.Context){
	cookie, err := c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username" + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	vid0  := c.PostForm("videoid")
	video := c.PostForm("video")
	vid,_ := strconv.Atoi(vid0)
	fmt.Println("video: " + video)
		if user.Deletevideo(uid, vid){
			c.JSON(200, gin.H{"status": http.StatusOK, "message": "删除视频成功"})
		}else {
			c.JSON(500,gin.H{"status":http.StatusInternalServerError , "message":"数据库Insert报错"})
		}
	}


//视频点赞
func Bingoes(c *gin.Context)  {
	cookie, err := c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username" + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	if user.Bingo(uid){
		c.JSON(500, gin.H{"cstatus": http.StatusOK, "message: ":"点赞成功"})
	}else {
		c.JSON(403, gin.H{"cstatus": http.StatusForbidden, "message: ":"点赞失败"})
	}
}


//视频投币
func Rewards(c *gin.Context)  {
	cookie, err := c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username", uid)
	num0  := c.PostForm("reward")
	vid0  := c.PostForm("video")
	num,_ := strconv.Atoi(num0)
	vid,_ := strconv.Atoi(vid0)
	fmt.Println(vid + num)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	if user.Rewards(vid, num){
		c.JSON(500, gin.H{"cstatus": http.StatusOK, "message: ":"投币成功"})
	}else {
		c.JSON(403, gin.H{"cstatus": http.StatusForbidden, "message: ":"点赞失败"})
	}
}


//视频收藏
func Collection(c *gin.Context)  {
	cookie, err := c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden, "message": "cookie读取失败"})
		return
	}
	vid0  := c.PostForm("video")
	vid,_ := strconv.Atoi(vid0)
	if user.Collections(uid, vid){
		c.JSON(500,gin.H{"status": http.StatusOK, "message": "收藏成功"})
	}else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "收藏失败"})
	}
	fmt.Println(uid)
}


//寻找OK
func Find(c *gin.Context)  {
	finder := c.PostForm("message")
	//finder := c.PostForm("find")
	findint,_ := strconv.Atoi(finder)
	fmt.Println(finder, findint)
	users := find.Findu(findint)
	if users != nil{
		c.JSON(200,gin.H{"status":http.StatusOK,"您所查找的人是:":users })
	}
	userer := find.Fuser(finder)
	if userer  != nil {
		c.JSON(200,gin.H{"status":http.StatusOK,"您所查找的人是:":userer})
	}
	videoes := find.Findv(findint)
	if videoes != nil {
		c.JSON(200, gin.H{"status": http.StatusOK, "您所查找的电影是:": videoes})
	}
	video   := find.Fvideo(finder)
	if video   != nil {
		c.JSON(200, gin.H{"status": http.StatusOK, "您所查找的电影是:": video  })
	}
}


//添加好友OK
func Add(c *gin.Context)  {
	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username" + id)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	 friends/*, err1 */  := c.PostForm("friends")
	      fid0/*, err2 */  := c.PostForm("fid")
	fmt.Println("username:", uid)
	/*if err0 !=nil{
		c.JSON(500,gin.H{"status":http.StatusForbidden,"username":"cookie读取失败"})
		return
	}
	if err1 !=nil{
		c.JSON(500,gin.H{"status":http.StatusForbidden,"username":"cookie读取失败"})
		return
	}
	if err2 !=nil{
		c.JSON(500,gin.H{"status":http.StatusForbidden,"username":"cookie读取失败"})
		return
	}*/
	fid, _ := strconv.Atoi(fid0)
	if user.Add(uid, fid){
		c.JSON(200,gin.H{"好友":friends + "添加成功"})
	}else{
		c.JSON(403,gin.H{"status": http.StatusForbidden, "message": "添加失败or好友达到上限"})
	}
}


//删除好友OK
func Delete(c *gin.Context)  {
	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username", uid)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}
	friend := c.PostForm("fid")
    fid,_  := strconv.Atoi(friend)
    fmt.Println(fid)
	if user.Deletef(fid,  uid){
		c.JSON(200,gin.H{"好友":id + "删除成功"})
	}else{
		c.JSON(403,gin.H{"status": http.StatusForbidden, "message": "删除失败or好友不存在"})
	}
}


//个人信息
func Person(c *gin.Context)  {
	cookie, err:=c.Request.Cookie("username")
	id := cookie.Value
	uid,_ := strconv.Atoi(id)
	fmt.Println("username", uid)
	if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden, "message": "cookie读取失败"})
		return
	}
	user, video, friend:= user.Pn(uid)
	c.JSON(200, gin.H{"status": http.StatusOK, "message": user  })
	c.JSON(200, gin.H{"status": http.StatusOK, "message": video })
	c.JSON(200, gin.H{"status": http.StatusOK, "message": friend})
	fmt.Println(user, video)
}
