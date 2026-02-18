package models

import "time"

type Customers struct {
	ID     uint     `gorm:"primaryKey"`
	Name   string   `gorm:"size:100;not null"`
	Email  string   `gorm:"size:100;not null;uniqueIndex"`
	Orders []Orders `gorm:"foreignKey:CustomerID;references:ID"`
}

type Orders struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint `gorm:"not null;index"`

	Total_ammount uint   `gorm:"not null"`
	Status        string `gorm:"size:10;not null;default:'pending'"`
	Created_at    time.Time
	Customer      Customers `gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	OrderItems []OrderItems `gorm:"foreignKey:OrderID;references:ID"`
	Payment    *Payments    `gorm:"foreignKey:OrderID;references:ID"`
}

type OrderItems struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
	UnitPrice uint `gorm:"not null"`
	Qty       uint `gorm:"not null"`

	Order    Orders   `gorm:"foreignKey:OrderID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE"`
	Products Products `gorm:"foreignKey:ProductID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE"`
}

type Products struct {
	ID         uint         `gorm:"primaryKey"`
	Name       string       `gorm:"size:100;not null"`
	Price      int          `gorm:"not null"`
	Stock      uint         `gorm:"not null;default:0"`
	OrderItems []OrderItems `gorm:"foreignKey:ProductID;references:ID"`
}

type Payments struct {
	ID      uint   `gorm:"primaryKey"`
	OrderID uint   `gorm:"uniqueIndex;not null"`
	Ammount int    `gorm:"not null"`
	Status  string `gorm:"size:10;not null"`
	Order   Orders `gorm:"foreignKey:OrderID;references:ID"`
}
