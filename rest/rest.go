package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPIWithHandler(address string, h HandlerInterface) error { // Restful API 의 진입점 함수 , 이곳에 HTTP 라우팅을 정의한다
	r := gin.Default()
	h, _ := NewHandler()
	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)
	r.POST("/users/signin", h.SignIn)
	r.POST("/users", h.AddUser)
	r.POST("/user/:id/signout", h.SignOut)
	r.GET("/user/:id/orders", h.GetOrders)

	return r.Run(address)
}

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
