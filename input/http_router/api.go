package http_router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nazobenkyou/gorm-manabu/repository"
	"github.com/nazobenkyou/gorm-manabu/service"
	"net/http"
	"time"
)

type (
	HttpUtils struct{}

	HttpUtilities interface {
		CreateErrorResponse(w http.ResponseWriter, code int, message string)
		CreateResponse(w http.ResponseWriter, code int, data ...interface{})

		methodNotAllowedHandler(w http.ResponseWriter, r *http.Request)
		notFoundHandler(w http.ResponseWriter, r *http.Request)
	}

	Controller struct {
		service.TaskServicer

		HttpUtilities
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Date    string `json:"date"`
	}
)

func (h *HttpUtils) CreateErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("application/json", "application/json")

	response, _ := json.Marshal(&ErrorResponse{Message: message, Code: code, Date: time.Now().UTC().Format(time.RFC3339)})

	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (h *HttpUtils) CreateResponse(w http.ResponseWriter, code int, data ...interface{}) {
	w.Header().Set("application/json", "application/json")

	w.WriteHeader(code)
	if len(data) == 0 {
		return
	}

	response, _ := json.Marshal(data[0])
	_, _ = w.Write(response)
}

func (h *HttpUtils) methodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	h.CreateErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	return
}

func (h *HttpUtils) notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	h.CreateErrorResponse(w, http.StatusNotFound, "url not found")
	return
}

func NewRouter(conn repository.Repository) *mux.Router {
	r := mux.NewRouter()

	c := &Controller{&service.TaskService{Repository: conn}, &HttpUtils{}}

	r.MethodNotAllowedHandler = http.Handler(http.HandlerFunc(c.methodNotAllowedHandler))
	r.NotFoundHandler = http.Handler(http.HandlerFunc(c.notFoundHandler))

	r.HandleFunc("/tasks", c.GetTasks).Methods(http.MethodGet)

	return r
}
