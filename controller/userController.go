package controller

import (
	"encoding/json"
	"net/http"
	"pos/config"
	"pos/model/user"
	"pos/utils"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	var users []user.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status : http.StatusOK,
		Message: "Get data successuly",
		Data	 : users,
	})
}

func CreateUser(e echo.Context) error {

	payload := user.UserRequest{}

	jsonData := make(map[string]interface{})
	err := json.NewDecoder(e.Request().Body).Decode(&jsonData)
	
	payload.Name 		 = jsonData["name"].(string)
	payload.Email 	 = jsonData["email"].(string)
	payload.Password = utils.EncryptPassword([]byte(jsonData["password"].(string)))
	
	if err != nil{
		panic(err)
	}

	config.DB.Create(&payload)

	return e.JSON(http.StatusCreated, utils.BaseResponse{
		Status: http.StatusCreated,
		Message: "Created",
		Data: payload,
	})
}

func GetById(c echo.Context) error {
	param := c.Param("id")
	var user user.User
	config.DB.First(&user, param)
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "Get data success",
		Data: user,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user user.UserRequest
	// config.DB.First(&user, id)

	jsonData := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonData)
	
	if err != nil {
		panic(err)
	}

	user.Name 		 = jsonData["name"].(string)
	user.Email 	 	 = jsonData["email"].(string)
	user.Password  = utils.EncryptPassword([]byte(jsonData["password"].(string)))
	
	config.DB.Model(&user).Where("id = ?", id).Updates(&user)

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "Updated",
		Data: user,
	})

}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user user.UserRequest

	config.DB.Where("id = ?", id).Delete(&user)
	
	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "deleted",
		Data: "",
	})
}