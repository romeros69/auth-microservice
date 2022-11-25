package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectID primitive.ObjectID `json:"_id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
