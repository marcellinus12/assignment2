package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Items struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}

type Order struct {
	OrderID      int     `json:"order_id"`
	CustomerName string  `json:"customer_name"`
	OrderedAt    string  `json:"ordered_at"`
	Items        []Items `json:"items"`
}

var OrderData = []Order{}

// func validation(order Order) error {

// }

func GetOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("orderID")
	orderID, err2 := strconv.Atoi(orderIDStr)
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadRequest, err2)
		return
	}

	condition := false

	var getOrder Order

	for _, order := range OrderData {
		if orderID == order.OrderID {
			condition = true

			getOrder = order
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("order with id %v not found", orderID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": getOrder,
	})
}

func CreateOrder(ctx *gin.Context) {
	var newOrder Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i := range newOrder.Items {
		newOrder.Items[i].OrderID = len(OrderData) + 1
	}

	newOrder.OrderID = len(OrderData) + 1
	OrderData = append(OrderData, newOrder)

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("orderID")
	orderID, err2 := strconv.Atoi(orderIDStr)
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadRequest, err2)
		return
	}

	condition := false

	var updatedOrder Order

	err := ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i, order := range OrderData {
		if orderID == order.OrderID {
			condition = true

			OrderData[i] = updatedOrder
			OrderData[i].OrderID = orderID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("order with id %v not found", orderID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meesage": fmt.Sprintf("order with id %v has been successfully updated", orderID),
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("orderID")
	orderID, err2 := strconv.Atoi(orderIDStr)
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadRequest, err2)
		return
	}

	condition := false

	var orderIndex int

	for i, order := range OrderData {
		if orderID == order.OrderID {
			condition = true

			orderIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("order with id %v not found", orderID),
		})

		return
	}

	copy(OrderData[orderIndex:], OrderData[orderIndex+1:])
	OrderData[orderID-1] = Order{}
	OrderData = OrderData[:len(OrderData)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"meesage": fmt.Sprintf("order with id %v has been successfully deleted", orderID),
	})
}
