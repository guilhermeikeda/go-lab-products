package repositories

import "go-lab/entities"

type ProductRepositoryInterface interface {
	GetAllProducts() []entities.Product
}
