package entities

//请求数据结构
type JobRequest struct {
	Uid string `json:"user_id" example:1212`
}

/*
 * validator.v9文档
 * 地址https://godoc.org/gopkg.in/go-playground/validator.v9
 * 列了几个大家可能会用到的，如有遗漏，请看上面文档
 */

//请求数据结构
type JobValidatorRequest struct {
	//tips，因为组件required不管是没传值或者传 0 or "" 都通过不了，但是如果用指针类型，那么0就是0，而nil无法通过校验
	Mobile string `json:"mobile" validate:"required" example:"157xxxxx"`
}