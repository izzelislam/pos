package controller

import (
	"net/http"
	"pos/config"
	"pos/model/vendor"
	"pos/utils"

	"github.com/labstack/echo/v4"
)

func GetVendor(c echo.Context) error {
	var vendor []vendor.Vendor
	config.DB.Find(&vendor)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "getdata success",
		Data: vendor,
	})
}

func ShowVendor(c echo.Context) error {
	id := c.Param("id")

	vendor := vendor.Vendor{}

	config.DB.First(&vendor, id)
	
	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: vendor,
	})
}

func StoreVendor(c echo.Context) error {
	var vendor vendor.Vendor
	
	c.Bind(&vendor)

	config.DB.Create(&vendor)

	return c.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "created",
		Data: vendor,
	})
}

func UpdateVendor(c echo.Context) error {
	id := c.Param("id")
	vendor := vendor.Vendor{}
	c.Bind(&vendor)

	config.DB.Where("id = ?", id).Updates(&vendor)

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "updated",
		Data: vendor,
	})
}

func DeleteVendor(c echo.Context) error {
	id := c.Param("id")
	vendor := vendor.Vendor{}

	config.DB.Where("id = ?", id).Delete(&vendor)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "deleted",
		Data: nil,
	})
}