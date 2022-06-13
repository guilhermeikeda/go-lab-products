package repositories

import "go-lab/entities"

type ProductRepositoryDatabaseImpl struct {
}

func (productRepository ProductRepositoryDatabaseImpl) GetAllProducts() []entities.Product {
	return []entities.Product{}
}
