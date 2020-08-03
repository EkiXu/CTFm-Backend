package response

import (
	"ctfm_backend/models"
)

type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UsersListResponse struct {
	Users []UserResponse `json:"challenges"`
}

type UserdAddedResponse struct {
	ID uint `json:"id"`
}
