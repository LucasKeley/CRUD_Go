package entity

type UserEntity struct {
	ID       string `bson:"_id,omitempy"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Age      int    `bson:"age"`
}
