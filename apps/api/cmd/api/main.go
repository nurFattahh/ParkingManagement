package main

import (
	"log"

	"WebParkir/apps/api/internal/delivery/http/handler"
	"WebParkir/apps/api/internal/infrastructure/config"
	"WebParkir/apps/api/internal/infrastructure/database"
	"WebParkir/apps/api/internal/repository"
	service "WebParkir/apps/api/internal/services"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectPostgres(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Auto migrate error", err)
	}

	log.Println("DATABASE_URL:", cfg.DatabaseURL)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	userRepo := repository.NewUserRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)

	authService := service.NewAuthService(userRepo)
	vehicleService := service.NewVehicleService(vehicleRepo)

	authHandler := handler.NewAuthHandler(authService)
	vehicleHandler := handler.NewVehicleHandler(vehicleService)

	api := r.Group("/api")

	api.GET("/health", handler.HealthCheckHandler)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.GET("/users", authHandler.GetUsers)
	api.POST("/adduser", authHandler.AddUser)

	api.GET("/vehicles", vehicleHandler.GetAllVehicles)
	api.GET("/vehicles/:ID", vehicleHandler.GetVehicleByID)
	api.GET("/vehicles/user/:lincensePlate", vehicleHandler.GetVehicleByLicensePlate)
	api.POST("/vehicle", vehicleHandler.AddVehicle)
	api.PUT("/vehicle/:ID", vehicleHandler.UpdateVehicle)
	api.DELETE("/vehicle/:ID", vehicleHandler.DeleteVehicle)

	log.Println("server running on :", cfg.Port)

	r.Run(":" + "8080")
}
