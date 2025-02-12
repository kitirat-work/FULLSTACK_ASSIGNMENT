package main

import (
	"api/config"
	"api/controller"
	"api/repository"
	"api/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	e := echo.New()
	cfg := config.New()

	gormDB, err := gorm.Open(mysql.Open(cfg.SQLConnStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			// NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatalf("fail %v!!", err)
	}

	defer func() {
		sqlDB, _ := gormDB.DB()
		if err = sqlDB.Close(); err != nil {
			log.Printf("fail to close sql's connection: %v\n", err)
		}
	}()

	// repository
	userRepo := repository.NewUserRepository(gormDB)

	// service
	userService := service.NewUserService(userRepo)

	// controller
	userController := controller.NewUserController(userService)

	// routing
	e.GET("/user/:id", userController.GetById)

	// add a middleware to write logs
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Printf("request: %s %s\n", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	})

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // ถ้าต้องการให้ API รองรับ Cookies
	}))

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.Logger.Fatal(e.Start(":" + cfg.PORT))
}
