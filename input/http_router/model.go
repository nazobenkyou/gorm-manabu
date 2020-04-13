package http_router

type (
	Task struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		DueDate     *string `json:"dueDate"`
		StartDate   *string `json:"startDate"`
	}
)
