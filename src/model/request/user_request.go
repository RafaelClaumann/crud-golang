package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int    `json:"age" binding:"required,numeric,gte=18,lte=150"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"`
}
