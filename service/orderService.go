package service

import (
	"golang-emarket/dto"
	"golang-emarket/models"
	"golang-emarket/repository"

	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(req dto.CreateOrderRequest) (*models.Orders, error)
	GetOrder(id uint) (*models.Orders, error)
}

type orderService struct {
	db   *gorm.DB
	repo repository.OrderRepository
}

func NewOrderService(db *gorm.DB, repo repository.OrderRepository) OrderService {
	return &orderService{db: db, repo: repo}
}

func (s *orderService) CreateOrder(req dto.CreateOrderRequest) (*models.Orders, error) {
	var CreatedID uint
	err := s.db.Transaction(func(tx *gorm.DB) error {
		txRepos := s.repo.WithTx(tx)

		orders := models.Orders{
			CustomerID:    req.CustomerID,
			Status:        "pending",
			Total_ammount: 0,
			OrderItems:    []models.OrderItems{},
		}

		for _, it := range req.Items {
			orders.OrderItems = append(orders.OrderItems, models.OrderItems{
				ProductID: it.ProductID,
				Qty:       it.Qty,
				UnitPrice: 0,
			})
		}
		if err := txRepos.Create(&orders); err != nil {
			return err
		}

		CreatedID = orders.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(CreatedID)
}

func (s *orderService) GetOrder(id uint) (*models.Orders, error) {
	return s.repo.FindByID(id)
}
