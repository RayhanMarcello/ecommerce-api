package repository

import (
	"golang-emarket/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	WithTx(tx *gorm.DB) OrderRepository
	Create(order *models.Orders) error
	FindByID(id uint) (*models.Orders, error)
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

func (r *orderRepository) Create(order *models.Orders) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id uint) (*models.Orders, error) {
	var o models.Orders
	err := r.db.Preload("Customers").Preload("OrderItems.Products").Preload("payments").First(&o, id).Error
	if err != nil {
		return nil, err
	}
	return &o, err
}
