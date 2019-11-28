package entities

type ABServiceProcessRequest struct {
	Params string `validate:"required" example:"POST"`
}
