package response

import (
	"Go_Code/src/project_training/gin/chapter4/model"
)

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUserDto
// DTO(Data Transfer Object) 用于展示层与服务层之间的数据传输对象
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
