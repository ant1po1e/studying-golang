package entity

type Person struct {
	Nickname string `json:"nickname" binding:"required"`
	Age      int8   `json:"age" binding:"gte=5,lte=80"`
	Email    string `json:"email" validate:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=100" validate:"is-cool"`
	Description string `json:"description" binding:"max=200"`
	Author      Person `json:"author" binding:"required"`
}
