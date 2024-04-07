package users

type Formatter struct {
	Id    int    `json:"id"`
	Name  string `json:"name`
	Email string `json:"email"`
}

func UserFormatter(user User) Formatter {
	processformatter := Formatter{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return processformatter
}
