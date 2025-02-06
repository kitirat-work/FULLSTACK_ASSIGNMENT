package main

import (
	"api/config"
	"log"

	"github.com/labstack/echo/v4"
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
			NoLowerCase:   true,
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

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.Logger.Fatal(e.Start(":" + cfg.PORT))
}
