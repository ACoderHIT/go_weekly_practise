package entities

//请求数据结构
type ServiceResponse struct {
	Code    interface{} `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
