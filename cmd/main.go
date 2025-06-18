// main.go
//
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
package main

import (
	"cushon/pkg/config"
	"cushon/pkg/customers"
	"cushon/pkg/deposit"
	"cushon/pkg/funds"
	"cushon/pkg/handlers"
	"cushon/pkg/storage"

	docs "cushon/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

func main() {
	config, _ := config.LoadConfig()
	store := storage.New(config)
	docs.SwaggerInfo.BasePath = ""
	defer store.Close()

	r := gin.Default()

	depositService := deposit.New(store)
	customerService := customers.New(store)
	fundsService := funds.New(store)

	// Register user routes
	handlers.RegisterFundRoutes(r, fundsService)
	handlers.RegisterCustomerRoutes(r, customerService)
	handlers.RegisterDepositRoutes(r, depositService)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
