package repository

import (
	"context"
	"golang-emarket/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	WithTx(tx *gorm.DB) OrderRepository
	Create(order *models.Orders, ctx context.Context) error
	FindByID(id uint, ctx context.Context) (*models.Orders, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) OrderRepository {
	return &orderRepository{db: DB}
}

func (r *orderRepository) WithTx(tx *gorm.DB) OrderRepository {
	return &orderRepository{db: tx}
}

func (r *orderRepository) Create(order *models.Orders, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) FindByID(id uint, ctx context.Context) (*models.Orders, error) {
	var o models.Orders
	err := r.db.WithContext(ctx).Preload("Customer").Preload("OrderItems.Products").First(&o, id).Error
	if err != nil {
		return nil, err
	}
	return &o, err
}
