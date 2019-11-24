package user

type UserRequest struct {
	UserId     int                    `json:"user_id" validate:"required" example:"123456"`
	UpdateData map[string]interface{} `json:"update_data"`
}
