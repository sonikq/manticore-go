package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserId  int    `json:"id"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type LogoutRequest struct {
	UserId int `json:"id" binding:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type Perms struct {
	Access int
	Code   int
}
