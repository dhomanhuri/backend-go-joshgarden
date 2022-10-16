package dto

import "backend-go/features/users"

type UserResquest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	RoleID   int    `json:"role_id" form:"role_id" binding:"required"`
}

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
}

type LoginResquest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
}

func ToCore(userReq LoginResquest) users.Core {
	userCore := users.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
