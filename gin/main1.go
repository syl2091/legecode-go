package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("./Arctime.ico"))
	//全局加载静态文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This data is come from Go background.",
		})
	})

	r.GET("/user/info", func(ctx *gin.Context) {
		userid := ctx.Query("userid")
		username := ctx.Query("username")
		ctx.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	//Restful方式
	r.GET("/user/info/:userid/:username", func(ctx *gin.Context) {
		userid := ctx.Param("userid")
		username := ctx.Param("username")
		ctx.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	//序列化
	r.POST("/json", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)
	})

	r.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	//路由
	r.GET("/test", func(ctx *gin.Context) {
		//重定向
		ctx.Redirect(301, "localhost:8080/index")

	})

	// 404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", nil)
	})

	r.Run(":8080")

}
