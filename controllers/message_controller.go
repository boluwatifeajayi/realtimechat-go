package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"chatapp/models"
)

func SendMessage(c *gin.Context, messageCollection *mongo.Collection) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.Timestamp = time.Now()

	_, err := messageCollection.InsertOne(context.TODO(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func GetMessages(c *gin.Context, messageCollection *mongo.Collection) {
	senderID := c.Param("sender_id")
	receiverID := c.Param("receiver_id")

	senderObjID, _ := primitive.ObjectIDFromHex(senderID)
	receiverObjID, _ := primitive.ObjectIDFromHex(receiverID)

	var messages []models.Message
	cursor, err := messageCollection.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"sender_id": senderObjID, "receiver_id": receiverObjID},
			{"sender_id": receiverObjID, "receiver_id": senderObjID},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var message models.Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, messages)
}
