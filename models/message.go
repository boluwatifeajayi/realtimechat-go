package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SenderID   primitive.ObjectID `json:"sender_id" bson:"sender_id"`
	ReceiverID primitive.ObjectID `json:"receiver_id" bson:"receiver_id"`
	Content    string             `json:"content" bson:"content"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
}
