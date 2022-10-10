package repository

import (
	"fmt"

	"github.com/enigma-mart/models"
	"github.com/enigma-mart/utils"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *sqlx.DB
}

func (p *productRepository) Create(product *models.Product) error {
	fmt.Println(product)
	_, err := p.db.NamedExec(utils.INSERT_PRODUCT, product)

	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Select(&products, utils.SELECT_PRODUCT_LIST); err != nil {
		return nil, err
	}
	return products, nil

}
func (p *productRepository) GetById(id string) (models.Product, error) {
	var product models.Product
	if err := p.db.Get(&product, utils.SELECT_PRODUCT_BY_ID, id); err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	if _, err := p.db.Exec(utils.DELETE_PRODUCT, id); err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
