package rest

import "github.com/gin-gonic/gin"

type HandlerInterface interface { // 핸들러의 모든 메서드를 포함하는 인터페이스를 만듬
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}
