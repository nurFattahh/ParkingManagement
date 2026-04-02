package handler

import (
	"WebParkir/apps/api/internal/domain"
	"WebParkir/apps/api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(service *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service}
}

func (h *VehicleHandler) AddVehicle(c *gin.Context) {
	var req domain.AddVehicleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	vehicle, err := h.service.AddVehicle(req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, vehicle)
}

func (h *VehicleHandler) GetAllVehicles(c *gin.Context) {
	vehicles, err := h.service.GetAllVehicles()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, vehicles)
}

func (h *VehicleHandler) GetVehicleByID(c *gin.Context) {
	id := c.Param("ID")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	vehicle, err := h.service.GetVehicleByID(uint(parsedID))

	if err != nil {
		c.JSON(404, gin.H{"error": "vehicle not found"})
		return
	}

	c.JSON(200, vehicle)
}

func (h *VehicleHandler) GetVehicleByLicensePlate(c *gin.Context) {
	licensePlate := c.Param("licensePlate")

	vehicles, err := h.service.GetVehicleByLicensePlate(licensePlate)
	if err != nil {
		c.JSON(404, gin.H{"error": "vehicle not found"})
		return
	}

	c.JSON(200, vehicles)
}

func (h *VehicleHandler) UpdateVehicle(c *gin.Context) {
	var req domain.Vehicle

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	vehicle, err := h.service.UpdateVehicle(req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, vehicle)
}

func (h *VehicleHandler) DeleteVehicle(c *gin.Context) {
	id := c.Param("ID")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	err = h.service.DeleteVehicle(uint(parsedID))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "vehicle deleted"})
}
