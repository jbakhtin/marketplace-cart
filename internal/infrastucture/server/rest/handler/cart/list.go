package cart

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/response"
	"net/http"
)

type ListRequest struct {
	Test string `json:"test" validate:"required"`
}

type ListResponse struct {
	Items []Item `json:"items" validate:"required"`
}

type Item struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Price string `json:"price,omitempty" validate:"required"`
}

func (o *Handler) List(w http.ResponseWriter, r *http.Request) {
	var request ListRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	// TODO: add logic
	// ...

	listResponse := ListResponse{
		Items: make([]Item, 0),
	}

	response.WriteStandardResponse(w, r, http.StatusOK, listResponse, nil)
}
