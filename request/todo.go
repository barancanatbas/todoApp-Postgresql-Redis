package request

type TodoInsert struct {
	Task string `json:"task" validate:"required"`
}

type Completed struct {
	ID uint `query:"id" validate:"required"`
}
