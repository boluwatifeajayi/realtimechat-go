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
	"chatapp/utils"
)

func Register(c *gin.Context, userCollection *mongo.Collection) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email, and password are required"})
		return
	}

	// Check if email already exists
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	if user.ProfilePic == "" {
		user.ProfilePic = "https://www.shutterstock.com/image-vector/user-profile-icon-vector-avatar-600nw-2247726673.jpg"
	} else {
		user.ProfilePic = utils.UploadToCloudinary(user.ProfilePic)
	}

	user.JoinDate = time.Now()
	user.LastActive = time.Now()
	user.IsActive = true

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context, userCollection *mongo.Collection) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetAllUsers(c *gin.Context, userCollection *mongo.Collection) {
	var users []models.User
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func SearchUsers(c *gin.Context, userCollection *mongo.Collection) {
	query := c.Query("query")
	var users []models.User
	cursor, err := userCollection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": query, "$options": "i"}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func GetChatList(c *gin.Context, messageCollection *mongo.Collection) {
	userID := c.Param("user_id")
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var chatList []primitive.ObjectID
	cursor, err := messageCollection.Distinct(context.TODO(), "receiver_id", bson.M{"sender_id": userObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, id := range cursor {
		chatList = append(chatList, id.(primitive.ObjectID))
	}

	cursor, err = messageCollection.Distinct(context.TODO(), "sender_id", bson.M{"receiver_id": userObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, id := range cursor {
		chatList = append(chatList, id.(primitive.ObjectID))
	}

	c.JSON(http.StatusOK, chatList)
}

func GetUserByID(c *gin.Context, userCollection *mongo.Collection) {
	userID := c.Param("user_id")
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserProfile(c *gin.Context, userCollection *mongo.Collection) {
	userID := c.Param("user_id")
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	profile := map[string]interface{}{
		"username":    user.Name,
		"email":       user.Email,
		"join_date":   user.JoinDate,
		"last_active": user.LastActive,
	}

	c.JSON(http.StatusOK, profile)
}
