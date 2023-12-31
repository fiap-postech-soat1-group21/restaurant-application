package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Product struct type
type Product struct {
	db *gorm.DB
}

// NewProductRepository instantiates product repository
func NewProductRepository(db *gorm.DB) *Product {
	return &Product{db}
}

// Create create a new product record. If record already exists, it does nothing.
func (p *Product) Create(ctx context.Context, pdt *domain.Product) (*domain.Product, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("product").Create(pdt); result.Error != nil {
		return nil, result.Error
	}

	return pdt, nil
}

// Update updates a product record in database.
func (p *Product) Update(ctx context.Context, pdt *domain.Product) error {
	dbFn := p.db.WithContext(ctx)

	filter := pdt.ID

	return dbFn.Table("product").Where(filter).Updates(pdt).Error
}

// Delete remove a product record in database.
func (p *Product) Delete(ctx context.Context, pdt *domain.Product) error {
	dbFn := p.db.WithContext(ctx)

	id := pdt.ID

	return dbFn.Table("product").Where("id = ?", id).Delete(&pdt).Error
}

// List retrives all categories.
func (o *Product) GetProducts(ctx context.Context) (*domain.ProductResponseList, error) {
	dbFn := o.db.WithContext(ctx)

	var count int64
	var products []*domain.Product

	result := dbFn.Table("product").Find(&products).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.ProductResponseList{
		Result: products,
		Count:  count,
	}, nil
}