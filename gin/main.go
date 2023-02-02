package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/crud-list?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	fmt.Println("db = ", db)
	fmt.Println("err = ", err)

	// 连接池
	s := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	s.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	s.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	s.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	type List struct {
		gorm.Model
		Name    string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
		State   string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
		Phone   string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
		Email   string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
		Address string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	}
	db.AutoMigrate(&List{})
	r := gin.Default()
	r.POST("/user/add", func(ctx *gin.Context) {
		var data List
		// 绑定方法
		err := ctx.ShouldBindJSON(&data)
		// 判断绑定是否有错误
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": "400",
			})
		} else {
			// 数据库的操作
			db.Create(&data) // 创建一条数据
			ctx.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": "200",
			})
		}
	})

	r.DELETE("/user/delete/:id", func(ctx *gin.Context) {
		var data []List
		id := ctx.Param("id")
		db.Where("id = ?", id).Find(&data)
		if len(data) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": 400,
			})
		} else {
			// 操作数据库删除（删除id所对应的那一条）
			// db.Where("id = ? ", id).Delete(&data) <- 其实不需要这样写，因为查到的data里面就是要删除的数据
			db.Delete(&data)

			ctx.JSON(200, gin.H{
				"msg":  "删除成功",
				"code": 200,
			})
		}
	})

	r.POST("/user/update/:id", func(ctx *gin.Context) {
		var data List
		id := ctx.Param("id")
		db.Select("id").Where("id = ?", id).Find(&data)
		if data.ID == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "用户id没有找到",
				"code": 400,
			})
		} else {
			// 绑定一下
			err := ctx.ShouldBindJSON(&data)
			if err != nil {
				ctx.JSON(200, gin.H{
					"msg":  "修改失败",
					"code": 400,
				})
			} else {
				// db修改数据库内容
				db.Where("id = ?", id).Updates(&data)
				ctx.JSON(200, gin.H{
					"msg":  "修改成功",
					"code": 200,
				})
			}
		}
	})

	r.GET("/user/list/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		var data []List
		db.Where("name = ?", name).Find(&data)
		if len(data) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			err := ctx.ShouldBindJSON(&data)
			if err != nil {
				ctx.JSON(200, gin.H{
					"msg":  "查询失败",
					"code": 500,
				})
			} else {
				ctx.JSON(200, gin.H{
					"msg":  "查询成功",
					"code": "200",
					"data": data,
				})
			}
		}
	})

	r.GET("/user/list", func(ctx *gin.Context) {
		var dataList []List
		// 查询全部数据 or 查询分页数据
		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

		// 判断是否需要分页
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}

		offsetVal := (pageNum - 1) * pageSize // 固定写法 记住就行
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}

		// 返回一个总数
		var total int64

		// 查询数据库
		db.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

		if len(dataList) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			ctx.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})
		}
	})

	r.Run(":3000")
}
