package models

type User struct {
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstname" json:"firstName"`
	LastName  string `bson:"lastname" json:"lastName"`
	Phone     string `bson:"phone" json:"phone"`
	Gender    string `bson:"gender" json:"gender"`
}


