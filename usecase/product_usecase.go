package usecase

import (
	"github.com/enigma-mart/models"
	"github.com/enigma-mart/repository"
)

type ProductUseCase interface {
	CreateNewProduct(product *models.Product) error
	FindAllProduct() ([]models.Product, error)
	FindProductByid(id string) (models.Product, error)
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(product *models.Product) error {
	return p.repo.Create(product)
}
func (p *productUseCase) FindAllProduct() ([]models.Product, error) {
	return p.repo.GetAll()
}
func (p *productUseCase) FindProductByid(id string) (models.Product, error) {
	return p.repo.GetById(id)
}
func (p *productUseCase) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	uc := new(productUseCase)
	uc.repo = repo
	return uc
}
