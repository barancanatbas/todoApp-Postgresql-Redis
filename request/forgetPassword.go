package request

type Forget struct {
	UserName string `json:"user_name" validate:"required"`
}
