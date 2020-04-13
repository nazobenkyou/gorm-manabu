package model

type (
	Task struct {
		EntityID    int
		Name        string
		Description string
		StartDate   *string
		DueDate     *string
	}
)
