package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	err = db.Create(&newOrder).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order

	db := database.GetDB()

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("orderID")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updatedOrder models.Order

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedOrder.OrderID = uint(orderID)

	db := database.GetDB()

	tx := db.Begin()

	var existingOrder models.Order

	if err := tx.First(&existingOrder, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data not found",
				"error_message": fmt.Sprintf("Order with ID %s not found", strconv.FormatUint(orderID, 10)),
			})
			return
		}

		tx.Rollback()
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := tx.Model(&models.Order{}).Where("id = ?", orderID).Updates(&updatedOrder).Error; err != nil {
		tx.Rollback()
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := range updatedOrder.Items {
		updatedOrder.Items[i].OrderID = updatedOrder.OrderID
		if err := tx.Model(&updatedOrder.Items[i]).Where("order_id = ?", orderID).Updates(&updatedOrder.Items[i]).Error; err != nil {
			tx.Rollback()
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with ID %d has been updated", orderID),
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

	db := database.GetDB()

	tx := db.Begin()

	var existingOrder models.Order

	if err := tx.First(&existingOrder, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data not found",
				"error_message": fmt.Sprintf("Order with ID %s not found", orderID),
			})
			return
		}

		tx.Rollback()
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := tx.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		tx.Rollback()
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := tx.Where("id = ?", orderID).Delete(&models.Order{}).Error; err != nil {
		tx.Rollback()
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with ID %s has been deleted", orderID),
	})
}
