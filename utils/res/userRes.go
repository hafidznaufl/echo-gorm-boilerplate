package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndex(users []model.User) []web.UserReponse {
	var results []web.UserReponse
	for _, user := range users {
		userResponse := web.UserReponse{
			Id:       int(user.ID),
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		results = append(results, userResponse)
	}

	return results
}

func ConvertLogin(user model.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Email: user.Email,
		Password: user.Password,
	}
}

func ConvertGeneral(user *model.User) web.UserReponse {
	return web.UserReponse{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
