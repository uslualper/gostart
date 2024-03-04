package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	Firstname   string             `bson:"firstname" json:"firstname"`
	Lastname    string             `bson:"lastname" json:"lastname"`
	Email       string             `bson:"email" json:"email"`
}
