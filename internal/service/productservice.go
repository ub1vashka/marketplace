package service

import "github.com/ub1vashka/marketplace/internal/domain/models"

type ProductStorage interface {
	SaveProduct(models.Product) (string, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(string) (models.Product, error)
	UpdateProduct(string, models.Product) error
	DeleteProduct(string) error
}

type ProductService struct {
	stor ProductStorage
}

func NewProductService(stor ProductStorage) ProductService {
	return ProductService{stor: stor}
}

func (pros *ProductService) SaveProduct(product models.Product) (string, error) {
	return pros.stor.SaveProduct(product)
}

func (pros *ProductService) GetAllProducts() ([]models.Product, error) {
	return pros.stor.GetAllProducts()
}

func (pros *ProductService) GetProductByID(productID string) (models.Product, error) {
	return pros.stor.GetProductByID(productID)
}

func (pros *ProductService) UpdateProduct(productID string, product models.Product) error {
	return pros.stor.UpdateProduct(productID, product)
}

func (pros *ProductService) DeleteProduct(productID string) error {
	return pros.stor.DeleteProduct(productID)
}
