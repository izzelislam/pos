package routes

import (
	controller "pos/controller"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string   `json:"email"`
	jwt.RegisteredClaims
}

func Route(e *echo.Echo) *echo.Echo{
	
	e.POST("/api/v1/login", controller.LoginController)

	
	// e.Use(echojwt.WithConfig(config))
	
	routeGroup := e.Group("/api/v1")
	routeGroup.Use(echojwt.JWT([]byte("secret")))
	routeGroup.GET("/user", controller.GetUser)
	routeGroup.POST("/user", controller.CreateUser)
	routeGroup.GET("/user/:id", controller.GetById)
	routeGroup.PUT("/user/:id", controller.UpdateUser)
	routeGroup.DELETE("/user/:id", controller.DeleteUser)

	routeGroup.GET("/category", controller.GetCategory)
	routeGroup.POST("/category", controller.StoreCategory)
	routeGroup.GET("/category/:id", controller.ShowCategory)
	routeGroup.PUT("/category/:id", controller.UpdateCategory)
	routeGroup.DELETE("/category/:id", controller.DeleteCategory)

	routeGroup.GET("/vendor", controller.GetVendor)
	routeGroup.POST("/vendor", controller.StoreVendor)
	routeGroup.GET("/vendor/:id", controller.ShowVendor)
	routeGroup.PUT("/vendor/:id", controller.UpdateVendor)
	routeGroup.DELETE("/vendor/:id", controller.DeleteVendor)

	routeGroup.GET("/product", controller.GetProduct)
	routeGroup.POST("/product", controller.StoreProduct)
	routeGroup.GET("/product/:id", controller.ShowProduct)
	routeGroup.PUT("/product/:id", controller.UpdateProduct)
	routeGroup.DELETE("/product/:id", controller.DeleteProduct)

	return e
}