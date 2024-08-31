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

// SendMessage handles the sending of a message.
func SendMessage(c *gin.Context, messageCollection *mongo.Collection) {
	// Create a variable to hold the message data
	var message models.Message

	// Bind the JSON request body to the message variable
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate the required fields in the message
	if message.SenderID.IsZero() || message.ReceiverID.IsZero() || message.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sender ID, Receiver ID, and Content are required"})
		return
	}

	// Set the timestamp for the message
	message.Timestamp = time.Now()

	// Insert the message into the MongoDB collection
	_, err := messageCollection.InsertOne(context.TODO(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

// GetMessages handles the retrieval of messages between a sender and a receiver.
func GetMessages(c *gin.Context, messageCollection *mongo.Collection) {
	// Get the sender and receiver IDs from the URL parameters
	senderID := c.Param("sender_id")
	receiverID := c.Param("receiver_id")

	// Validate that both sender and receiver IDs are provided
	if senderID == "" || receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sender ID and Receiver ID are required"})
		return
	}

	// Convert the sender ID to a MongoDB ObjectID
	senderObjID, err := primitive.ObjectIDFromHex(senderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Sender ID"})
		return
	}

	// Convert the receiver ID to a MongoDB ObjectID
	receiverObjID, err := primitive.ObjectIDFromHex(receiverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Receiver ID"})
		return
	}

	// Initialize a slice to hold the retrieved messages
	var messages []models.Message

	// Query the MongoDB collection for messages between the sender and receiver
	cursor, err := messageCollection.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"sender_id": senderObjID, "receiver_id": receiverObjID},
			{"sender_id": receiverObjID, "receiver_id": senderObjID},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode each message
	for cursor.Next(context.TODO()) {
		var message models.Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	// Return the retrieved messages
	c.JSON(http.StatusOK, messages)
}
