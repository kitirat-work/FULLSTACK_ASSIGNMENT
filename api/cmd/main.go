package main

import (
	"api/cache"
	"api/config"
	"api/controller"
	"api/repository"
	"api/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	/* cache */
	cacheInstance := cache.NewLoginCache()
	cacheInstance.InitCacheReset()

	// repository
	userRepo := repository.NewUserRepository(gormDB)

	// service
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(cacheInstance, userRepo)

	// controller
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	// routing
	e.GET("/user/:id", userController.GetById)

	e.POST("/auth/login/pin", authController.LoginPin)

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

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

	// Start server
	go func() {
		if err := e.Start(":" + cfg.PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}
