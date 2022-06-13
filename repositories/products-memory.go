package repositories

import "go-lab/entities"

type ProductRepositoryMemory struct {
}

func NewProductRepositoryMemory() ProductRepositoryMemory {
	return ProductRepositoryMemory{}
}

func (productRepository ProductRepositoryMemory) GetAllProducts() []entities.Product {
	products := []entities.Product{
		{

			Id:          2,
			Name:        "MCA",
			Description: "Merchant Cash Advance",
		},
		{
			Id:          1,
			Name:        "Split",
			Description: "Split your payment",
		},
	}

	return products
}
