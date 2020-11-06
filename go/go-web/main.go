package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong9",
		})

		fmt.Println(c.Request)
		fmt.Println(111)
	})

	err := r.Run(":3000") // 指定端口
	if err != nil {
		fmt.Println(err)
	}
}
