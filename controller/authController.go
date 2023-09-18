package controller

import (
	"encoding/json"
	"net/http"
	"pos/config"
	"pos/model/user"
	"pos/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string   `json:"email"`
	jwt.RegisteredClaims
}

type CustomReponse struct {
	Name  string `json:"name"`
	Email string   `json:"email"`
	Token string   `json:"token"`
}

func LoginController(c echo.Context) error {
	user := user.UserRequest{}
	
	jsonData := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonData)

	if err != nil{
		panic(err)
	}
	// user.Email = jsonData["email"].(string)
	// user.Password = utils.EncryptPassword([]byte(jsonData["password"].(string)))
	// inputPassword := "mySecretPassword" 
	// storedHashedPassword := "$2a$12$D5Ei1Ht7uDbTpKXZLUUK.2kF0opTaaR7DntAJ39e7r4RmBC7AtFk3"

	password := jsonData["password"].(string)
	result := config.DB.Where("email", jsonData["email"].(string)).First(&user)

	if (result.Error != nil){
		return c.JSON(http.StatusBadRequest, utils.BaseResponse{
			Status: http.StatusBadRequest,
			Message: "bad request",
			Data: jsonData,
		})
	}

	err  = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		panic(err)
	}

	claims := &jwtCustomClaims{
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	// c.Bind(&user)

	res := CustomReponse{}

	res.Name = user.Name
	res.Email = user.Email
	res.Token = t

	return c.JSON(http.StatusOK, utils.BaseResponse{
		Status: http.StatusOK,
		Message: "getdata success",
		Data: res,
	})
}