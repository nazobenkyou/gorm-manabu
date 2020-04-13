package service

import "github.com/nazobenkyou/gorm-manabu/model"

func (t *TaskService) GetAll() ([]model.Task, error) {
	tasksRepository := t.Repository.GetAll()

	tasks := make([]model.Task, 0)
	for _, t := range tasksRepository {
		tt := model.Task{
			EntityID:    int(t.ID),
			Name:        t.Name,
			Description: t.Description,
			StartDate:   t.StartDate,
			DueDate:     t.DueDate,
		}
		tasks = append(tasks, tt)
	}

	return tasks, nil
}
