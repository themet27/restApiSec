package controllers

import (
	"net/http"
	"rest-api-mod02-http-method/services"
)

// ProductController handles HTTP requests related to items
type ProductController struct {
	ProductService services.ProductService
}

// GetAllProductsHandler handles the GET request to retrieve all items
func (c *ProductController) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := c.ProductService.GetAllProducts()
	sendJSONResponse(w, "200", "Items retrieved successfully", products)
}
