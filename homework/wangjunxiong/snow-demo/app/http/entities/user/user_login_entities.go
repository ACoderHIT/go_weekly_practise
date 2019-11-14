package user

import "time"

type UserLoginRequest struct {
	UserId    int       `json:"user_id" validate:"required" example:"123456"`
	Ip        string    `json:"ip" validate:"" example:"10.11.11.11"`
	LoginTime time.Time `json:"login_time"`
}
