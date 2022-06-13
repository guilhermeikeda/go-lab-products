package main

import (
	"fmt"
	"go-lab/repositories"
	"go-lab/use_cases"
)

func main() {
	productRepository := repositories.NewProductRepositoryMemory()

	getProductUseCase := use_cases.NewGetProductUseCase(productRepository)
	product, err := getProductUseCase.GetById(2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Encontrei produto %s", product.Name)
	}
}
