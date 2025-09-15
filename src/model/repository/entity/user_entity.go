package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempy"`
	Password string             `bson:"password,omitempy"`
	Name     string             `bson:"name,omitempy"`
	Age      int                `bson:"age,omitempy"`
}
