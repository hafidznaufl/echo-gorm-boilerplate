package req

import (
	"app/model"
	"app/model/web"
)

func PassBody(users web.UserRequest) *model.User {
	return &model.User{
		Name: users.Name,
		Email: users.Email,
		Password: users.Password,
	}
}
