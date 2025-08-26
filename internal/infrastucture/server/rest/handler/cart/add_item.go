package cart

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"net/http"
)

type AddItemRequest struct {
	Item  domain.SKU   `json:"item,omitempty" validate:"required"`
	Count domain.Count `json:"count,omitempty" validate:"required"`
}

func (o *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	var request AddItemRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	item := domain.Item{
		Sku:   request.Item,
		Count: request.Count,
	}

	err = o.useCase.AddItem(r.Context(), item)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusInternalServerError, nil, err)
		return
	}

	response.WriteStandardResponse(w, r, http.StatusCreated, nil, nil)
}
