package controller

import (
	"net/http"
	"pos/config"
	"pos/model/product"
	"pos/utils"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	var products []product.Product
	config.DB.Preload("Vendor").Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "getdata success",
		Data: products,
	})
}

func ShowProduct(c echo.Context) error {
	id := c.Param("id")

	product := product.Product{}

	config.DB.First(&product, id)
	
	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: product,
	})
}

func StoreProduct(c echo.Context) error {
	var product product.Product
	
	c.Bind(&product)

	config.DB.Create(&product)

	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: product,
	})
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := product.Product{}
	c.Bind(&product)

	config.DB.Where("id = ?", id).Updates(&product)

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "updated",
		Data: product,
	})
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	product := product.Product{}

	config.DB.Where("id = ?", id).Delete(&product)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "deleted",
		Data: nil,
	})
}