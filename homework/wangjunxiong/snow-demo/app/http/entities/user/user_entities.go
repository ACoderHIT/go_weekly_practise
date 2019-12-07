package user

type UserRequest struct {
	UserId     int                    `json:"user_id" validate:"required" example:"123456"`
	UpdateData map[string]interface{} `json:"update_data"`
}

type InfoRequest struct {
	UserId int `form:"user_id"`
}
type InfoJsonRequest struct {
	UserId interface{} `json:"user_id"`
}
