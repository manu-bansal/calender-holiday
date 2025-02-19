package controllers

import (
	"context"
	"holiday-calendar/config"
	"holiday-calendar/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddHoliday(c *gin.Context) {
	var holiday models.Holiday

	if err := c.ShouldBindJSON(&holiday); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	holiday.ID = primitive.NewObjectID()

	_, err := config.DB.InsertOne(context.TODO(), holiday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add holiday"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Holiday added successfully", "holiday": holiday})
}

func GetHolidays(c *gin.Context) {
	cursor, err := config.DB.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve holidays"})
		return
	}
	defer cursor.Close(context.Background())

	var holidays []models.Holiday
	for cursor.Next(context.Background()) {
		var holiday models.Holiday
		cursor.Decode(&holiday)
		holidays = append(holidays, holiday)
	}
	if holidays == nil {
		holidays = []models.Holiday{}
	}

	c.JSON(http.StatusOK, holidays)
}

func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = config.DB.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete holiday"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Holiday deleted successfully"})
}
