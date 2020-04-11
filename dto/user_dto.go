package dto

import "hackweek/main/model"

type UserDTO struct {
	Name string `json:"name"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name: user.Name,
	}
}
