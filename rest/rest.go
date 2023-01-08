package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error { // Restful API 의 진입점 함수 , 이곳에 HTTP 라우팅을 정의한다
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {})
	r.GET("/promos", func(c *gin.Context) {})
	r.POST("/users/signin", func(c *gin.Context) {})
	r.POST("/users", func(c *gin.Context) {})
	r.POST("/user/:id/signout", func(c *gin.Context) {})
	r.GET("/user/:id/orders", func(c *gin.Context) {})
}
