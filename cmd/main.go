package main

import (
	"Leadgen/api"
	_ "Leadgen/docs"
	"Leadgen/internal/config"
	"Leadgen/internal/repositories"
	"Leadgen/internal/repositories/Building"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Leadgen
// @version 1.0
// @description Тестовое задание.

// @contact.name Константин
// @contact.url https://t.me/London68

func main() {
	//Инициализация
	router := gin.Default()
	fmt.Println("Чтение конфигурации")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cfg := config.New()
	fmt.Println("Подключение к базе данных")
	db := repositories.New(cfg.DB)

	buildingDB, err := Building.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	//
	router.POST("/building/", func(context *gin.Context) {
		api.BuildingHandlerAdd(context, buildingDB)
	})
	router.GET("/buildings/", func(context *gin.Context) {
		api.BuildingHandlerGet(context, buildingDB)
	})

	err = router.Run(fmt.Sprintf(":%d", cfg.API.Port))
	if err != nil {
		fmt.Printf(":%d", cfg.API.Port)
		log.Fatal(err)
		return
	}
}
