package routes

import (
	"github.com/gin-gonic/gin"

	"chatapp/config"
	"chatapp/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userCollection := config.Client.Database("chat_app").Collection("users")
	messageCollection := config.Client.Database("chat_app").Collection("messages")

	api := r.Group("/api")
	{
		api.POST("/register", func(c *gin.Context) {
			controllers.Register(c, userCollection)
		})
		api.POST("/login", func(c *gin.Context) {
			controllers.Login(c, userCollection)
		})
		api.POST("/send-message", func(c *gin.Context) {
			controllers.SendMessage(c, messageCollection)
		})
		api.GET("/messages/:sender_id/:receiver_id", func(c *gin.Context) {
			controllers.GetMessages(c, messageCollection)
		})
		api.GET("/users", func(c *gin.Context) {
			controllers.GetAllUsers(c, userCollection)
		})
		api.GET("/users/search", func(c *gin.Context) {
			controllers.SearchUsers(c, userCollection)
		})
		api.GET("/users/:user_id/chat-list", func(c *gin.Context) {
			controllers.GetChatList(c, messageCollection)
		})
		api.GET("/users/:user_id", func(c *gin.Context) {
			controllers.GetUserByID(c, userCollection)
		})
		api.GET("/users/:user_id/profile", func(c *gin.Context) {
			controllers.GetUserProfile(c, userCollection)
		})
	}

	return r
}
