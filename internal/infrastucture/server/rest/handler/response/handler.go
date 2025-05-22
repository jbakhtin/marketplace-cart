package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func NewSuccessResponse(status int, data any) Response {
	return Response{
		Success: true,
		Status:  status,
		Data:    data,
	}
}

func NewErrorResponse(status int, err error) Response {
	return Response{
		Success: false,
		Status:  status,
		Error: &Error{
			Title: err.Error(),
		},
	}
}

func WriteStandardResponse(w http.ResponseWriter, r *http.Request, code int, payload any, err error) {
	var jsonResponse Response
	if err != nil {
		jsonResponse = NewErrorResponse(code, err)
	} else {
		jsonResponse = NewSuccessResponse(code, payload)
	}

	buf, err := json.Marshal(&jsonResponse)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
}
