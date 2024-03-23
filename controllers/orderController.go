package controllers

import (
	"net/http"
	"strconv"
	"web-server/database"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderIDStr := ctx.Param("orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var order models.Order
	if err := db.Preload("Item").First(&order, orderID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderIDStr := ctx.Param("orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var updatedOrder models.Order
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	updatedOrder.ID = orderID

	if err := db.Save(&updatedOrder).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order":   updatedOrder,
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderIDStr := ctx.Param("orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&models.Order{}, orderID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
