package main

import (
	"golang-emarket/controller"
	"golang-emarket/db"
	"golang-emarket/repository"
	"golang-emarket/router"
	"golang-emarket/service"
)

func main() {
	db := db.Database()

	repo := repository.NewOrderRepository(db)

	service := service.NewOrderService(db, repo)

	controller := controller.NewOrderController(service)

	r := router.Route(controller)

	r.Run(":8080")

}
