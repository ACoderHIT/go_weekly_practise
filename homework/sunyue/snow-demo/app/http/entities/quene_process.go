package entities

type QueneProcessRequest struct {
	Name string `validate:"required" example:"xxx"`
	Id int `validate:"required" example:1`
	Update bool
	Num int
}