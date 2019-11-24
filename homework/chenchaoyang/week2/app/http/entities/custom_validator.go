package entities

//请求数据结构
type CustomValidatorRequest struct {
	From  interface{} `json:"from" validate:"required" example:"snow"`
	Value interface{} `json:"value" validate:"required" example:"3"`
}

//请求数据结构
type ServiceResponse struct {
	Code    interface{} `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	//From interface{} `json:"from" validate:"required" example:"snow"`
	//Value interface{} `json:"value" validate:"required" example:"3"`
}
