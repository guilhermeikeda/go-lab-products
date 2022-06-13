package use_cases

import (
	"errors"
	"go-lab/entities"
	"go-lab/repositories"
)

type GetProductUseCase interface {
	GetById(requestId int64) (entities.Product, error)
}

type getProductUseCaseImpl struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewGetProductUseCase(productRepository repositories.ProductRepositoryInterface) GetProductUseCase {
	return getProductUseCaseImpl{
		productRepository: productRepository,
	}
}

func (g getProductUseCaseImpl) GetById(requestId int64) (entities.Product, error) {
	products := g.productRepository.GetAllProducts()

	for _, product := range products {
		if product.Id == requestId {
			return product, nil
		}
	}

	return entities.Product{}, errors.New("PRODUCT_NOT_FOUND")
}
