package services

import "rest-api-mod02-http-method/models"

var products = []models.Product{
	{ID: 1, ProductName: "Dell vostro", Price: 15500, ItemID: 1, ItemName: "Laptop"},
	{ID: 2, ProductName: "Dell delan", Price: 15900, ItemID: 1, ItemName: "Laptop"},
}

type ProductService struct{}

func (s *ProductService) GetAllProducts() []models.Product {
	return products
}
