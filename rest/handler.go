package rest

import (
	"GoMusicAPI/dblayer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface { // 핸들러의 모든 메서드를 포함하는 인터페이스를 만듬
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer.DBlayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}
func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
