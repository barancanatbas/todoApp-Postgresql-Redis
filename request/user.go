package request

type Login struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
