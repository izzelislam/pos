package controller

import (
	"net/http"
	"pos/config"
	"pos/model/category"
	"pos/utils"

	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {
	var category []category.Category
	config.DB.Find(&category)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "getdata success",
		Data: category,
	})
}

func ShowCategory(c echo.Context) error {
	id := c.Param("id")

	category := category.Category{}

	config.DB.First(&category, id)
	
	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: category,
	})
}

func StoreCategory(c echo.Context) error {
	var category category.Category
	
	c.Bind(&category)

	config.DB.Create(&category)

	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: category,
	})
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	category := category.Category{}
	c.Bind(&category)

	config.DB.Where("id = ?", id).Updates(&category)

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "updated",
		Data: category,
	})
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	category := category.Category{}

	config.DB.Where("id = ?", id).Delete(&category)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "deleted",
		Data: nil,
	})
}