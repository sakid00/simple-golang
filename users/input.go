package users

type RegisterUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type LoginIniput struct {
	Email string `json:"email" binding:"required"`
}
