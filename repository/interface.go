package repository

type (
	Repository interface {
		GetAll() []Task
	}
)
