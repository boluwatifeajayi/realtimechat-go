package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	ProfilePic string             `json:"profile_pic" bson:"profile_pic"`
	JoinDate   time.Time          `json:"join_date" bson:"join_date"`
	LastActive time.Time          `json:"last_active" bson:"last_active"`
	IsActive   bool               `json:"is_active" bson:"is_active"`
}
