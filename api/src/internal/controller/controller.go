package controller

import "encoding/json"
import "net/http"

type Response struct {
	Status  string `json:"status, omitempty"`
	Message string `json:"message, omitempty"`
	Data    any    `json:"data, omitempty"`
}

func (r *Response) toJson() ([]byte, error) {
	response, err := json.Marshal(*r)
	return response, err
}

func NewResponseJson(status string, message string, data any) ([]byte, error) {
	response := &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response.toJson()
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Ping(res http.ResponseWriter, req *http.Request) {
	response := &Response{
		Message: "Pong",
	}

	res.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(res)
	err := encoder.Encode(response)
	if err != nil {
		encoder.Encode(&Response{
			Status:  "failed",
			Message: "Internal Error",
		})
	}
}
