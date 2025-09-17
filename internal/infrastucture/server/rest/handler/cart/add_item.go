package cart

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	context2 "github.com/jbakhtin/marketplace-cart/internal/infrastucture/context"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"net/http"
)

type AddItemRequest struct {
	Sku   domain.SKU   `json:"sku,omitempty" validate:"required,numeric,min=0"`
	Count domain.Count `json:"count,omitempty" validate:"required,numeric,min=1"`
}

func (o *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	var request AddItemRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	item := domain.Item{
		Sku:   request.Sku,
		Count: request.Count,
	}

	userID := r.Context().Value(context2.UserIDKey) //ToDo: перенести ключ в домен, отвязать от логики контекста

	err = o.useCase.AddItem(r.Context(), userID.(domain.UserID), item)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusInternalServerError, nil, err)
		return
	}

	response.WriteStandardResponse(w, r, http.StatusCreated, nil, nil)
}
