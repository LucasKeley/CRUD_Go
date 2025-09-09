package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	PassWord string `json:"password" binding:"required,min=6,max=20,containsany=@_!#$%&*-?"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int    `json:"age" binding:"required,min=12,max=140"`
}
