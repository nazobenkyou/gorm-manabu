package http_router

import "net/http"

func (c *Controller) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.TaskServicer.GetAll()
	if err != nil {
		c.CreateErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.CreateResponse(w, http.StatusOK, tasks)
}
