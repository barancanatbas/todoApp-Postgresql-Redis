package request

type TodoInsert struct {
	Task string `json:"task" validate:"required"`
}
