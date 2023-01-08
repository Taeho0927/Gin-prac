package dblayer

import (
	"GoMusicAPI/models"
)

type DBLayer interface { // 데이터베이스 레이어의 모든 기능 파악 후 모든 기능을 캡슐화 하는 인터페이스 생성
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
